package timeline

import (
	"time"
)

//信
type Letter struct {
	IPAddress string //收件人所在服务器ip
	From string //发件人
	To string // 收件人
	Message string //消息
	SendTime time.Time //发送时间
}


//每个人所有的一个Timeline
type TimeLine struct {
	Owner string
	InBox chan *Letter
}



//func (l *Letter) ReadL() {
//	
//}

func MakeLetter(from,to string ,message string)*Letter  {
	return &Letter{
		IPAddress: "127.0.0.1",
		//From:      getOwnerId(),
		From: from,
		To:        to,
		Message:   message,
		SendTime:  time.Now(),
	}
}

func MakeTimeLine(id string) *TimeLine {
	return &TimeLine{
		//Owner: getOwnerId(),
		Owner: id,
		InBox: make(chan *Letter,1000),
	}
}


//获取写信人的id
func getOwnerId() string {
	return ""
}


