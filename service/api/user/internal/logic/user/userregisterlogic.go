package logic

import (
	"context"
	"gitee.com/phillip_huang/redis-demo/service/api/user/internal/svc"
	"gitee.com/phillip_huang/redis-demo/service/api/user/internal/types"
	userRpc "gitee.com/phillip_huang/redis-demo/service/rpc/user/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type UserRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) UserRegisterLogic {
	return UserRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRegisterLogic) UserRegister(req types.UserRegisterRequest) (*types.UserRegisterResponse, error) {
	if (req.Username == "" && req.Phone == "" && req.Email == "") || req.Password == "" || len(req.Password) > 16 {
		return &types.UserRegisterResponse{ErrorCode: 1, ErrorMessage: "Invalid request value!"}, nil
	}
	rpcReq := userRpc.UserRegisterRequest{
		Username:             req.Username,
		Password:             req.Password,
		Phone:                req.Phone,
		Email:                req.Email,
	}
	rpcResp, err := l.svcCtx.UserRpc.UserRegister(l.ctx, &rpcReq)
	if err != nil || rpcResp == nil{
		l.Error("UserRpc.UserRegister(l.ctx, &rpcReq) error", err)
		return &types.UserRegisterResponse{ErrorCode: 500, ErrorMessage: "server error!"}, nil
	}
	return &types.UserRegisterResponse{ErrorCode: rpcResp.ErrorCode, ErrorMessage: rpcResp.ErrorMessage}, nil
}
