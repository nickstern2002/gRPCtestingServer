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

	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
	_ "google.golang.org/protobuf/runtime/protoimpl"
)

var (
	port = flag.String("port", "8080", "Port to listen on")
)

type server struct {
	v1.UnimplementedJunkyardServiceServer
}

// MakeMyDayBetter returns a string depedning on the value given
func (s *server) MakeMyDayBetter(_ context.Context, in *v1.MakeMyDayBetterRequest) (*v1.MakeMyDayBetterResponse, error) {
	resp := &v1.MakeMyDayBetterResponse{}
	mood := in.GetMood()
	switch mood {
	case 0:
		fmt.Println("Mood 0")
	case 1:
		fmt.Println("Mood 1")
	case 2:
		fmt.Println("Mood 2")
	case 3:
		fmt.Println("Mood 3")
	default:
		resp.Message = fmt.Sprintf("Invalid mood, mood must be [0,3]")
		log.Printf("Invalid mood")
		return nil, errors.New("Invalid mood")
	}
	log.Printf("Valid mood")
	resp.Message = fmt.Sprintf("Mood %d", mood)
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
