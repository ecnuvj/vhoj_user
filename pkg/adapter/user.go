package adapter

import (
	"github.com/ecnuvj/vhoj_db/pkg/dao/model"
	"github.com/ecnuvj/vhoj_user/pkg/sdk/userpb"
)

func RpcUserToModelUser(user *userpb.User) *model.User {
	roles := make([]*model.Role, len(user.Roles))
	for i, r := range user.Roles {
		roles[i] = &model.Role{
			RoleName: r,
		}
	}
	return &model.User{
		UserAuth: &model.UserAuth{
			UserID:   uint(user.UserAuthId),
			Password: user.Password,
		},
		Email:     user.Email,
		Nickname:  user.Username,
		School:    user.School,
		Roles:     roles,
		Submitted: user.Submitted,
		Accepted:  user.Accepted,
	}
}

func ModelUserToRpcUser(user *model.User) *userpb.User {
	roles := make([]string, len(user.Roles))
	for i, r := range user.Roles {
		roles[i] = r.RoleName
	}
	return &userpb.User{
		UserId:     uint64(user.ID),
		Username:   user.Nickname,
		UserAuthId: uint64(user.UserAuth.ID),
		Password:   user.UserAuth.Password,
		Email:      user.Email,
		School:     user.School,
		Roles:      roles,
		Accepted:   user.Accepted,
		Submitted:  user.Submitted,
	}
}
