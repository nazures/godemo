package main

import (
	"context"
	"fmt"
	pb "godemo/rpcdemo/grpc/grpc_crt/proto"
	"google.golang.org/grpc"
	"net"
)

const (
	port = ":8848"
)

type server struct {
	pb.UnimplementedGreeterServer
}

// 服务端实现方法
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	fmt.Printf("服务端收到请求：%v\n", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	// 设置监听
	lis, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Printf("create listen failed, err: %v\n", err)
		return
	}
	// 得到一个服务器
	s := grpc.NewServer()
	// 将服务注册给服务器
	pb.RegisterGreeterServer(s, &server{})

	err = s.Serve(lis)
	if err != nil {
		fmt.Printf("failed to serve, err: %v\n", err)
		return
	}
}
