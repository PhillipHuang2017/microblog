package logic

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"gitee.com/phillip_huang/redis-demo/service/rpc/user/internal/svc"
	userModel "gitee.com/phillip_huang/redis-demo/service/rpc/user/sql/model"
	userRpc "gitee.com/phillip_huang/redis-demo/service/rpc/user/user"
	"github.com/go-sql-driver/mysql"
	"github.com/tal-tech/go-zero/core/logx"
)

type UserRegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserRegisterLogic) UserRegister(in *userRpc.UserRegisterRequest) (*userRpc.UserRegisterResponse, error) {
	if (in.Username == "" && in.Phone == "" && in.Email == "") || in.Password == "" || len(in.Password) > 16 {
		return &userRpc.UserRegisterResponse{ErrorCode: 1, ErrorMessage: "Invalid request value!"}, nil
	}

	hash := hmac.New(sha256.New, []byte(l.svcCtx.Config.UserPasswordSalt))
	_, _ = hash.Write([]byte(in.Password))
	newUser := userModel.User{
		Password:   base64.StdEncoding.EncodeToString(hash.Sum(nil)),
		Phone:      in.Phone,
		Email:      in.Email,
		Username:   in.Username,
		Gender:     userModel.UserDefaultGender,
		Birthday:   userModel.UserDefaultBirthday,
	}
	_, err := l.svcCtx.UserModel.Insert(newUser)
	if err != nil {
		sqlErr, ok := err.(*mysql.MySQLError)
		if ok && sqlErr.Number == userModel.MysqlDuplicateErrCode{
			return &userRpc.UserRegisterResponse{ErrorCode: 400, ErrorMessage: "User existed!"}, nil
		}
		l.Infof("UserModel.Insert(newUser) error!", err, newUser)
		return &userRpc.UserRegisterResponse{ErrorCode: 500, ErrorMessage: "Server error!"}, err
	}
	return &userRpc.UserRegisterResponse{ErrorCode: 0, ErrorMessage: ""}, nil
}


