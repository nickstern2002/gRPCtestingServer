package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	v1 "github.com/nickstern2002/gRPCtestingServer/pkg/protogen/compute/v1"
	"google.golang.org/grpc"
	"log"
	"net"
)

var (
	port = flag.String("port", "8080", "Port to listen on")
)

type server struct {
	v1.JunkyardServiceServer
}

func MakeMyDayBetter(_ context.Context, in *v1.MakeMyDayBetterRequest) (*v1.MakeMyDayBetterResponse, error) {
	switch mood := in.GetMood(); mood {
	case 0:
		fmt.Println("Mood 0, L bozo")
	case 1:
		fmt.Println("Mood 1, L bozo")
	case 2:
		fmt.Println("Mood 2, L bozo")
	case 3:
		fmt.Println("Mood 3, L bozo")
	default:
		fmt.Println("Invalid mood, mood must be [0,3]")
		return nil, errors.New("Invalid mood")
	}
	resp := &v1.MakeMyDayBetterResponse{
		Message: "oof",
	}
	return resp, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	v1.RegisterJunkyardServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

/*
conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

*/
