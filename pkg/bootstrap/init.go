package bootstrap

import (
	"github.com/ecnuvj/vhoj_user/pkg/bootstrap/database"
	"github.com/ecnuvj/vhoj_user/pkg/bootstrap/rpc_service"
)

func Init() {
	database.Init()
	rpc_service.InitService()
}
