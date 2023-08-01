package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_wx_webhook/config"
	"go_wx_webhook/routers"
)

func main()  {
	yaml := config.Yaml{}
	yaml.LoadToml()
	r := gin.Default()

	routers.Router.InitApiRouter(r)
	addressBind := fmt.Sprintf("%s:%d", config.Conf.Server.Host ,config.Conf.Server.Port)
	fmt.Println(addressBind)
	r.Run(addressBind)
}
