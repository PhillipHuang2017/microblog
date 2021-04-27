// Code generated by goctl. DO NOT EDIT!
// Source: user.proto

//go:generate mockgen -destination ./user_mock.go -package userclient -source $GOFILE

package userclient

import (
	"context"

	"gitee.com/phillip_huang/redis-demo/service/rpc/user/user"

	"github.com/tal-tech/go-zero/zrpc"
)

type (
	UserRegisterRequest  = user.UserRegisterRequest
	UserRegisterResponse = user.UserRegisterResponse
	GetUserInfoRequest   = user.GetUserInfoRequest
	GetUserInfoResponse  = user.GetUserInfoResponse

	User interface {
		UserRegister(ctx context.Context, in *UserRegisterRequest) (*UserRegisterResponse, error)
		GetUserInfo(ctx context.Context, in *GetUserInfoRequest) (*GetUserInfoResponse, error)
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

func (m *defaultUser) UserRegister(ctx context.Context, in *UserRegisterRequest) (*UserRegisterResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.UserRegister(ctx, in)
}

func (m *defaultUser) GetUserInfo(ctx context.Context, in *GetUserInfoRequest) (*GetUserInfoResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.GetUserInfo(ctx, in)
}