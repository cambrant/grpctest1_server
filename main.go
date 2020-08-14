package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/cambrant/grpctest1_pb/go"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedMessageServer
}

func (s *server) SendMessage(ctx context.Context, in *pb.Msg) (*pb.Msg, error) {
	log.Printf("Received: %v", in.GetText())
	return &pb.Msg{Text: "Hello"}, nil
}

func SomeProtobufAttempts() {
	// msg := &pb.Msg{
	// 	Text:   "a bit of text",
	// 	Number: 15,
	// }

	// out, err := proto.Marshal(msg)
	// if err != nil {
	// 	log.Fatalf("Marshal error: %s", err.Error())
	// }

	// msg2 := &pb.Msg{}
	// err = proto.Unmarshal(out, msg2)
	// if err != nil {
	// 	log.Fatalf("Marshal error: %s", err.Error())
	// }

	// fmt.Println("text:", msg2.GetText())
	// fmt.Println("number:", msg2.GetNumber())
	// fmt.Printf("size: %d bytes\n", proto.Size(msg2))
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterMessageServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
