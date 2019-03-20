package main

import (
	"context"
	"github.com/golang/protobuf/ptypes"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/riimi/grpc-simple/protocol"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
)

//go:generate protoc -I../protocol -I/usr/local/include -I$GOPATH/src -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc:../protocol --grpc-gateway_out=logtostderr=true,grpc_api_configuration=../protocol/count_service.yaml:../protocol --swagger_out=logtostderr=true,grpc_api_configuration=../protocol/count_service.yaml:../protocol --csharp_out=../protocol --grpc_out=../protocol --plugin=protoc-gen-grpc=/root/.nuget/packages/grpc.tools/1.19.0/tools/linux_x64/grpc_csharp_plugin count.proto
const (
	CODE_SUCCESS = iota
	CODE_REPO
)

type CountLogger interface {
	Fatalf(format string, a ...interface{})
	Warningf(foramt string, a ...interface{})
	Infof(format string, a ...interface{})
	Errorf(format string, a ...interface{})
}

type CountRepo interface {
	Incr(key string) (int, error)
}

type CountService struct {
	addr   string
	repo   CountRepo
	logger CountLogger
}

func NewCountService(serverAddr, repoAddr string) *CountService {
	l := grpclog.NewLoggerV2(os.Stdout, os.Stdout, os.Stderr)
	return &CountService{
		addr:   serverAddr,
		repo:   NewCountRepoRedis(repoAddr, l),
		logger: l,
	}
}

func (s *CountService) SetLogger(l CountLogger) {
	s.logger = l
}

func (s *CountService) Run(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	protocol.RegisterCountServiceServer(grpcServer, s)

	mux := http.NewServeMux()
	mux.HandleFunc("/swagger.json", func(w http.ResponseWriter, req *http.Request) {
		content, err := ioutil.ReadFile("../protocol/count.swagger.json")
		if err != nil {
			s.logger.Errorf("failed to read swagger file: %v", err)
			http.Error(w, "failed to read swagger file", http.StatusNotFound)
		}
		if _, err := w.Write(content); err != nil {
			s.logger.Errorf("failed to write: %v", err)
		}
	})

	dopts := []grpc.DialOption{grpc.WithInsecure()}
	gwmux := runtime.NewServeMux()
	if err := protocol.RegisterCountServiceHandlerFromEndpoint(ctx, gwmux, s.addr, dopts); err != nil {
		s.logger.Fatalf("failed to serve gateway: %v", err)
	}
	mux.Handle("/", gwmux)

	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		s.logger.Fatalf("failed to listen: %v", err)
	}

	server := &http.Server{
		Addr:    s.addr,
		Handler: grpcHandlerFunc(grpcServer, mux),
	}

	go func() {
		s.logger.Infof("server is running: %s", s.addr)
		s.logger.Errorf("serve: %v", server.Serve(lis))
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	cancel()
	grpcServer.GracefulStop()
	return nil
}

func grpcHandlerFunc(grpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if req.ProtoMajor == 2 && strings.Contains(req.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, req)
		} else {
			otherHandler.ServeHTTP(w, req)
		}
	})
}

func (s *CountService) Incr(ctx context.Context, req *protocol.IncrRequest) (*protocol.IncrResponse, error) {
	res := &protocol.IncrResponse{
		Timestamp: ptypes.TimestampNow(),
		Api:       "Incr",
	}
	cnt, err := s.repo.Incr(req.Key)
	if err != nil {
		s.logger.Errorf("s.repo.Incr(%s): %v", req.Key, err)
		res.Code = CODE_REPO
		res.Error = "failed to incr in repo"
		return res, nil
	}

	res.Count = int32(cnt)
	res.Code = CODE_SUCCESS
	return res, nil
}
