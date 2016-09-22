package main

import (
	"google.golang.org/grpc"
	"log"
	pb "../lib/gacha"
	"golang.org/x/net/context"
)

const (
	address = "localhost:50055"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGachaClient(conn)

	// define cards
	cards := []*pb.Card{
		&pb.Card{Name: "card1"},
		&pb.Card{Name: "card2"},
	}

	r, err := c.Lottery(context.Background(), &pb.Request{Cards: cards})
	if err != nil {
		log.Fatalf("could not get card %v", err)
	}
	log.Printf("gain card: %v", r.Card.Name)
}