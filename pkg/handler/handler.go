package handler

import (
	"context"
	"fmt"
	"github.com/ecnuvj/vhoj_user/pkg/sdk/base"
	"github.com/ecnuvj/vhoj_user/pkg/sdk/userpb"
	"github.com/ecnuvj/vhoj_user/pkg/service"
	"github.com/ecnuvj/vhoj_user/pkg/util"
)

type UserHandler struct {
	userpb.UnimplementedUserServiceServer
	UserService *service.UserService
}

func NewUserHandler() (*UserHandler, error) {
	return &UserHandler{
		UserService: &service.UserService{},
	}, nil
}

func (u *UserHandler) Login(ctx context.Context, request *userpb.LoginRequest) (*userpb.LoginResponse, error) {
	if request == nil {
		return &userpb.LoginResponse{
			BaseResponse: util.PbReplyf(base.REPLY_STATUS_FAILURE, "login request is nil"),
		}, fmt.Errorf("login request is nil")
	}
	user, err := u.UserService.Login(request.Nickname, request.Password)
	if err != nil {
		return &userpb.LoginResponse{
			BaseResponse: util.PbReplyf(base.REPLY_STATUS_FAILURE, "service error: %v", err),
		}, fmt.Errorf("service error: %v", err)
	}
	return &userpb.LoginResponse{
		User:         user,
		BaseResponse: util.PbReplyf(base.REPLY_STATUS_SUCCESS, "success"),
	}, nil
}

func (u *UserHandler) Register(ctx context.Context, request *userpb.RegisterRequest) (*userpb.RegisterResponse, error) {
	if request == nil {
		return &userpb.RegisterResponse{
			BaseResponse: util.PbReplyf(base.REPLY_STATUS_FAILURE, "register request is nil"),
		}, fmt.Errorf("register request is nil")
	}
	user, err := u.UserService.Register(request.User)
	if err != nil {
		return &userpb.RegisterResponse{
			BaseResponse: util.PbReplyf(base.REPLY_STATUS_FAILURE, "service error: %v", err),
		}, fmt.Errorf("service error: %v", err)
	}
	return &userpb.RegisterResponse{
		User:         user,
		BaseResponse: util.PbReplyf(base.REPLY_STATUS_SUCCESS, "success"),
	}, nil
}

func (u *UserHandler) GenerateUsers(ctx context.Context, request *userpb.GenerateUsersRequest) (*userpb.GenerateUsersResponse, error) {
	if request == nil {
		return &userpb.GenerateUsersResponse{
			BaseResponse: util.PbReplyf(base.REPLY_STATUS_FAILURE, "request is nil"),
		}, fmt.Errorf("request is nil")
	}
	users, err := u.UserService.GenerateUsers(uint(request.ContestId), request.GenerateCount)
	if err != nil {
		return &userpb.GenerateUsersResponse{
			BaseResponse: util.PbReplyf(base.REPLY_STATUS_FAILURE, "service error: %v", err),
		}, fmt.Errorf("service error: %v", err)
	}
	return &userpb.GenerateUsersResponse{
		Users:        users,
		BaseResponse: util.PbReplyf(base.REPLY_STATUS_SUCCESS, "success"),
	}, nil
}

func (u *UserHandler) GetUserRoles(ctx context.Context, request *userpb.GetUserRolesRequest) (*userpb.GetUserRolesResponse, error) {
	if request == nil {
		return &userpb.GetUserRolesResponse{
			BaseResponse: util.PbReplyf(base.REPLY_STATUS_FAILURE, "request is nil"),
		}, fmt.Errorf("request is nil")
	}
	roles, err := u.UserService.GetUserRoles(uint(request.UserId))
	if err != nil {
		return &userpb.GetUserRolesResponse{
			BaseResponse: util.PbReplyf(base.REPLY_STATUS_FAILURE, "service error: %v", err),
		}, fmt.Errorf("service error: %v", err)
	}
	return &userpb.GetUserRolesResponse{
		Roles:        roles,
		BaseResponse: util.PbReplyf(base.REPLY_STATUS_SUCCESS, "success"),
	}, nil
}

//role id 必须给到
func (u *UserHandler) UpdateUserRoles(ctx context.Context, request *userpb.UpdateUserRolesRequest) (*userpb.UpdateUserRolesResponse, error) {
	if request == nil {
		return &userpb.UpdateUserRolesResponse{
			BaseResponse: util.PbReplyf(base.REPLY_STATUS_FAILURE, "request is nil"),
		}, fmt.Errorf("request is nil")
	}
	user, err := u.UserService.UpdateUserRoles(uint(request.UserId), request.Roles)
	if err != nil {
		return &userpb.UpdateUserRolesResponse{
			BaseResponse: util.PbReplyf(base.REPLY_STATUS_FAILURE, "service error: %v", err),
		}, fmt.Errorf("service error: %v", err)
	}
	return &userpb.UpdateUserRolesResponse{
		User:         user,
		BaseResponse: util.PbReplyf(base.REPLY_STATUS_SUCCESS, "success"),
	}, nil
}

func (u *UserHandler) GetAllUsers(ctx context.Context, request *userpb.GetAllUsersRequest) (*userpb.GetAllUsersResponse, error) {
	if request == nil {
		return &userpb.GetAllUsersResponse{
			BaseResponse: util.PbReplyf(base.REPLY_STATUS_FAILURE, "request is nil"),
		}, fmt.Errorf("request is nil")
	}
	users, pageInfo, err := u.UserService.GetAllUsers(request.PageNo, request.PageSize)
	if err != nil {
		return &userpb.GetAllUsersResponse{
			BaseResponse: util.PbReplyf(base.REPLY_STATUS_FAILURE, "service error: %v", err),
		}, fmt.Errorf("service error: %v", err)
	}
	return &userpb.GetAllUsersResponse{
		Users:        users,
		TotalPages:   pageInfo.TotalPages,
		TotalCount:   pageInfo.TotalCount,
		BaseResponse: util.PbReplyf(base.REPLY_STATUS_SUCCESS, "success"),
	}, nil
}

func (u *UserHandler) DeleteUsers(ctx context.Context, request *userpb.DeleteUsersRequest) (*userpb.DeleteUsersResponse, error) {
	if request == nil {
		return &userpb.DeleteUsersResponse{
			BaseResponse: util.PbReplyf(base.REPLY_STATUS_FAILURE, "request is nil"),
		}, fmt.Errorf("request is nil")
	}
	userIds := make([]uint, len(request.UserIds))
	for i, u := range request.UserIds {
		userIds[i] = uint(u)
	}
	err := u.UserService.DeleteUsers(userIds)
	if err != nil {
		return &userpb.DeleteUsersResponse{
			BaseResponse: util.PbReplyf(base.REPLY_STATUS_FAILURE, "service error: %v", err),
		}, fmt.Errorf("service error: %v", err)
	}
	return &userpb.DeleteUsersResponse{
		BaseResponse: util.PbReplyf(base.REPLY_STATUS_SUCCESS, "success"),
	}, nil
}

func (u *UserHandler) GetUsersByIds(ctx context.Context, request *userpb.GetUsersByIdsRequest) (*userpb.GetUsersByIdsResponse, error) {
	if request == nil {
		return &userpb.GetUsersByIdsResponse{
			BaseResponse: util.PbReplyf(base.REPLY_STATUS_FAILURE, "request is nil"),
		}, fmt.Errorf("request is nil")
	}
	userIds := make([]uint, len(request.UserIds))
	for i, u := range request.UserIds {
		userIds[i] = uint(u)
	}
	users, err := u.UserService.GetUsersByIds(userIds)
	if err != nil {
		return &userpb.GetUsersByIdsResponse{
			BaseResponse: util.PbReplyf(base.REPLY_STATUS_FAILURE, "service error: %v", err),
		}, fmt.Errorf("service error: %v", err)
	}
	return &userpb.GetUsersByIdsResponse{
		Users:        users,
		BaseResponse: util.PbReplyf(base.REPLY_STATUS_SUCCESS, "success"),
	}, nil
}

func (u *UserHandler) UpdateUserInfo(ctx context.Context, request *userpb.UpdateUserInfoRequest) (*userpb.UpdateUserInfoResponse, error) {
	if request == nil {
		return &userpb.UpdateUserInfoResponse{
			BaseResponse: util.PbReplyf(base.REPLY_STATUS_FAILURE, "request is nil"),
		}, fmt.Errorf("request is nil")
	}
	user, err := u.UserService.UpdateUserInfo(uint(request.UserId), request.User)
	if err != nil {
		return &userpb.UpdateUserInfoResponse{
			BaseResponse: util.PbReplyf(base.REPLY_STATUS_FAILURE, "service error: %v", err),
		}, fmt.Errorf("service error: %v", err)
	}
	return &userpb.UpdateUserInfoResponse{
		User:         user,
		BaseResponse: util.PbReplyf(base.REPLY_STATUS_SUCCESS, "success"),
	}, nil
}
