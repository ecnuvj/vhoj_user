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

func (u *UserHandler) Login(ctx context.Context, request *userpb.LoginRequest) (*userpb.LoginResponse, error) {
	if request == nil {
		return &userpb.LoginResponse{
			BaseResponse: util.PbReplyf(base.REPLY_STATUS_FAILURE, "login request is nil"),
		}, fmt.Errorf("login request is nil")
	}
	//
	return &userpb.LoginResponse{
		BaseResponse: util.PbReplyf(base.REPLY_STATUS_SUCCESS, "success"),
	}, nil
}

func (u *UserHandler) Register(ctx context.Context, request *userpb.RegisterRequest) (*userpb.RegisterResponse, error) {
	if request == nil {
		return &userpb.RegisterResponse{
			BaseResponse: util.PbReplyf(base.REPLY_STATUS_FAILURE, "register request is nil"),
		}, fmt.Errorf("register request is nil")
	}
	return &userpb.RegisterResponse{
		BaseResponse: util.PbReplyf(base.REPLY_STATUS_SUCCESS, "success"),
	}, nil
}

func (u *UserHandler) GenerateUsers(ctx context.Context, request *userpb.GenerateUsersRequest) (*userpb.GenerateUsersResponse, error) {
	if request == nil {
		return &userpb.GenerateUsersResponse{
			BaseResponse: util.PbReplyf(base.REPLY_STATUS_FAILURE, "request is nil"),
		}, fmt.Errorf("request is nil")
	}

	return &userpb.GenerateUsersResponse{
		Users:        nil,
		BaseResponse: util.PbReplyf(base.REPLY_STATUS_SUCCESS, "success"),
	}, nil
}

func (u *UserHandler) GetUserRoles(ctx context.Context, request *userpb.GetUserRolesRequest) (*userpb.GetUserRolesResponse, error) {
	if request == nil {
		return &userpb.GetUserRolesResponse{
			BaseResponse: util.PbReplyf(base.REPLY_STATUS_FAILURE, "request is nil"),
		}, fmt.Errorf("request is nil")
	}

	return &userpb.GetUserRolesResponse{
		Roles:        nil,
		BaseResponse: util.PbReplyf(base.REPLY_STATUS_SUCCESS, "success"),
	}, nil
}

func (u *UserHandler) UpdateUserRoles(ctx context.Context, request *userpb.UpdateUserRolesRequest) (*userpb.UpdateUserRolesResponse, error) {
	if request == nil {
		return &userpb.UpdateUserRolesResponse{
			BaseResponse: util.PbReplyf(base.REPLY_STATUS_FAILURE, "request is nil"),
		}, fmt.Errorf("request is nil")
	}
	return &userpb.UpdateUserRolesResponse{
		BaseResponse: util.PbReplyf(base.REPLY_STATUS_SUCCESS, "success"),
	}, nil
}

func (u *UserHandler) GetAllUsers(ctx context.Context, request *userpb.GetAllUsersRequest) (*userpb.GetAllUsersResponse, error) {
	if request == nil {
		return &userpb.GetAllUsersResponse{
			BaseResponse: util.PbReplyf(base.REPLY_STATUS_FAILURE, "request is nil"),
		}, fmt.Errorf("request is nil")
	}

	return &userpb.GetAllUsersResponse{
		User:         nil,
		TotalPages:   0,
		TotalCount:   0,
		BaseResponse: util.PbReplyf(base.REPLY_STATUS_SUCCESS, "success"),
	}, nil
}

func (u *UserHandler) DeleteUsers(ctx context.Context, request *userpb.DeleteUsersRequest) (*userpb.DeleteUsersResponse, error) {
	if request == nil {
		return &userpb.DeleteUsersResponse{
			BaseResponse: util.PbReplyf(base.REPLY_STATUS_FAILURE, "request is nil"),
		}, fmt.Errorf("request is nil")
	}

	return &userpb.DeleteUsersResponse{
		BaseResponse: util.PbReplyf(base.REPLY_STATUS_SUCCESS, "success"),
	}, nil
}
