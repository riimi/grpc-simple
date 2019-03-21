package main

import (
	"context"
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
	"github.com/riimi/grpc-simple/server/grpc"
	"github.com/riimi/grpc-simple/server/json"
	"os"
	"os/signal"
	"runtime/pprof"
)

func main() {
	serverAddr := flag.String("server", "127.0.0.1:40040", "server address")
	repoAddr := flag.String("repo", "127.0.0.1:6379", "repo address")
	loglevel := flag.Int("log", 4, "log level")
	mode := flag.String("mode", "grpc", "mode")
	cpuprofile := flag.String("cpu", "", "path to cpu profile")
	memporfile := flag.String("mem", "", "path to mem profile")
	flag.Parse()

	if *cpuprofile != "" {
		cf, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		defer cf.Close()
		pprof.StartCPUProfile(cf)
	}

	switch *mode {
	case "grpc":
		srv := grpc.NewCountService(*serverAddr, *repoAddr, *loglevel)
		go func() {
			log.Print(srv.Run(context.Background()))
		}()
	case "json":
		srv := json.NewCountService(*serverAddr, *repoAddr)
		r := gin.New()
		r.POST("/v1/count/incr", srv.Incr)

		go func() {
			log.Print(r.Run(srv.Addr))
		}()
	default:
		log.Fatal("unsupported mode")
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	if *cpuprofile != "" {
		pprof.StopCPUProfile()
	}
	if *memporfile != "" {
		mf, err := os.Create(*memporfile)
		if err != nil {
			log.Fatal(err)
		}
		defer mf.Close()
		if err := pprof.WriteHeapProfile(mf); err != nil {
			log.Fatal(err)
		}
	}
}
