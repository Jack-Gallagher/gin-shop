package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HelloController struct {
}

func (hello *HelloController) Router(eng *gin.Engine) {
	eng.GET("/hello", hello.Hello)
}
func (hello *HelloController) Hello(context *gin.Context) {
	context.JSON(http.StatusOK, map[string]interface{}{
		"message": "hello welcome",
	})

}
