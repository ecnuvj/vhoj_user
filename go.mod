module github.com/ecnuvj/vhoj_user

go 1.14

replace github.com/ecnuvj/vhoj_db => ../vhoj_db

replace github.com/ecnuvj/vhoj_common => ../vhoj_common

require (
	github.com/ecnuvj/vhoj_db v0.0.0-00010101000000-000000000000
	github.com/golang/protobuf v1.4.3
	google.golang.org/grpc v1.35.0
	google.golang.org/protobuf v1.25.0
)
