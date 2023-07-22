package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	pb "github.com/ayo-ajayi/grpctry/gogen/grpctry"
	"google.golang.org/grpc"
)

func main() {

	//create new grpc server
	server := grpc.NewServer()

	//register GreaterServer with grpc server
	pb.RegisterGreeterServer(server, &GreeterServer{
		i: 0,
	})

	//listening on port
	list, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalln("failed to listen: ", err)
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		//serve grpc requests
		log.Println("starting grpc server")
		if err := server.Serve(list); err != nil {
			log.Fatalln("failed to serve: ", err)
		}		
	}()
	

	<-stop
	func() {
		server.GracefulStop()
		log.Println("gracefully stopped grpc server")
	}()

}

type GreeterServer struct {
	i int64
	pb.UnimplementedGreeterServer
}

func (g *GreeterServer) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	i := &g.i
	for (*i) < 2000 {
		(*i)++
	}
	reqName := req.Name
	res := &pb.HelloReply{
		Message: fmt.Sprintf("succcessfully written grpc response for req: %v\n i: %v", reqName, *i),
	}
	return res, nil
}
