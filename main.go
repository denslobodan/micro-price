// 1:12
// https://www.youtube.com/watch?v=367qYRy39zw&t=3162s
// http://localhost:3030/?ticker=ETH
package main

import (
	"flag"
)

func main() {
	/* client := client.New("http://localhost:3030")

	price, err := client.FetchPrice(context.Background(), "ETH")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", price)
	return
	*/
	listenAddr := flag.String("listenAddr", ":3030", "listen address the srvice is running")
	svc := NewLoggingService(NewMetricSevice(&priceFetch{}))

	server := NewJSONAPIServer(*listenAddr, svc)
	server.Run()
}
