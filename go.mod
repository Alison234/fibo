module github.com/go/http-rest-api

go 1.16

require (
	github.com/BurntSushi/toml v0.4.1
	github.com/bradfitz/gomemcache v0.0.0-20190913173617-a41fca850d0b
	github.com/gorilla/mux v1.8.0
	github.com/sirupsen/logrus v1.8.1
	github.com/stretchr/testify v1.7.0
	golang.org/x/sys v0.0.0-20211210111614-af8b64212486 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	google.golang.org/grpc v1.42.0
	google.golang.org/protobuf v1.27.1
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)

replace google.golang.org/grpc => github.com/grpc/grpc-go v1.42.0
