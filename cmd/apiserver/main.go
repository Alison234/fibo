package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/go/http-rest-api/internal/app/apiserver"

	grpcfibo "github.com/go/http-rest-api/grpc/proto/grpcfibo"
)

var (
	configPath string
	ServerType string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
	flag.StringVar(&ServerType, "serverType", "http", "server type")
}

func main() {
	cacheAddr := "localhost:11211"
	flag.Parse()

	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	if ServerType == "grpc" {
		gprcServ := grpcfibo.NewGRPCServer(cacheAddr)
		gprcServ.Start(
			&grpcfibo.GRPCServerConfig{
				Port:        "11564",
				Host:        "127.0.0.1",
				ConnTimeout: time.Minute * 5,
			},
		)
	} else {
		s := apiserver.New(config)
		fmt.Println(s)
		if err := s.Start(); err != nil {
			log.Fatal(err)
		}
	}

}
