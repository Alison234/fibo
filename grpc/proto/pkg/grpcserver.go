package adder

import (
	"context"
	"net"
	"time"

	"github.com/go/http-rest-api/cache"
	"github.com/go/http-rest-api/fibonacci"
	"github.com/go/http-rest-api/grpc/proto/genproto/fiboproto"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type GRPCServerConfig struct {
	Host        string
	Port        string
	ConnTimeout time.Duration
}

type server struct {
	fiboproto.FibonaciApiServer
	fibonaciProvider *fibonacci.FibonacciProvider
	Cache            cache.Cache
	Logger           *logrus.Logger
}

func NewGRPCServer() *server {
	cacheStore := cache.NewMemCacher("localhost:11211")
	return &server{fibonaciProvider: fibonacci.NewFibonacciProvider(), Cache: cacheStore, Logger: logrus.New()}
}

func (serv *server) Start(cfg *GRPCServerConfig) {
	conn, err := net.Listen("tcp", net.JoinHostPort(cfg.Host, cfg.Port))
	if err != nil {
		serv.Logger.Fatalln(err)
	}
	server := grpc.NewServer(grpc.ConnectionTimeout(cfg.ConnTimeout))
	fiboproto.RegisterFibonaciApiServer(server, serv)

	serv.Logger.Infof("grpc server starting and listen %v:%v", cfg.Host, cfg.Port)

	if err := server.Serve(conn); err != nil {
		serv.Logger.Info(err)
	}
}

func (serv *server) Seq(ctx context.Context, req *fiboproto.SeqRequest) (*fiboproto.SeqResponse, error) {
	from := int(req.From)
	to := int(req.To)

	err := serv.fibonaciProvider.Calculate(from, to)
	if err != nil {
		return nil, err
	}

	result := make([]*fiboproto.Fib, 0)

	for _, f := range serv.fibonaciProvider.FibonacciSequence {
		out := fiboproto.Fib{Index: int32(f.Index), Value: int32(f.Value)}
		result = append(result, &out)
	}
	resp := &fiboproto.SeqResponse{
		Seq: result,
	}
	return resp, nil

}
