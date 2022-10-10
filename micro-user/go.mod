module micro-user

go 1.14

require (
	github.com/golang/protobuf v1.5.2
	github.com/micro/micro/v3 v3.13.0
	google.golang.org/protobuf v1.28.1
)

require (
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/jordan-wright/email v4.0.1-0.20210109023952-943e75fe5223+incompatible // indirect
	github.com/micro/go-micro/v2 v2.9.1 // indirect
	github.com/micro/go-plugins/registry/consul/v2 v2.9.1 // indirect
	gorm.io/driver/mysql v1.3.6 // indirect
	gorm.io/gorm v1.23.10 // indirect
)

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.27.1
