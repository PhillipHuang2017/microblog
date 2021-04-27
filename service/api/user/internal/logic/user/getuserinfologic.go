package logic

import (
	"context"
	userRpc "gitee.com/phillip_huang/redis-demo/service/rpc/user/user"

	"gitee.com/phillip_huang/redis-demo/service/api/user/internal/svc"
	"gitee.com/phillip_huang/redis-demo/service/api/user/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetUserInfoLogic {
	return GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo() (*types.GetUserInfoResponse, error) {
	uid := l.ctx.Value("uid").(string)
	rpcResp, err := l.svcCtx.UserRpc.GetUserInfo(l.ctx, &userRpc.GetUserInfoRequest{Id: uid})
	if err != nil || rpcResp == nil {
		l.Error("getUserInfoByRpc failed", err)
		return &types.GetUserInfoResponse{ErrorCode: 500, ErrorMessage: "server error"}, nil
	}
	resp := types.GetUserInfoResponse{
		ErrorCode:    rpcResp.ErrorCode,
		ErrorMessage: rpcResp.ErrorMessage,
		Id:           rpcResp.Id,
		Username:     rpcResp.Username,
		Gender:       rpcResp.Gender,
		Phone:        rpcResp.Phone,
		Email:        rpcResp.Email,
		Nickname:     rpcResp.Nickname,
		Birthday:     rpcResp.Birthday,
	}

	return &resp, nil
}
