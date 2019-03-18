package main

import (
	"context"
	"fmt"
	"log"
	pbBounce "relay/reports/bounce/proto/bounce"
	"time"

	"google.golang.org/grpc"
)

const (
	address = "localhost:8080"
)

//Dummy client for testing
func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect to server  %v", err)
	}
	defer conn.Close()

	c := pbBounce.NewBounceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	resp, err := c.GetData(ctx, &pbBounce.BounceGetData{AccountId: 1})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("this is bounce response", resp)

}
