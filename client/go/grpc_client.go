package main

import (
	"context"
	"github.com/riimi/grpc-simple/protocol"
	"google.golang.org/grpc"
	"log"
)

type GrpcClient struct {
	conn *grpc.ClientConn
}

func NewGrpcClient(addr string) *GrpcClient {
	opts := []grpc.DialOption{grpc.WithInsecure()}
	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		log.Fatal(err)
	}

	return &GrpcClient{
		conn: conn,
	}
}

func (g *GrpcClient) Data() interface{} {
	return g.conn
}

func (g *GrpcClient) LoadTest() func(context.Context, interface{}) {
	return func(ctx context.Context, x interface{}) {
		conn := x.(*grpc.ClientConn)
		client := protocol.NewCountServiceClient(conn)
		_, err := client.Incr(ctx, &protocol.IncrRequest{
			Api: "Incr",
			Sid: "nadana",
			Uid: "123132",
			Key: "myspoon",
		})
		if err != nil {
			log.Fatal(err)
		}
		//log.Print(res)
	}
}
