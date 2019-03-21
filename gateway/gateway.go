package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/riimi/grpc-simple/protocol"
	"google.golang.org/grpc"
	"log"
	"net/http"
)

func main() {
	EndPoint := flag.String("endpoint", "localhost:40040", "endpoint of chatserver")
	port := flag.Int("port", 8081, "gateway port")
	flag.Parse()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	if err := protocol.RegisterCountServiceHandlerFromEndpoint(ctx, mux, *EndPoint, opts); err != nil {
		log.Fatal(err)
	}
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), mux))
}
