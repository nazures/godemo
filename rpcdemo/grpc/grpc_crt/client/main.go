package main

import (
	"context"
	"fmt"
	pb "godemo/rpcdemo/grpc/grpc_crt/proto"
	"google.golang.org/grpc"
	"time"
)

const (
	address = "localhost:8848"
)

func main() {
	// 建立与服务器的连接
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		fmt.Printf("did not connect, err: %v\n", err)
		return
	}
	defer conn.Close()

	// 创建连接客户端
	client := pb.NewGreeterClient(conn)
	// 得到一个ctx
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	// 远程调用
	name := "小王子"
	reply, err := client.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		fmt.Printf("greet failed, err: %v\n", err)
		return
	}
	fmt.Printf("Greet: %s\n", reply.GetMessage())

}
