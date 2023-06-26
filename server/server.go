package server

import (
	"context"
	"errors"
	"github.com/smallnest/rpcx/protocol"
	"github.com/smallnest/rpcx/server"
	"golang.org/x/time/rate"
)

type singlelimiter struct {
	Name    string
	Limit   float64
	Cap     int
	Limiter *rate.Limiter
}

type LimitService struct {
	port     string
	limiters []singlelimiter
	//limiter *rate.Limiter
}

func NewLimitService(port string) *LimitService {
	return &LimitService{
		port: port,
	}

}

// the second parameter is not a pointer
func (t *LimitService) Run() error {
	s := server.NewServer()
	err := s.RegisterName("Limit", t, "")
	if err != nil {
		panic(err)
	}
	s.AuthFunc = auth
	err = s.Serve("tcp", "0.0.0.0:"+t.port)
	if err != nil {
		panic(err)
	}
	return nil
}

func (l *LimitService) NewSpeedLimiter(limit float64, cap int, name string) {
	l.limiters = append(l.limiters, singlelimiter{
		Name:    name,
		Limit:   limit,
		Cap:     cap,
		Limiter: rate.NewLimiter(rate.Limit(limit), cap),
	})
}

func auth(ctx context.Context, req *protocol.Message, token string) error {

	if token == "bearer tGzv3JOkF0XG5Qx2TlKWIA" {
		return nil
	}

	return errors.New("invalid token")
}
