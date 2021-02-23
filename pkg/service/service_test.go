package service

import (
	"encoding/json"
	"fmt"
	"github.com/ecnuvj/vhoj_common/pkg/common/constants/role"
	"github.com/ecnuvj/vhoj_user/pkg/bootstrap/database"
	"github.com/ecnuvj/vhoj_user/pkg/sdk/userpb"
	"testing"
)

func TestUserService_Register(t *testing.T) {
	database.Init()
	rpcUser := &userpb.User{
		Username: "test account",
		Password: "123456666",
		Email:    "1231241@qq.com",
		School:   "ecnu",
		Roles: []*userpb.Role{
			{
				RoleName: role.CODE_REVIEWER,
			},
			{
				RoleName: role.CONTEST_CREATOR,
			},
		},
	}
	service := &UserService{}
	user, err := service.Register(rpcUser)
	if err != nil {
		fmt.Printf("err: %v", err)
		return
	}
	str, _ := json.Marshal(user)
	fmt.Println(string(str))
}

func TestUserService_Login(t *testing.T) {
	database.Init()
	service := &UserService{}
	user, err := service.Login("tt", "123456")
	if err != nil {
		fmt.Printf("err: %v", err)
		return
	}
	str, _ := json.Marshal(user)
	fmt.Println(string(str))
}

func TestUserService_DeleteUsers(t *testing.T) {
	database.Init()
	service := &UserService{}
	err := service.DeleteUsers([]uint{1})
	if err != nil {
		fmt.Printf("err: %v",err)
		return 
	}
}
