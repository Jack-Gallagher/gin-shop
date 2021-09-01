package controller

import (
	"ginShop/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type MemberController struct {
}

func (mc *MemberController) Router(engine *gin.Engine) {
	engine.GET("/api/sendCode", mc.SendSmsCode)
}

// SendSmsCode https://localhost:8080/api/sendCode?phone=<phone num>
func (mc *MemberController) SendSmsCode(context *gin.Context) {
	//TODO  finish send code
	phone, exist := context.GetQuery("phone")
	if !exist {
		context.JSON(http.StatusOK, map[string]interface{}{
			"code":    0,
			"message": "参数解析失败",
		})
		return
	}

	ms := service.MemberService{}

	isSend := ms.SendCode(phone)
	if isSend {
		context.JSON(http.StatusOK, map[string]interface{}{
			"code": 1,
			"msg":  "发送成功",
		})
		return
	}

	context.JSON(http.StatusOK, map[string]interface{}{
		"code": 0,
		"msg":  "发送失败",
	})

}
