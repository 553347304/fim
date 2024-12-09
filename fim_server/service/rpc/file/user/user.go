// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: file.proto

package user

import (
	"context"

	"fim_server/service/rpc/file/file_rpc"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	FileInfoRequest  = file_rpc.FileInfoRequest
	FileInfoResponse = file_rpc.FileInfoResponse

	User interface {
		FileInfo(ctx context.Context, in *FileInfoRequest, opts ...grpc.CallOption) (*FileInfoResponse, error)
	}

	defaultUser struct {
		cli zrpc.Client
	}
)

func NewUser(cli zrpc.Client) User {
	return &defaultUser{
		cli: cli,
	}
}

func (m *defaultUser) FileInfo(ctx context.Context, in *FileInfoRequest, opts ...grpc.CallOption) (*FileInfoResponse, error) {
	client := file_rpc.NewUserClient(m.cli.Conn())
	return client.FileInfo(ctx, in, opts...)
}
