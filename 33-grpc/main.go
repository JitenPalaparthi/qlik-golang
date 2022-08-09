package main

import (
	pb "demo/protos"
	"fmt"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedHelloServiceServer
}

func (s *server) SayHello(in *pb.HelloRequest, stream pb.HelloService_SayHelloServer) error {
	for i := 0; i <= 5; i++ {
		out := &pb.HelloResponse{Reply: "Hello-->" + fmt.Sprint(i)}
		if err := stream.Send(out); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":50082")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()

	pb.RegisterHelloServiceServer(srv, &server{})

	err = srv.Serve(lis)
	if err != nil {
		panic(err)
	}
}
