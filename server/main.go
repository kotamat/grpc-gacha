package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "../lib/gacha"
	"math/rand"
	"time"
	"net"
	"log"
	"errors"
)

const (
	port = ":50055"
)

type server struct {}

func (s *server) Lottery(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	if len(in.Cards) < 1 {
		return &pb.Response{Card:nil, RetCode:0}, errors.New("empty cards")
	}
	rand.Seed(time.Now().UnixNano())
	chosenKey := rand.Intn(len(in.Cards))
	return &pb.Response{Card: in.Cards[chosenKey], RetCode: 1}, nil

}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGachaServer(s, &server{})
	s.Serve(lis)
}
