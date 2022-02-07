package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"

	"demo/conf"
	pb "demo/user/pb"
)

func init() {
	// flags
	flag.StringVar(&conf.GrpcAddr, "grpc-addr", conf.GetEnv("GrpcAddr", "0.0.0.0:5000"), "grpc服务地址")
}

func main() {
	conn, err := grpc.Dial(conf.GrpcAddr, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("创建grpc连接失败! Error: %s", err)
	}
	defer conn.Close()

	Create(conn)
	Delete(conn)
}

func Create(conn *grpc.ClientConn) {
	client := pb.NewUserServiceClient(conn)
	resp, _ := client.Create(context.Background(), &pb.CreateReq{
		Name: "wss",
		Age:  19,
	})

	fmt.Printf("%#v\n", resp)
	fmt.Printf("%#v\n", resp.Data.Name)
}

func Delete(conn *grpc.ClientConn) {
	client := pb.NewUserServiceClient(conn)
	resp, _ := client.Delete(context.Background(), &pb.DeleteReq{
		Name: "wss",
		Id:   19,
	})

	fmt.Printf("%#v\n", resp)
	fmt.Printf("%#v\n", resp.Data)
}