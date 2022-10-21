# gerrors

#### 自定义异常处理包
* 创建业务异常(默认不设置堆栈，可以再constants中预先定义)
```go
    gerrors.New(constants.ErrorConfigCode, constants.ErrorConfigMsg)
```

* 包装异常
```go
    gerrors.Wrap(err, "exec1 wrap")
```
