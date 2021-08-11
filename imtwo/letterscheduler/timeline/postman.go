package timeline

import (
	"context"
	"google.golang.org/grpc"
	"imdemo/imtwo/letterscheduler/pb"
	"log"
	"math/rand"
	"strings"
)

var Users =make(map[int64]*TimeLine)
var ErrChan =make(chan string,1000)

func AddUser(id int64,timeline *TimeLine) {
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
		log.Println(msg)
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
		//client
		//timeline:=IsExist(msg.To)
		//switch  {
		//case timeline==nil:
		//	ErrChan<- "没有该用户"
		//case timeline!=nil:
		//	timeline.InBox<- msg
		////case : 错误处理
		mp:=&pb.Msg{
			AddressIp: msg.IPAddress,
			From:      msg.From,
			To:        msg.To,
			Msg:       msg.Message,
			SendTime:  msg.SendTime,
		}
		address :=strings.Split(msg.IPAddress,":")
		conn, err := grpc.Dial(":"+address[1], grpc.WithInsecure())
		if err != nil {
			log.Fatalf("grpc.Dial err: %v", err)
		}
		//todo
		client:=pb.NewUserServiceClient(conn)
		resp,err:=client.ReceiveMsg(context.Background(),&pb.ReceiveMsgRequest{MP: mp})
		if err != nil {
			log.Fatalf("client.Search err: %v", err)
		}

		log.Printf("resp: %s", resp)

		conn.Close()
	}

}


//key userid value user of timeline
func IsExist(userID int64) *TimeLine {
	timeline,ok:=Users[userID]
	if !ok {
		return nil
	}
	return timeline
}

