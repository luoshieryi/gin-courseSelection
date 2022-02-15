# design document

Contributors:

- luoshieryi : https://github.com/luoshieryi , developer, debugger
- the-xin : https://github.com/the-xin , developer
- z-anshun : https://github.com/z-anshun , developer

## global

- use `\mod\github.com\go-playground\validator\v10` to validate request params 
  - through tags in struct for example ```	Nickname string   `form:"Nickname" json:"Nickname" binding:"required,min=4,max=20,alpha"` ```
  - in `*gin.Context.ShouldBind`
- use new function XXXRes to create response structure
- a log tool in project/util/logs
- use relational database MySQL, gorm as db and orm

## auth module

- login @luoshieryi 
  - create new random session for member
  - use gin.Context.setCookie() to store session
- logout @luoshieryi
  - use gin.Context.Cookie() to get session-code
  - delete session 
  - delete cookie through setting cookie's Expires to -1  
- whoami @luoshieryi
  - use gin.Context.Cookie() to get session-code
  - find memberID in model.Session
  - find member in model.Member

## member module

- create @luoshieryi
  - create new member and store it in db.Member
  - Nickname can only be upper or low case English
  - Username Ditto
  - Password must include upper, low case English, number and can only include them
- get @luoshieryi
  - get member by ID from db.Member
- getList @luoshieryi
  - get member in page from db.Member
- update @luoshieryi
  - can only update nickname
- delete @luoshieryi
  - delete member in db.Member
  - delete member's sessions in db.Session

## course module

### course

- create @luoshieryi
  - create new course and store it in db.Course
- get @luoshieryi
  - get course by ID from db.Member
### teacher

- bind @the-xin
  - add course.Teacher in db.Course.Teacher
- unbind @the-xin
  - delete course.Teacher in db.Course.Teacher
- get @the-xin
  - get teacher's binding courses by TeacherID
- schedule @the-xin
  - return best idea for scheduling courses and teachers

### student

- book @z-anshun
  - sync.Map stores all course info
  - goroutine pool to reuse goroutine
  - redis or go-cache to cache data
  - *gin.Context.lock to protect data
- getStudent @luoshieryi
  - get courseIDs in db.Stu_Courses by studentID
  - get courses in db.Course by ID

