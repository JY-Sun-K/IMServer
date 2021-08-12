package services

import (
	"context"
	"google.golang.org/grpc"
	"imdemo/imtwo/user/pb"
	"imdemo/imtwo/user/protocol"
	"imdemo/imtwo/user/timeline"
	"io"
	"log"
	"time"
)

var WriteChan chan *protocol.Letter
var WS *timeline.WriteScheduler

func InitWriteChan() error {
	writeChan := make(chan *protocol.Letter,1000)
	WriteChan=writeChan
	conn, err := grpc.Dial(":"+"8888", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc.Dial err: %v", err)
	}


	client:=pb.NewStreamServiceClient(conn)
	stream, err := client.WriteStream(context.Background())
	if err != nil {
		return err
	}
	defer conn.Close()
	defer stream.CloseSend()
	for  {
		msg:=<-WriteChan
		mp:=&pb.MsgPoint{
			AddressIp: msg.IPAddress,
			From:      msg.From,
			To:        msg.To,
			Msg:       msg.Message,
			SendTime:  time.Now().String(),
		}
		r:=&pb.WriteStreamRequest{MP: mp}
		//conn, err := grpc.Dial(":"+"8888", grpc.WithInsecure())
		//if err != nil {
		//	log.Fatalf("grpc.Dial err: %v", err)
		//}
		//
		//
		//client:=pb.NewStreamServiceClient(conn)
		//stream, err := client.WriteStream(context.Background())
		//if err != nil {
		//	return err
		//}

			log.Println("send message:",r)
			err = stream.Send(r)
			if err != nil {
				return err
			}

			_, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				return err
			}
			//stream.CloseSend()


		//conn.Close()
	}

	return nil
}

func InitReceiveWS() {
	ws := timeline.MakeWriteScheduler(2)
	WS=ws
	go WS.Run()
}



