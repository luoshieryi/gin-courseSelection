# 开发记录

## 2.xx day0

熟悉go与gin的使用中..
不适应的目录结构折腾好一会, 强迫症++

##2.10 day1

### 代码与服务器

代码自动同步到服务器中
1. 配置代码上传
   1. goland -> Tools -> Development -> Configuration
   2. +(Add) -> SFTP -> New Servier Name -> 随便输入个名字 
   3. Connection -> SSH Configuration -> ... -> +(Add) -> Host（输入 SSH Host），User Name（输入 SSH User），Password（输入 SSH Password）-> Test Connection -> 成功后点击 OK
   4. 回到 Connection 页面 -> Root Path（输入 /root），点击 Mappings -> Deployment path（输入 /go/helloWord）-> OK
2. 配置自动同步
   1. Tools -> Deployment -> Automatic Upload
3. 上传代码
   1. 右键要上传到云端的文件（点击整个项目）-> Deployment -> Upload to ...

部署运行
1. ssh连接服务器, `ssh root@{SSH Host}`
2. 启动服务 
    ```bash
   cd /root/go/helloWorld/
   go run . 
   ```
3. 使用 tmux 部署服务
   
### 接口实现
- 内嵌匿名结构体的初始化有点麻烦
- 函数参数不支持默认值
- 不能 throw-catch 的error处理, 使用多返回值层级返回
- 实现了member, auth 的 8 个接口

### 参数绑定与验证
使用 shouldBind 绑定参数与结构体, 在 binding 标签中设置校验规则(基于validate)

- 例如 ``` Username string   `form:"Username" json:"Username" binding:"required,min=8,max=20,alpha"` ``` 设置传参方式为form/json, 表单参数/json字段名为Username, [8,20], 仅包含字母
- 实际shouldBind会自动匹配绑定方式, from/json标签的作用仅指定字段名称, 也可以使用shouldBindJson具体指定绑定方式
- 校验同时包括大小写、数字...`alphanum,containsany=abcdefghijklmnopqrstuvwxyz,containsany=0123456789,containsany=ABCDEFGHIJKLMNOPQRSTUVWXYZ"`

## 2.11 day2
接了一点 react 咕咕咕中

- validator 验证字符串长度时自动匹配了字符的长度, 而不是字节长度, 为nickname写了一个验证器
- 仅有 CreateMember 需要鉴权, 鉴权方法内置在相关 api 与 service 中, 减少其他请求的开销