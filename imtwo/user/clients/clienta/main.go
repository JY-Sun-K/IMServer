package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"imdemo/imtwo/user/pb"
	"io"
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
		Email:    "654321@qq.com",
		Password: "654321",

	})
	log.Println(r)
	conn1, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		grpclog.Errorf("failed to connect: %v", err)
	}
	defer conn1.Close()
	client1:=pb.NewUserServiceClient(conn1)
	stream,err:=client1.AcceptMsg(context.Background(),&pb.AcceptMsgRequest{UserId: r.UserId})
	if err != nil {
		grpclog.Errorln("login failed:",err)
	}
	for  {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}
		log.Println(resp)
	}
	

	if err != nil {
		grpclog.Errorln("login failed:",err)
	}

}