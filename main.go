// part 1: all
// part 2: 01:20:
// https://www.youtube.com/watch?v=D0St2LH158Q
// http://localhost:3030/?ticker=ETH
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"pricefetcher/client"
	"pricefetcher/proto"
	"time"
)

func main() {
	var (
		jsonAddr = flag.String("json", ":3030", "listen address of the json transport")
		grpcAddr = flag.String("grpc", ":4040", "listen address of the grpc transport")
		svc      = loggingService{priceService{}}
		ctx      = context.Background()
	)

	flag.Parse()

	grpcClient, err := client.NewGRPCClient(":4040")
	if err != nil {
		log.Fatal(err)
	}

	// ticker := time.NewTicker()
	go func() {
		for {
			time.Sleep(3 * time.Second)
			resp, err := grpcClient.FetchPrice(ctx, &proto.PriceRequest{Ticker: "BTC"})
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("%+v\n", resp)
		}
	}()

	go makeGRPCServerAndRun(*grpcAddr, svc)

	jsonServer := NewJSONAPIServer(*jsonAddr, svc)
	jsonServer.Run()
}
