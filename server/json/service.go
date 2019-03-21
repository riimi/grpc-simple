package json

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/riimi/grpc-simple/server/repo"
	"google.golang.org/grpc/grpclog"
	"net/http"
	"os"
	"time"
)

const (
	CODE_SUCCESS = iota + 1
	CODE_REPO
	CODE_INVALID_REQUEST
)

type CountRepo interface {
	Incr(key string) (int, error)
}

type CountService struct {
	Addr   string
	repo   CountRepo
	logger repo.CountLogger
}

func NewCountService(serverAddr, repoAddr string) *CountService {
	l := grpclog.NewLoggerV2(os.Stdout, os.Stdout, os.Stderr)
	return &CountService{
		Addr:   serverAddr,
		repo:   repo.NewCountRepoRedis(repoAddr, l),
		logger: l,
	}
}

func (s *CountService) Run(ctx context.Context) error {
	http.HandleFunc("/v1/count/incr", s.CountHandler)
	s.logger.Fatalf("run: %v", http.ListenAndServe(s.Addr, nil))
	return nil
}

type IncrRequest struct {
	Api string `json:"api"`
	Sid string `json:"sid"`
	Uid string `json:"uid"`
	Key string `json:"key"`
}

type IncrResponse struct {
	Timestamp string `json:"timestamp"`
	Api       string `json:"api"`
	Code      int32  `json:"code"`
	Error     string `json:"error"`
	Count     int32  `json:"count"`

	Picture    string   `json:"picture"`
	Age        int32    `json:"age"`
	Name       string   `json:"name"`
	Gender     string   `json:"gender"`
	Company    string   `json:"company"`
	Email      string   `json:"email"`
	Phone      string   `json:"phone"`
	Address    string   `json:"address"`
	About      string   `json:"about"`
	Registered string   `json:"registered"`
	Latitude   float32  `json:"latitude"`
	Longitude  float32  `json:"longitude"`
	Tags       []string `json:"tags"`
	IsActive   bool     `json:"is_active"`
}

func (s *CountService) CountHandler(w http.ResponseWriter, r *http.Request) {
	req := &IncrRequest{}
	res := &IncrResponse{
		Timestamp: time.Now().String(),
		Api:       "Incr",
	}
	decoder := json.NewDecoder(r.Body)
	encoder := json.NewEncoder(w)
	w.Header().Set("Content-Type", "application/json")

	if err := decoder.Decode(req); err != nil {
		s.logger.Errorf("failed to decode http body: %v", err)
		res.Error = "failed to decode http body"
		res.Code = CODE_INVALID_REQUEST
		if err := encoder.Encode(res); err != nil {
			s.logger.Errorf("failed to write http body: %v", err)
		}
		return
	}

	cnt, err := s.repo.Incr(req.Key)
	if err != nil {
		s.logger.Errorf("s.repo.Incr(%s): %v", req.Key, err)
		res.Code = CODE_REPO
		if err := encoder.Encode(res); err != nil {
			s.logger.Errorf("failed to write http body: %v", err)
		}
		return
	}

	res.Count = int32(cnt)
	res.Code = int32(CODE_SUCCESS)
	res.Error = "none"
	s.logger.Infof("call Incr: %v", res)
	if err := encoder.Encode(res); err != nil {
		s.logger.Errorf("failed to write http body: %v", err)
	}
}

func (s *CountService) Incr(c *gin.Context) {
	req := &IncrRequest{}
	res := &IncrResponse{
		Timestamp: time.Now().String(),
		Api:       "Incr",
	}

	if err := c.Bind(req); err != nil {
		s.logger.Errorf("failed to bind: %v", err)
		res.Error = "failed to decode http body"
		res.Code = CODE_INVALID_REQUEST
		c.JSON(http.StatusOK, res)
		return
	}

	cnt, err := s.repo.Incr(req.Key)
	if err != nil {
		s.logger.Errorf("s.repo.Incr(%s): %v", req.Key, err)
		res.Code = CODE_REPO
		c.JSON(http.StatusOK, res)
		return
	}

	res.Count = int32(cnt)
	res.Code = int32(CODE_SUCCESS)
	res.Error = "none"
	//s.logger.Infof("call Incr: %v", res)
	c.JSON(http.StatusOK, res)
}
