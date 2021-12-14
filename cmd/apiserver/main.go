package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/go/http-rest-api/internal/app/apiserver"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}
	s := apiserver.New(config)
	fmt.Println(s)

	// gprcServ := addr.NewGRPCServer()
	// gprcServ.Start(
	// 	&addr.GRPCServerConfig{
	// 		Port:        "11564",
	// 		Host:        "127.0.0.1",
	// 		ConnTimeout: time.Minute * 5,
	// 	},
	// )
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}

}
