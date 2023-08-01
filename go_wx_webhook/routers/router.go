package routers

import (
	"github.com/gin-gonic/gin"
	"go_wx_webhook/controller"
)
var Router router


type router struct {

}

func (r *router) InitApiRouter(router *gin.Engine) {

	//shell
	wx := router.Group("/api/asdasdadadadssadasdasd/")

	wx.GET("wxpush",controller.GetWxController)
	wx.POST("wxpush",controller.PostWxController)

}
