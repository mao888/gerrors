# gerrors方法

- **func New(code int, msg string) error**
- **func Wrap(err error, msg string) error**
- **func WrapCode(err error, code int, msg string) error**

## 使用场景

**系统error：不是通过gerrors创建的错误。**

| **方法**     | **说明**                       | **示例**                                         | **日志**                                                     | **HTTP返回数据格式**                                   |
| ------------ | ------------------------------ | ------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------ |
| **New**      | 创建一个业务错误               | gerrors.New(100001,"这是一个业务错误")           | 会打印Warn级别日志                                           | `{    "code": 100001,    "msg": "这是一个业务错误" }`  |
| **Wrap**     | 包装一个错误                   | gerrors.Wrap(err, "类.方法 调用异常.")           | 是否包含系统error.存在打印Error级别日志不存在打印Warn级别日志 | `{    "code": -1,    "msg": "internal server error" }` |
| **WrapCode** | 包装一个错误，且返回业务错误码 | gerrors.WrapCode(err, 100002,"这是一个包装消息") | 是否包含系统error.存在打印Error级别日志不存在打印Warn级别日志 | `{    "code": 100002,    "msg": "这是一个包装消息" }`  |

# **案例描述**

**需求：上传文件当文件存在时告诉用户文件已经存在，如果创建的文件名称为aaa时提示用户名称错误。**

- handler

```go
func (h *FileHandler) File(c *gin.Context) {
   ctx := c.Request.Context()
   req := &bean.FileReq{}
   //true：绑定Header，false：不绑定
   if !h.Bind(c, req, false) {
      return
   }
   resp, err := h.fileLogic.File(ctx, req.ID)
   if err != nil {
      h.Fail(c, err)
      return
   }
   h.Success(c, resp)
}
```

- logic

```go
func (l *FileLogic) File(ctx context.Context, id int64) (string, error) {
   data, err := l.FileService.File(ctx, id)
   if err != nil {
      return "", gerrors.Wrap(err,"FileLogic.File upload file.")
   }
   return data, nil
}
```

- service

```go
func (s *FileService) File(ctx context.Context, id int64) (string, error) {
   file,err := os.Open("1.txt");
   if  errors.Is(os.ErrExist,err) {
      return "",gerrors.WrapCode(err,100001,"文件已经存在")
   }else if err != nil {
      return "", err
   }
   if file.Name() == "aaa" {
      return "", gerrors.New(100002,"文件名称不能为aaa")
   }
   return "", nil
}
```
