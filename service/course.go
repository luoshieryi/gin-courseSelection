package service

import (
	"errors"
	"project/dao"
	"project/util/logs"
	"sync"
)

/*
 @Author: as
 @Date: Creat in 15:52 2022/2/12
 @Description: user.go
*/

var (
	CourseOver = errors.New("Course No Cap")
	StuHaveCourse=errors.New("Stu have this Course")
)

func BookCourse(courseId, stuId string) error {
	//TODO: 查看是否绑定
	if dao.StuHaveCourse(courseId,stuId){
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
	}
	info, ok = val.(*courseInfo)
	if !ok {
		return info, errors.New("Unknown error")
	}
	return info, nil
}

// 执行抢课
func (c *courseInfo) bookOne(courseId,stuId string) error {
	c.lock.Lock()
	defer c.lock.Unlock()
	if c.isOver || c.getLeft() == 0 {
		return CourseOver
	}
	c.getOne()
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
