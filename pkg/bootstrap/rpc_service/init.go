package rpc_service

import (
	"fmt"
	"github.com/ecnuvj/vhoj_common/pkg/common/constants/rpc_config"
	user "github.com/ecnuvj/vhoj_user/pkg/handler"
	"github.com/ecnuvj/vhoj_user/pkg/sdk/userpb"
	"google.golang.org/grpc"
	"log"
	"net"
)

func InitService() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", rpc_config.UserRpc.Address.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	handler, err := user.NewUserHandler()
	if err != nil {
		log.Fatalf("failed to create handler: %v", err)
	}

	s := grpc.NewServer()
	userpb.RegisterUserServiceServer(s, handler)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
