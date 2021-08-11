package main

import (
	"google.golang.org/grpc"
	"imdemo/imtwo/user/dao"
	"imdemo/imtwo/user/pb"
	"imdemo/imtwo/user/services"
	"log"
	"net"
)

func main() {
	err:= dao.InitDB()
	if err != nil {
		panic(err)
	}
	lis, err := net.Listen("tcp",":8080")
	if err != nil{
		log.Fatalf("failed to listen :%v",err)
	}
	server := grpc.NewServer()

	pb.RegisterUserServiceServer(server,&services.UserService{UserDAO: &dao.UserDAOImpl{}})
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
