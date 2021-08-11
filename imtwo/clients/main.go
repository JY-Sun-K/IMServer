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
	requestToken := new(utils.AuthToken)
	requestToken.Token=r.Token

	log.Println(requestToken.Token)

	conn, err = grpc.Dial(":8080", grpc.WithInsecure(),grpc.WithPerRPCCredentials(requestToken))
	if err != nil {
		grpclog.Errorf("failed to connect: %v", err)
	}
	defer conn.Close()
	client=pb.NewUserServiceClient(conn)
	userInfo,err:=client.GetUserInfo(context.Background(),&pb.GetUserInfoRequest{
		Email: "654321@qq.com",
		Token: requestToken.Token,
	})
	if err != nil {
		grpclog.Errorln("get user info failed:",err)
	}
	log.Println(userInfo)

	conn, err = grpc.Dial(":8080", grpc.WithInsecure(),grpc.WithPerRPCCredentials(requestToken))
	if err != nil {
		grpclog.Errorf("failed to connect: %v", err)
	}
	defer conn.Close()
	client=pb.NewUserServiceClient(conn)
	SUser,err:=client.SearchUser(context.Background(),&pb.SearchUserRequest{UserName: "432"})
	if err != nil {
		grpclog.Errorln("get user info failed:",err)
	}
	log.Println(SUser)

	conn, err = grpc.Dial(":8080", grpc.WithInsecure(),grpc.WithPerRPCCredentials(requestToken))
	if err != nil {
		grpclog.Errorf("failed to connect: %v", err)
	}
	defer conn.Close()
	client=pb.NewUserServiceClient(conn)
	AUser,err:=client.AddUser(context.Background(),&pb.AddUserRequest{UserId: 4})
	if err != nil {
		grpclog.Errorln("get user info failed:",err)
	}
	log.Println(AUser)
}
