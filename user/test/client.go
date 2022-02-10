package main

import (
	"context"
	"demo/user/conf"
	"flag"
	"fmt"
	"google.golang.org/grpc"

	pb "demo/user/pb"
)

func init() {
	// flags
	flag.StringVar(&conf.GrpcHost, "grpc-host", "0.0.0.0", "grpc服务IP/域名")
	flag.IntVar(&conf.GrpcPort, "grpc-port", 5000, "grpc服务端口")
}

func main() {
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", conf.GrpcHost, conf.GrpcPort), grpc.WithInsecure())
	if err != nil {
		fmt.Printf("创建grpc连接失败! Error: %s", err)
	}
	defer conn.Close()

	create(conn)
	delete(conn)
}

func create(conn *grpc.ClientConn) {
	client := pb.NewUserServiceClient(conn)
	resp, _ := client.Create(context.Background(), &pb.CreateReq{
		Name: "wss",
		Age:  19,
	})

	fmt.Printf("%#v\n", resp)
	fmt.Printf("%#v\n", resp.Data.Name)
}

func delete(conn *grpc.ClientConn) {
	client := pb.NewUserServiceClient(conn)
	resp, _ := client.Delete(context.Background(), &pb.DeleteReq{
		Name: "wss",
		Id:   19,
	})

	fmt.Printf("%#v\n", resp)
	fmt.Printf("%#v\n", resp.Data)
}