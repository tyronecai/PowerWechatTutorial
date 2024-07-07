package merchantService

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/payment/merchantService/request"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"power-wechat-tutorial/services"
)

func APIComplaints(c *gin.Context) {

	para := &request.RequestComplaints{
		Limit:            1,
		Offset:           10,
		BeginDate:        "2024-01-01",
		EndDate:          "2024-04-01",
		ComplaintedMchId: "1616273230",
	}

	rs, err := services.PaymentApp.MerchantService.Complaints(c.Request.Context(), para)
	if err != nil {
		log.Println("出错了： ", err)
		c.String(400, err.Error())
		return
	}
	c.JSON(http.StatusOK, rs)

}
