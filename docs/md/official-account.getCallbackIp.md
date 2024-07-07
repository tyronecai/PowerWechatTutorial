
SDK产品接口的代码展示：
```
返回类型定义如下：
type ResponseGetCallBackIP struct {
	response.ResponseOfficialAccount

	IPList []string `json:"ip_list"`
}

具体使用接口方式：
func GetCallbackIP(ctx *gin.Context) {
	data, err := services.OfficialAccountApp.Base.GetCallbackIP(ctx.Request.Context())
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, data)
}
```