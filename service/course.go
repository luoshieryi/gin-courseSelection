package service

import (
	"errors"
	"project/dao"
	"project/model"
	"project/types"
	"project/util/logs"
	"strconv"
	"sync"
)

var (
	CourseOver    = errors.New("Course No Cap")
	StuHaveCourse = errors.New("Stu have this Course")
)

// CreateCourse 创建课程
func CreateCourse(request types.CreateCourseRequest) (string, types.ErrNo) {
	course := dao.GetCourseByName(request.Name)
	if course.ID != 0 {
		return strconv.FormatInt(course.ID, 10), types.UnknownError
	}

	entity := model.Course{
		Name: request.Name,
		Cap:  request.Cap,
	}
	id, err := dao.CreateCourse(entity)
	if err != nil {
		return "", types.UnknownError
	}
	return strconv.FormatInt(id, 10), types.OK
}

func GetCourse(request types.GetCourseRequest) (types.TCourse, types.ErrNo) {
	tCourse := types.TCourse{}
	_id, err := strconv.ParseInt(request.CourseID, 10, 64)
	if err != nil {
		return tCourse, types.UnknownError
	}

	course := dao.GetCourseById(_id)
	if course.ID == 0 {
		return tCourse, types.CourseNotExisted
	}

	res := types.TCourse{
		CourseID:  strconv.FormatInt(course.ID, 10),
		Name:      course.Name,
		TeacherID: strconv.FormatInt(course.TeacherID, 10),
	}
	return res, types.OK
}

func BindCourse(request types.BindCourseRequest) types.ErrNo {
	Id, err := strconv.ParseInt(request.CourseID, 10, 64)
	if err != nil {
		return types.UnknownError
	}

	teacherId, err := strconv.ParseInt(request.TeacherID, 10, 64)
	if err != nil {
		return types.UnknownError
	}

	res := dao.GetCourseById(Id)
	if res.ID == 0 {
		return types.CourseNotExisted
	}
	if res.TeacherID != 0 {
		return types.CourseHasBound
	}

	res.TeacherID = teacherId
	err = dao.UpdateCourse(res)
	if err != nil {
		return types.UnknownError
	}
	return types.OK
}

//其实teacherid用不上
func UnbindCourse(request types.UnbindCourseRequest) types.ErrNo {
	Id, err := strconv.ParseInt(request.CourseID, 10, 64)
	if err != nil {
		return types.UnknownError
	}

	teacherId, err := strconv.ParseInt(request.TeacherID, 10, 64)
	if err != nil {
		return types.UnknownError
	}

	res := dao.GetCourseById(Id)
	if res.ID == 0 {
		return types.CourseNotExisted
	}
	if res.TeacherID == 0 {
		return types.CourseNotBind
	}
	if res.TeacherID != teacherId {
		return types.UnknownError
	}
	res.TeacherID = 0
	//println(res.ID)
	err = dao.DeleteTeacherByID(res.ID)
	if err != nil {
		return types.UnknownError
	}
	return types.OK
}

func GetTeacherCourse(request types.GetTeacherCourseRequest) ([]*types.TCourse, types.ErrNo) {
	tcource := make([]*types.TCourse, 0)
	TeacherID, err := strconv.ParseInt(request.TeacherID, 10, 64)
	if err != nil {
		return tcource, types.UnknownError
	}
	tcource, err = dao.GetCourseByTeacherId(TeacherID)
	if err != nil {
		return tcource, types.UnknownError
	}
	return tcource, types.OK
}

func BookCourse(courseId, stuId string) error {
	if dao.StuHaveCourse(courseId, stuId) {
		return StuHaveCourse
	}
	// 抢课
	info, err := getCourseInfo(courseId)
	if err != nil {
		return err
	}
	// 课程抢完了
	if info.isOver {
		return CourseOver
	}
	return info.bookOne(courseId, stuId)

}

type courseInfo struct {
	cap    int
	left   int // 剩余
	leftCh chan int
	done   chan struct{}
	close  chan struct{}
	lock   *sync.Mutex
	isOver bool
}

var allCourse = sync.Map{}

func getCourseInfo(courseId string) (*courseInfo, error) {
	var info *courseInfo
	val, ok := allCourse.Load(courseId)
	if !ok {
		ccap, err := dao.GetCourseCap(courseId)
		if err != nil {
			logs.PrintLogErr(logs.Service, courseId+"get Cap error", err)
			return info, err
		}
		info = &courseInfo{
			cap:    ccap,
			left:   ccap,
			leftCh: make(chan int),
			done:   make(chan struct{}),
			close:  make(chan struct{}),
			lock:   &sync.Mutex{},
		}
		allCourse.Store(courseId, &info)
		info.monitor() // 开始监听
	} else {
		//fmt.Println(val)
		info, ok = val.(*courseInfo)
		if !ok {
			//fmt.Println("notOk")
			return info, errors.New("Unknown error")
		}
	}

	return info, nil
}

// 执行抢课
func (c *courseInfo) bookOne(courseId, stuId string) error {
	c.lock.Lock()
	defer c.lock.Unlock()
	if c.isOver || c.getLeft() == 0 {
		return CourseOver
	}
	c.getOne()

	//fmt.Println("bookOneError")
	// DB写入肯定比抢课的流程慢
	return dao.BookCourse(courseId, stuId)
}

func (c *courseInfo) getOne() {
	c.done <- struct{}{}
}

func (c *courseInfo) getLeft() int {
	l := <-c.leftCh
	return l
}

func (c *courseInfo) monitor() {
	go c.bookCourse()
}

func (c *courseInfo) bookCourse() {
	for {
		select {
		case <-c.done:
			c.left -= 1
			//fmt.Println(c.left)
			if c.left == 0 {
				c.isOver = true
				close(c.close)
			}
		case c.leftCh <- c.left: // 保持同步
		case <-c.close:
			//fmt.Println("close")
			return
		}
	}
}
