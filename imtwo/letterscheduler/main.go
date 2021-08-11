package main

import (
	"google.golang.org/grpc"
	"imdemo/imtwo/letterscheduler/pb"
	"imdemo/imtwo/letterscheduler/services"
	"log"
	"net"
)


//:8888
func main() {
	services.InitWs()

	server := grpc.NewServer()
	pb.RegisterStreamServiceServer(server,&services.StreamService{})

	lis, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}

	server.Serve(lis)
}
