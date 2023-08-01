package reqdata


type MsgContent struct {
	ToUsername   string `xml:"ToUserName"`
	FromUsername string `xml:"FromUserName"`
	CreateTime   uint32 `xml:"CreateTime"`
	MsgType      string `xml:"MsgType"`
	Content      string `xml:"Content"`
	Msgid        string `xml:"MsgId"`
	Agentid      uint32 `xml:"AgentId"`
}


type ReplyTextMsg struct {
	ToUsername   string `xml:"ToUserName"`
	FromUsername string `xml:"FromUserName"`
	CreateTime   uint32 `xml:"CreateTime"`
	MsgType      string `xml:"MsgType"`
	Content      string `xml:"Content"`
}