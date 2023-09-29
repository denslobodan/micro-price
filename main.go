// part 1: all
// part 2: 01:00:
// https://www.youtube.com/watch?v=D0St2LH158Q
// http://localhost:3030/?ticker=ETH
package main

import (
	"flag"
)

func main() {
	var (
		jsonAddr = flag.String("json", ":3030", "listen address of the json transport")
		grpcAddr = flag.String("grpc", ":4040", "listen address of the grpc transport")
	)

	flag.Parse()

	svc := loggingService{priceService{}}

	go makeGRPCServerAndRun(*grpcAddr, svc)

	jsonServer := NewJSONAPIServer(*jsonAddr, svc)
	jsonServer.Run()
}
