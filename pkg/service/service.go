package service

import (
	"fmt"
	"github.com/ecnuvj/vhoj_db/pkg/dao/mapper/user_mapper"
	"github.com/ecnuvj/vhoj_user/pkg/adapter"
	"github.com/ecnuvj/vhoj_user/pkg/sdk/userpb"
)

type UserService struct {
}

func (UserService) Login(username string, password string) error {
	user, err := user_mapper.UserMapper.FindUserByUsername(username)
	if err != nil {
		return err
	}
	if user == nil {
		return fmt.Errorf("user not exist")
	}
	if password != user.UserAuth.Password {
		return fmt.Errorf("password error")
	}
	return nil
}

func (UserService) Register(rpcUser *userpb.User) (*userpb.User, error) {
	user, err := user_mapper.UserMapper.AddUser(adapter.RpcUserToModelUser(rpcUser))
	if err != nil {
		return nil, err
	}
	//todo normal user
	return adapter.ModelUserToRpcUser(user), nil
}
