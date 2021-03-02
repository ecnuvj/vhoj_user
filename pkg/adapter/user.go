package adapter

import (
	"github.com/ecnuvj/vhoj_db/pkg/dao/model"
	"github.com/ecnuvj/vhoj_user/pkg/sdk/userpb"
	"github.com/jinzhu/gorm"
)

func RpcUserToModelUser(user *userpb.User) *model.User {
	roles := RpcRolesToModelRoles(user.Roles)
	userAuth := &model.UserAuth{
		Model:    gorm.Model{ID: uint(user.UserAuthId)},
		Password: user.Password,
	}
	if user.UserAuthId == 0 || user.Password == "" {
		userAuth = nil
	}
	return &model.User{
		Model: gorm.Model{
			ID: uint(user.UserId),
		},
		UserAuth:  userAuth,
		Email:     user.Email,
		Nickname:  user.Username,
		School:    user.School,
		Roles:     roles,
		Submitted: user.Submitted,
		Accepted:  user.Accepted,
	}
}

func ModelUserToRpcUser(user *model.User) *userpb.User {
	roles := ModelRolesToRpcRoles(user.Roles)
	var userAuthId uint64
	var password string
	if user.UserAuth != nil {
		userAuthId = uint64(user.UserAuth.ID)
		password = user.UserAuth.Password
	}
	return &userpb.User{
		UserId:     uint64(user.ID),
		Username:   user.Nickname,
		UserAuthId: userAuthId,
		Password:   password,
		Email:      user.Email,
		School:     user.School,
		Roles:      roles,
		Accepted:   user.Accepted,
		Submitted:  user.Submitted,
	}
}

func RpcUsersToModelUsers(users []*userpb.User) []*model.User {
	retUsers := make([]*model.User, len(users))
	for i, u := range users {
		retUsers[i] = RpcUserToModelUser(u)
	}
	return retUsers
}

func ModelUsersToRpcUsers(users []*model.User) []*userpb.User {
	retUsers := make([]*userpb.User, len(users))
	for i, u := range users {
		retUsers[i] = ModelUserToRpcUser(u)
	}
	return retUsers
}
