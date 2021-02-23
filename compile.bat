cd pkg/sdk
protoc -I . --go_out=.  --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative userpb/service.proto
protoc -I . --go_out=.  --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative userpb/user.proto
protoc -I . --go_out=.  --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative userpb/role.proto
protoc -I . --go_out=.  --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative base/base.proto
cd ../..

xcopy pkg\sdk\userpb\*.go ..\vhoj_rpc\model\userpb\ /s /y