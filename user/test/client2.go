package main

import (
	"context"
	pb "demo/user/pb"
	"fmt"
	"google.golang.org/grpc/credentials/insecure"

	_ "github.com/mbobakov/grpc-consul-resolver" // It's important
	"google.golang.org/grpc"
)

func main() {
	// 从注册中心获取server地址进行连接
	userConn, err := grpc.Dial(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s", "consul", 8500, "user-srv"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`), // 利用grpc负载均衡
	)
	if err != nil {
		panic(err)
	}

	create2(userConn)
	delete2(userConn)
}

func create2(conn *grpc.ClientConn)  {
	client := pb.NewUserServiceClient(conn)
	resp, err := client.Create(context.Background(), &pb.CreateReq{
		Name: "xiaoming",
		Age: 18,
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v\n", resp)
	fmt.Printf("%#v\n", resp.Data.Name)
}

func delete2(conn *grpc.ClientConn) {
	client := pb.NewUserServiceClient(conn)
	resp, _ := client.Delete(context.Background(), &pb.DeleteReq{
		Name: "wss",
		Id:   19,
	})

	fmt.Printf("%#v\n", resp)
	fmt.Printf("%#v\n", resp.Data)
}