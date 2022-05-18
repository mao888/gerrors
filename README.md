# gerrors

#### 自定义异常处理包
* 创建业务异常
```go
    gerrors.New(constants.ErrorConfigCode, constants.ErrorConfigMsg)
```

* 使用（Handler调用Logic，Logic抛出一个业务错误）
```go
func (h *InitHandler) Post() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		request := &bean.InitRequest{}
		if !h.Bind(c, request, true) {
			return
		}
		initConfig, err := h.initLogic.GetConfig(ctx, request)
		if err != nil {
			glog.Errorf(ctx, "generate init pos error. app_id: %s, sdk_id: %s", request.Header.AppID)
			code, msg := gerrors.Resp(err)
			h.Fail(c, code, msg)
			return
		}
		h.Success(c, initConfig)
	}
}    

func (s *InitLogic) GetConfig(ctx context.Context, platforms []string, appID string) (map[string]interface{}, error) {
	glog.Infof(ctx, "GetConfig start. app_id: %s", appID)
	platformConfig := cache.GetPlatformConfig(appID)
	if platformConfig == nil {
		return nil, gerrors.New(constants.ErrorConfigCode, constants.ErrorConfigMsg)
	}
}

```
