package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"imdemo/imtwo/user/pb"
	"imdemo/imtwo/user/utils"
	"log"
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
	//r,err:=client.Register(context.Background(),&pb.RegisterRequest{
	//	UserName: "654321",
	//	Email:    "654321@qq.com",
	//	Password: "654321",
	//})
	log.Println(r)
	if err != nil {
		grpclog.Errorln("login failed:",err)
	}

	conn, err = grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		grpclog.Errorf("failed to connect: %v", err)
	}
	defer conn.Close()
	client=pb.NewUserServiceClient(conn)

	w,err:=client.SendMsg(context.Background(),&pb.SendMsgRequest{MP: &pb.Msg{
		AddressIp: "127.0.0.1:8080",
		From:      1,
		To:        4,
		Msg:       "hello",
		SendTime:  "",
	}})
	log.Println("sss")
	log.Println(w)

	conn, err = grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		grpclog.Errorf("failed to connect: %v", err)
	}
	defer conn.Close()
	client=pb.NewUserServiceClient(conn)
	r,err=client.Login(context.Background(),&pb.LoginRequest{
		Email:    "123456@qq.com",
		Password: "123456",

	})
	//r,err:=client.Register(context.Background(),&pb.RegisterRequest{
	//	UserName: "654321",
	//	Email:    "654321@qq.com",
	//	Password: "654321",
	//})
	log.Println(r)
	if err != nil {
		grpclog.Errorln("login failed:",err)
	}


	requestToken := new(utils.AuthToken)
	requestToken.Token=r.Token

	log.Println(requestToken.Token)
}
