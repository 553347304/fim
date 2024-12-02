// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: user.proto

package server

import (
	"context"

	"fim_server/service/rpc/user/internal/logic"
	"fim_server/service/rpc/user/internal/svc"
	"fim_server/service/rpc/user/user_rpc"
)

type UserServer struct {
	svcCtx *svc.ServiceContext
	user_rpc.UnimplementedUserServer
}

func NewUserServer(svcCtx *svc.ServiceContext) *UserServer {
	return &UserServer{
		svcCtx: svcCtx,
	}
}

func (s *UserServer) UserCreate(ctx context.Context, in *user_rpc.UserCreateRequest) (*user_rpc.UserCreateResponse, error) {
	l := logic.NewUserCreateLogic(ctx, s.svcCtx)
	return l.UserCreate(in)
}

func (s *UserServer) UserInfo(ctx context.Context, in *user_rpc.UserInfoRequest) (*user_rpc.UserInfoResponse, error) {
	l := logic.NewUserInfoLogic(ctx, s.svcCtx)
	return l.UserInfo(in)
}

func (s *UserServer) UserListInfo(ctx context.Context, in *user_rpc.UserListInfoRequest) (*user_rpc.UserListInfoResponse, error) {
	l := logic.NewUserListInfoLogic(ctx, s.svcCtx)
	return l.UserListInfo(in)
}
