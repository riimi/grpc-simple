package main

import (
	"context"
	"flag"
	"github.com/labstack/gommon/log"
)

func main() {
	serverAddr := flag.String("server", "127.0.0.1:40040", "server address")
	repoAddr := flag.String("repo", "127.0.0.1:6379", "repo address")
	flag.Parse()

	srv := NewCountService(*serverAddr, *repoAddr)
	log.Fatal(srv.Run(context.Background()))
}
