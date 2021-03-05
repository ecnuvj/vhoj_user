package service

import (
	"fmt"
	"github.com/ecnuvj/vhoj_common/pkg/common/constants/role"
	"github.com/ecnuvj/vhoj_db/pkg/dao/mapper/user_mapper"
	"github.com/ecnuvj/vhoj_db/pkg/dao/model"
	"github.com/ecnuvj/vhoj_user/pkg/adapter"
	"github.com/ecnuvj/vhoj_user/pkg/common"
	"github.com/ecnuvj/vhoj_user/pkg/sdk/userpb"
	"math/rand"
	"time"
)

type UserService struct {
}

func (UserService) Login(username string, password string) (*userpb.User, error) {
	user, err := user_mapper.UserMapper.FindUserByUsername(username)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, fmt.Errorf("user not exist")
	}
	if password != user.UserAuth.Password {
		return nil, fmt.Errorf("password error")
	}
	return adapter.ModelUserToRpcUser(user), nil
}

func (UserService) Register(rpcUser *userpb.User) (*userpb.User, error) {
	_, err := user_mapper.UserMapper.FindUserByUsername(rpcUser.Username)
	if err == nil {
		return nil, fmt.Errorf("user already exist")
	}
	//注册用户都是normal用户
	rpcUser.Roles = []*userpb.Role{
		{
			RoleName: role.NORMAL_USER,
		},
	}
	user, err := user_mapper.UserMapper.AddUser(adapter.RpcUserToModelUser(rpcUser))
	if err != nil {
		return nil, err
	}
	return adapter.ModelUserToRpcUser(user), nil
}

func GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzQWERTYUIOPASDFGHJKLZXCVBNM.+-"
	bytes := []byte(str)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func (UserService) GenerateUsers(contestId uint, count int32) ([]*userpb.User, error) {
	users := make([]*userpb.User, count)
	for i := 0; i < int(count); {
		generateUser := &model.User{
			UserAuth: &model.UserAuth{
				Password: GetRandomString(8),
			},
			Nickname: GetRandomString(10),
			Roles: []*model.Role{
				{
					RoleName: role.GENERATOR_USER,
				},
			},
			GenerateUser: true,
		}
		user, err := user_mapper.UserMapper.AddUser(generateUser)
		if err == nil {
			users[i] = adapter.ModelUserToRpcUser(user)
			i++
		}
	}
	return users, nil
}

func (UserService) GetUserRoles(userId uint) ([]*userpb.Role, error) {
	roles, err := user_mapper.UserMapper.FindUserRolesById(userId)
	if err != nil {
		return nil, err
	}
	return adapter.ModelRolesToRpcRoles(roles), nil
}

//role id 必须给到
func (UserService) UpdateUserRoles(userId uint, roles []*userpb.Role) (*userpb.User, error) {
	user, err := user_mapper.UserMapper.UpdateUserRoles(userId, adapter.RpcRolesToModelRoles(roles))
	if err != nil {
		return nil, err
	}
	return adapter.ModelUserToRpcUser(user), nil
}

func (UserService) GetAllUsers(pageNo int32, pageSize int32) ([]*userpb.User, *common.PageInfo, error) {
	users, count, err := user_mapper.UserMapper.FindAllUsers(pageNo, pageSize)
	if err != nil {
		return nil, &common.PageInfo{}, nil
	}
	rpcUsers := make([]*userpb.User, len(users))
	for i, u := range users {
		rpcUsers[i] = adapter.ModelUserToRpcUser(u)
	}
	return rpcUsers, &common.PageInfo{
		TotalPages: (count + pageSize - 1) / pageSize,
		TotalCount: count,
	}, nil
}

func (UserService) DeleteUsers(userIds []uint) error {
	for _, id := range userIds {
		err := user_mapper.UserMapper.DeleteUserById(id)
		if err != nil {
			return err
		}
	}
	return nil
}

func (UserService) GetUsersByIds(userIds []uint) ([]*userpb.User, error) {
	users, err := user_mapper.UserMapper.FindUsersByIds(userIds)
	if err != nil {
		return nil, err
	}
	return adapter.ModelUsersToRpcUsers(users), nil
}

func (UserService) UpdateUserInfo(userId uint, user *userpb.User) (*userpb.User, error) {
	user.UserId = uint64(userId)
	retUser, err := user_mapper.UserMapper.UpdateUser(adapter.RpcUserToModelUser(user))
	if err != nil {
		return nil, err
	}
	return adapter.ModelUserToRpcUser(retUser), nil
}

func (UserService) GetUserById(userId uint) (*userpb.User, error) {
	user, err := user_mapper.UserMapper.FindUserById(userId)
	if err != nil {
		return nil, err
	}
	return adapter.ModelUserToRpcUser(user), nil
}

func (UserService) GetUserByUsername(username string) (*userpb.User, error) {
	user, err := user_mapper.UserMapper.FindUserByUsername(username)
	if err != nil {
		return nil, err
	}
	return adapter.ModelUserToRpcUser(user), nil
}
