// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: chat.proto

package server

import (
	"context"

	"fim_server/service/rpc/chat/chat_rpc"
	"fim_server/service/rpc/chat/internal/logic"
	"fim_server/service/rpc/chat/internal/svc"
)

type ChatServer struct {
	svcCtx *svc.ServiceContext
	chat_rpc.UnimplementedChatServer
}

func NewChatServer(svcCtx *svc.ServiceContext) *ChatServer {
	return &ChatServer{
		svcCtx: svcCtx,
	}
}

func (s *ChatServer) UserChat(ctx context.Context, in *chat_rpc.UserChatRequest) (*chat_rpc.UserChatResponse, error) {
	l := logic.NewUserChatLogic(ctx, s.svcCtx)
	return l.UserChat(in)
}
