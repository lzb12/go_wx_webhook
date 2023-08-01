package servicer

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"go_wx_webhook/reqdata"
	"go_wx_webhook/wxbizmsgcrypt"
	"net/http"
)






var (
	webhookURL = "" //群通知的机器人
	token = ""  //自建应用的token
	receiverId = ""  //企业微信CorpID
	encodingAeskey = ""
)

func GetWxService(reqMsgSign,reqTimestamp,reqNonce,verifyEchoStr string)  string {

	//reqData := []byte("<xml><ToUserName><![CDATA[wx5823bf96d3bd56c7]]></ToUserName><Encrypt><![CDATA[RypEvHKD8QQKFhvQ6QleEB4J58tiPdvo+rtK1I9qca6aM/wvqnLSV5zEPeusUiX5L5X/0lWfrf0QADHHhGd3QczcdCUpj911L3vg3W/sYYvuJTs3TUUkSUXxaccAS0qhxchrRYt66wiSpGLYL42aM6A8dTT+6k4aSknmPj48kzJs8qLjvd4Xgpue06DOdnLxAUHzM6+kDZ+HMZfJYuR+LtwGc2hgf5gsijff0ekUNXZiqATP7PF5mZxZ3Izoun1s4zG4LUMnvw2r+KqCKIw+3IQH03v+BCA9nMELNqbSf6tiWSrXJB3LAVGUcallcrw8V2t9EL4EhzJWrQUax5wLVMNS0+rUPA3k22Ncx4XXZS9o0MBH27Bo6BpNelZpS+/uh9KsNlY6bHCmJU9p8g7m3fVKn28H3KDYA5Pl/T8Z1ptDAVe0lXdQ2YoyyH2uyPIGHBZZIs2pDBS8R07+qN+E7Q==]]></Encrypt><AgentID><![CDATA[218]]></AgentID></xml>")
	//reqData,err := xml.Marshal(data)
	//if err != nil {
	//	fmt.Println(err)
	//}

	wxcpt := wxbizmsgcrypt.NewWXBizMsgCrypt(token,encodingAeskey,receiverId,wxbizmsgcrypt.XmlType)
	echoStr, cryptErr := wxcpt.VerifyURL(reqMsgSign, reqTimestamp, reqNonce, verifyEchoStr)
	if nil != cryptErr {
		fmt.Println("verifyUrl fail", cryptErr)
		return ""
	}
	fmt.Println("verifyUrl success echoStr", string(echoStr))
	return string(echoStr)


}

func PostWxService(reqMsgSign,reqTimestamp,reqNonce string,data []byte)  {
	wxcpt := wxbizmsgcrypt.NewWXBizMsgCrypt(token,encodingAeskey,receiverId,wxbizmsgcrypt.XmlType)
	msg, cryptErr := wxcpt.DecryptMsg(reqMsgSign, reqTimestamp, reqNonce, data)
	if cryptErr != nil {
		fmt.Println("DecryptMsg fail", cryptErr)
	}
	//fmt.Println("after decrypt msg: ", string(msg))
	var msgContent reqdata.MsgContent
	xmlerr := xml.Unmarshal(msg, &msgContent)
	if xmlerr != nil {
		fmt.Println("Unmarshal fail")
	} else {
		//fmt.Println("struct", msgContent)
		fmt.Println(msgContent.Content)
		if msgContent.Content == "" {   //获取企业微信的内容
			resp,err := http.Post("","",nil) //你要触发的webhook
			if err != nil {
				fmt.Println(err)
				return
			}
			msg := fmt.Sprintf("%s",msgContent.FromUsername) //群机器人通知的内容
			SendWx(msg)
			defer resp.Body.Close()
		}

	}
}

func SendWx(msgContent string) {
	message := map[string]interface{}{
		"msgtype": "text",
		"text": map[string]string{
			"content": msgContent,
		},
	}

	// 将消息转换为JSON格式
	messageData, err := json.Marshal(message)
	if err != nil {
		fmt.Println("Error converting message to JSON:", err)
		return
	}
	// 发起POST请求发送消息
	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(messageData))
	if err != nil {
		fmt.Println("Error occurred:", err)
		return
	}
	defer resp.Body.Close()
	// 检查响应状态码
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Server returned status:", resp.Status)
		return
	}
	fmt.Println("Message sent successfully!")
}
