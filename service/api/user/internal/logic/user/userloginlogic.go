package logic

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	userModel "gitee.com/phillip_huang/redis-demo/service/rpc/user/sql/model"
	userRpc "gitee.com/phillip_huang/redis-demo/service/rpc/user/user"
	"time"

	"gitee.com/phillip_huang/redis-demo/service/api/user/internal/svc"
	"gitee.com/phillip_huang/redis-demo/service/api/user/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) UserLoginLogic {
	return UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req types.UserLoginRequest) (*types.UserLoginResponse, error) {
	// 获取用户信息，检查密码
	rpcUserInfo, err := l.svcCtx.UserRpc.GetUserInfo(l.ctx, &userRpc.GetUserInfoRequest{Username: req.Username, Phone: req.Phone, Email: req.Email})
	if err != nil || rpcUserInfo == nil {
		l.Error("getUserInfoLogic.getUserInfoByRpc error", err)
		return &types.UserLoginResponse{ErrorCode: 500, ErrorMessage: "Server get user info error!"}, nil
	}
	if rpcUserInfo.ErrorCode == userModel.MysqlNotFoundErrCode {
		return &types.UserLoginResponse{ErrorCode: 400, ErrorMessage: "User not existed!"}, nil
	}
	if !checkPassword(req.Password, rpcUserInfo.Password, l.svcCtx.Config.UserPasswordSalt) {
		return &types.UserLoginResponse{ErrorCode: 400, ErrorMessage: "Incorrect user or password"}, nil
	}

	// 生成jwt的token
	token, err := getToken(l.svcCtx.Config.Auth.AccessSecret, time.Now().Unix(), l.svcCtx.Config.Auth.AccessExpire, rpcUserInfo.Id)
	if err != nil {
		l.Error("Generate token error!")
		return &types.UserLoginResponse{ErrorCode: 500, ErrorMessage: "Generate token failed！"}, nil
	}
	return &types.UserLoginResponse{ErrorCode: 0, Token: token}, nil
}

func checkPassword(reqPassword, realPasswordHash, salt string) bool {
	hash := hmac.New(sha256.New, []byte(salt))
	_, _ = hash.Write([]byte(reqPassword))
	expect, _ :=base64.StdEncoding.DecodeString(realPasswordHash)
	return hmac.Equal(hash.Sum(nil), expect)  // hmac.New对比哈希值，防止时间序列攻击
}