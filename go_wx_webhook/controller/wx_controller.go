package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_wx_webhook/servicer"
	"io/ioutil"
	"net/http"
)

func GetWxController(c *gin.Context)  {
	verifyMsgSign := c.Query("msg_signature")
	verifyTimestamp := c.Query("timestamp")
	verifyNonce := c.Query("nonce")
	verifyEchoStr := c.Query("echostr")



	echoStr := servicer.GetWxService(verifyMsgSign,verifyTimestamp,verifyNonce,verifyEchoStr)
	c.String(http.StatusOK,echoStr)
}

func PostWxController(c *gin.Context)  {
	reqMsgSign := c.Query("msg_signature")
	reqTimestamp := c.Query("timestamp")
	reqNonce := c.Query("nonce")
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println(err)
	}
	servicer.PostWxService(reqMsgSign,reqTimestamp,reqNonce,body)

}