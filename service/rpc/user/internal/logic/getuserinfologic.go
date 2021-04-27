package logic

import (
	"context"
	userModel "gitee.com/phillip_huang/redis-demo/service/rpc/user/sql/model"
	"strconv"

	"gitee.com/phillip_huang/redis-demo/service/rpc/user/internal/svc"
	"gitee.com/phillip_huang/redis-demo/service/rpc/user/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *user.GetUserInfoRequest) (*user.GetUserInfoResponse, error) {
	var findUser *userModel.User
	var err error
	switch {
	case in.Id != "":
		uid, err := strconv.ParseInt(in.Id, 10, 64)
		if err != nil{
			return &user.GetUserInfoResponse{ErrorCode: 400, ErrorMessage: "Invalid user id!"}, nil
		}
		findUser, err = l.svcCtx.UserModel.FindOne(uid)
	case in.Username != "":
		findUser, err = l.svcCtx.UserModel.FindOneByUsername(in.Username)
	case in.Phone != "":
		findUser, err = l.svcCtx.UserModel.FindOneByPhone(in.Phone)
	case in.Email != "":
		findUser, err = l.svcCtx.UserModel.FindOneByEmail(in.Email)
	default:
		return &user.GetUserInfoResponse{ErrorCode: 400, ErrorMessage: "Invalid request!"}, nil

	}
	if err != nil || findUser == nil {
		if err == userModel.ErrNotFound{
			return &user.GetUserInfoResponse{ErrorCode: 400, ErrorMessage: "User not found!"}, nil
		}
		l.Error("UserModel FindUser Error", err)
		return &user.GetUserInfoResponse{ErrorCode: userModel.MysqlNotFoundErrCode, ErrorMessage: "Server find user error!"}, nil
	}
	resp := user.GetUserInfoResponse{
		ErrorCode:            0,
		ErrorMessage:         "",
		Id:                   strconv.FormatInt(findUser.Id, 10),
		Username:             findUser.Username,
		Gender:               findUser.Gender,
		Phone:                findUser.Phone,
		Email:                findUser.Email,
		Nickname:             findUser.Nickname,
		Birthday:             findUser.Birthday.Format(userModel.UserBirthdayFormat),
		Password:             findUser.Password,
	}
	return &resp, nil
}
