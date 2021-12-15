package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"google.golang.org/grpc"

	"github.com/go/http-rest-api/grpc/proto/genproto/fiboproto"
)

const defaultPort = "11564"

func main() {
	fmt.Println("grpc client")
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	opts := grpc.WithInsecure()
	cc, err := grpc.Dial("localhost:11564", opts)
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()
	client := fiboproto.NewFibonaciApiClient(cc)

	var from int
	fmt.Println("Index from")
	fmt.Scan(&from)

	fmt.Println("Index to")
	var to int
	fmt.Scan(&to)

	fmt.Printf("send request from: %v, to:%v \n", from, to)

	req := fiboproto.SeqRequest{From: int32(from), To: int32(to)}
	resp, err := client.Seq(context.Background(), &req)
	if err != nil {
		log.Fatalf("could not send request: %v", err)
	}

	fmt.Println(resp.Seq)
}
