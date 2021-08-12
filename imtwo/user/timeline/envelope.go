package timeline

import (
	"imdemo/imtwo/user/protocol"
	"time"
)

////信
//type Letter struct {
//	IPAddress string //收件人所在服务器ip
//	From int64 //发件人
//	To int64 // 收件人
//	Message string //消息
//	SendTime string //发送时间
//}
//
//
////每个人所有的一个Timeline
//type TimeLine struct {
//	Owner int64
//	InBox chan *Letter
//}
//


//func (l *Letter) ReadL() {
//	
//}

func MakeLetter(from,to int64 ,message string)*protocol.Letter  {
	return &protocol.Letter{
		IPAddress: "127.0.0.1",
		//From:      getOwnerId(),
		From: from,
		To:        to,
		Message:   message,
		SendTime:  time.Now().String(),
	}
}

func MakeTimeLine(id int64) *protocol.TimeLine {
	return &protocol.TimeLine{
		//Owner: getOwnerId(),
		Owner: id,
		InBox: make(chan *protocol.Letter,1000),
	}
}


//获取写信人的id
func getOwnerId() string {
	return ""
}


