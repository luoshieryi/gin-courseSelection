# design document

Contributors:

- 冯浩泽 : https://github.com/luoshieryi , developer, debugger
- 杨新纪 : https://github.com/the-xin , developer
- 张安顺 : https://github.com/z-anshun , developer

## global

- use `\mod\github.com\go-playground\validator\v10` to validate request params 
  - through tags in struct for example ```	Nickname string   `form:"Nickname" json:"Nickname" binding:"required,min=4,max=20,alpha"` ```
  - in `*gin.Context.ShouldBind`
- use new function XXXRes to create response structure
- a log tool in project/util/logs
- use relational database MySQL, gorm as db and orm

## auth module

- login @冯浩泽 
  - create new random session for member
  - use gin.Context.setCookie() to store session
- logout @冯浩泽
  - use gin.Context.Cookie() to get session-code
  - delete session 
  - delete cookie through setting cookie's Expires to -1  
- whoami @冯浩泽
  - use gin.Context.Cookie() to get session-code
  - find memberID in model.Session
  - find member in model.Member

## member module

- create @冯浩泽
  - create new member and store it in db.Member
  - Nickname can only be upper or low case English
  - Username Ditto
  - Password must include upper, low case English, number and can only include them
- get @冯浩泽
  - get member by ID from db.Member
- getList @冯浩泽
  - get member in page from db.Member
- update @冯浩泽
  - can only update nickname
- delete @冯浩泽
  - delete member in db.Member
  - delete member's sessions in db.Session

## course module

### course

- create @冯浩泽
  - create new course and store it in db.Course
- get @冯浩泽
  - get course by ID from db.Member
### teacher

- bind @杨新纪
  - add course.Teacher in db.Course.Teacher
- unbind @杨新纪
  - delete course.Teacher in db.Course.Teacher
- get @杨新纪
  - get teacher's binding courses by TeacherID
- schedule @杨新纪
  - return best idea for scheduling courses and teachers

### student

- book @张安顺
  - sync.Map stores all course info
  - goroutine pool to reuse goroutine
  - redis or go-cache to cache data
  - *gin.Context.lock to protect data
- getStudent @冯浩泽
  - get courseIDs in db.Stu_Courses by studentID
  - get courses in db.Course by ID

