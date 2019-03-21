package grpc

import (
	"context"
	"github.com/golang/protobuf/ptypes"
	"github.com/riimi/grpc-simple/protocol"
	"github.com/riimi/grpc-simple/server/repo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
)

//go:generate protoc -I../protocol -I/usr/local/include -I$GOPATH/src -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc:../protocol --grpc-gateway_out=logtostderr=true,grpc_api_configuration=../protocol/count_service.yaml:../protocol --swagger_out=logtostderr=true,grpc_api_configuration=../protocol/count_service.yaml:../protocol --csharp_out=../protocol --grpc_out=../protocol --plugin=protoc-gen-grpc=/root/.nuget/packages/grpc.tools/1.19.0/tools/linux_x64/grpc_csharp_plugin count.proto
const (
	LOGLEVEL_FATAL = iota + 1
	LOGLEVEL_ERROR
	LOGLEVEL_WARNING
	LOGLEVEL_INFO

	CODE_SUCCESS
	CODE_REPO
)

type CountRepo interface {
	Incr(key string) (int, error)
}

type CountService struct {
	addr   string
	repo   CountRepo
	logger repo.CountLogger
}

func NewCountService(serverAddr, repoAddr string, level int) *CountService {
	l := grpclog.NewLoggerV2WithVerbosity(os.Stdout, os.Stdout, os.Stderr, level)
	return &CountService{
		addr:   serverAddr,
		repo:   repo.NewCountRepoRedis(repoAddr, l),
		logger: l,
	}
}

func (s *CountService) SetLogger(l repo.CountLogger) {
	s.logger = l
}

func (s *CountService) Run(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	protocol.RegisterCountServiceServer(grpcServer, s)

	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		s.logger.Fatalf("failed to listen: %v", err)
	}

	go func() {
		s.logger.Infof("server is running: %s", s.addr)
		s.logger.Errorf("serve: %v", grpcServer.Serve(lis))
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	cancel()
	grpcServer.GracefulStop()
	return nil
}

func (s *CountService) grpcHandlerFunc(grpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		buf := make([]byte, 1024)
		body := req.Body
		body.Read(buf)
		if s.logger.V(LOGLEVEL_INFO) {
			s.logger.Infof("proto: %d, type: %s", req.ProtoMajor, string(buf))
		}
		if req.ProtoMajor == 2 && strings.HasPrefix(
			req.Header.Get("Content-Type"), "application/grpc") {
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
		if s.logger.V(LOGLEVEL_ERROR) {
			s.logger.Errorf("s.repo.Incr(%s): %v", req.Key, err)
		}
		res.Code = CODE_REPO
		res.Error = "failed to incr in repo"
		return res, nil
	}

	res.Count = int32(cnt)
	res.Code = int32(CODE_SUCCESS)
	res.Error = "none"
	if s.logger.V(LOGLEVEL_INFO) {
		s.logger.Infof("call Incr: %v", res)
	}
	return res, nil
}
