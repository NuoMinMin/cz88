package main

import (
	"fmt"

	"cz88/internal/server/http"
	"cz88/internal/server/rpc"

	"cz88/internal/service"
)

func init() {
	fmt.Printf("IP num: %d\n", service.GetMetricsNum().Data)
}

func main() {
	go rpc.New()
	go http.New()
	select {}
}
