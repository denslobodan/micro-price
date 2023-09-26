// part 1: all
// part 2: 00:00
// https://www.youtube.com/watch?v=D0St2LH158Q
// http://localhost:3030/?ticker=ETH
package main

import (
	"flag"
)

func main() {
	port := flag.String("port", ":3030", "listen address the service is running")
	flag.Parse()

	svc := loggingService{priceService{}}
	server := NewJSONAPIServer(*port, svc)
	server.Run()
}
