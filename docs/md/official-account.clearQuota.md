
SDK产品接口的代码展示：
```
返回类型定义如下：
type ResponseOfficialAccount struct {
	ResponseBase

	ErrCode int    `json:"errcode,omitempty"`
	ErrMsg  string `json:"errmsg,omitempty"`

	ResultCode string `json:"resultcode,omitempty"`
	ResultMsg  string `json:"resultmsg,omitempty"`
}

具体使用接口方式：
func ClearQuota(ctx *gin.Context) {
	data, err := services.OfficialAccountApp.Base.ClearQuota(ctx.Request.Context())
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, data)
}
```