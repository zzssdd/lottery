module micro-prize

go 1.14

require (
	github.com/golang/protobuf v1.5.2
	google.golang.org/protobuf v1.28.1
)

require (
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/google/go-cmp v0.5.6 // indirect
	github.com/google/uuid v1.1.2 // indirect
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/kr/pretty v0.2.0 // indirect
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/registry/consul/v2 v2.9.1
	github.com/onsi/gomega v1.15.0 // indirect
	github.com/stretchr/testify v1.7.0 // indirect
	go.etcd.io/bbolt v1.3.5 // indirect
	go.uber.org/atomic v1.6.0 // indirect
	golang.org/x/crypto v0.0.0-20201016220609-9e8e0b390897 // indirect
	google.golang.org/genproto v0.0.0-20200806141610-86f49bd18e98 // indirect
	google.golang.org/grpc v1.40.0 // indirect
)

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
