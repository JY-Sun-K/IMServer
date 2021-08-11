package timeline

import (
	"log"
	"math/rand"
)

var Users =make(map[string]*TimeLine)
var ErrChan =make(chan string,1000)

func AddUser(id string,timeline *TimeLine) {
	Users[id]=timeline
}


type ReceiveScheduler struct {
	ChanId int //以便于扩建,多个消费者
	AcceptChan chan *Letter//接受信件

}

type WriteScheduler struct {
	WriteChan chan *Letter
	AccChannels []*ReceiveScheduler

}

func Init() {
	ws := MakeWriteScheduler(2)
	go ws.Run()

}

func MakeWriteScheduler(nums int) *WriteScheduler {
	rs:=make([]*ReceiveScheduler,0,nums)
	for i := 0; i < nums; i++ {
		r:=MakeReceiveScheduler(i)

		rs=append(rs,r)
		log.Println("add ReceiveScheduler")
		go r.Run()
	}
	return &WriteScheduler{AccChannels: rs,WriteChan: make(chan *Letter,1000)}
}

func MakeReceiveScheduler(id int)*ReceiveScheduler  {
	log.Println("init ReceiveScheduler")
	return &ReceiveScheduler{
		ChanId:     id,
		AcceptChan: make(chan *Letter,1000),
	}

}

func (w *WriteScheduler) Run() {
	n:=len(w.AccChannels)
	for  {
		msg:= <-w.WriteChan
		//i:=rand.Intn(n)
		//w.AccChannels[i].AcceptChan <- msg
		switch  {
		case msg!=nil:
			i:=rand.Intn(n)
			w.AccChannels[i].AcceptChan <- msg

		}

	}
}

func (r *ReceiveScheduler)Run()  {
	for  {
		msg := <-r.AcceptChan
		timeline:=IsExist(msg.To)
		switch  {
		case timeline==nil:
			ErrChan<- "没有该用户"
		case timeline!=nil:
			timeline.InBox<- msg
		//case : 错误处理


		}
	}
}


//key userid value user of timeline
func IsExist(userID string) *TimeLine {
	timeline,ok:=Users[userID]
	if !ok {
		return nil
	}
	return timeline
}

