package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"imdemo/imtwo/user/pb"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		grpclog.Errorf("failed to connect: %v", err)
	}
	defer conn.Close()
	client:=pb.NewUserServiceClient(conn)
	r,err:=client.Login(context.Background(),&pb.LoginRequest{
		Email:    "123456@qq.com",
		Password: "123456",

	})
	log.Println(r)
	//r,err:=client.Register(context.Background(),&pb.RegisterRequest{
	//	UserName: "654321",
	//	Email:    "654321@qq.com",
	//	Password: "654321",
	//})
	conn1, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		grpclog.Errorf("failed to connect: %v", err)
	}
	defer conn1.Close()
	client1:=pb.NewUserServiceClient(conn1)
	for i := 0; i < 20; i++ {
		r1,err:=client1.SendMsg(context.Background(),&pb.SendMsgRequest{MP: &pb.Msg{
			AddressIp: "127.0.0.1:8080",
			From:      1,
			To:        4,
			Msg:       "hello",
			SendTime:  time.Now().String(),
		}})
		log.Println(r1)
		if err != nil {
			grpclog.Errorln("login failed:",err)
		}
	}


}