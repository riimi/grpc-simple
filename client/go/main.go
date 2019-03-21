package main

import (
	"context"
	"flag"
	worker2 "github.com/riimi/grpc-simple/client/worker"
	"log"
	"time"
)

type Client interface {
	LoadTest() func(ctx context.Context, x interface{})
	Data() interface{}
}

type IncrRequest struct {
	Api string `json:"api"`
	Sid string `json:"sid"`
	Uid string `json:"uid"`
	Key string `json:"key"`
}

func main() {
	addr := flag.String("server", "127.0.0.1:40040", "server address")
	worker := flag.Int("worker", 100, "the number of worker")
	try := flag.Int("try", 1000, "the number of tries")
	mode := flag.String("mode", "grpc", "mode")
	flag.Parse()

	var client Client
	if *mode == "grpc" {
		client = NewGrpcClient(*addr)
	} else if *mode == "json" {
		client = NewJsonClient(*addr, *worker)
	} else {
		log.Fatalf("unsupported mode")
	}

	d := worker2.NewDispatcher(*worker)
	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()
	d.SetHandler(client.LoadTest())

	for i := 0; i < *try; i++ {
		d.Add(client.Data())
	}
	start := time.Now()
	d.Run(ctx)
	if err := ctx.Err(); err != nil {
		log.Fatal(err)
	}
	d.Wait()
	log.Print(time.Since(start).Seconds())
}
