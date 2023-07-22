package main

import (
	"context"
	"log"

	pb "github.com/ayo-ajayi/grpctry/gogen/grpctry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	addr := "localhost:50051"


	//connection server
	conn, err:=grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("failed to connect: ", err)
	}
	defer conn.Close()
	client:=pb.NewGreeterClient(conn)

	//contact server
	res, err := client.SayHello(context.Background(), &pb.HelloRequest{Name: "Ayo"})
	if err != nil {
		log.Fatalln("could not say hello: ", err)
	}
	log.Println("Message is: ", res.GetMessage())
}