package logic

import (
	"context"
	"tiny_service/apps/user/user/internal/common"
	"tiny_service/apps/user/user/internal/svc"
	"tiny_service/apps/user/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.UserLoginRequest) (*user.UserLoginResponse, error) {
	// todo: add your logic here and delete this line
	token, err := common.GetJwtToken(l.svcCtx.Config.AuthInfo.AccessSecret, l.svcCtx.Config.AuthInfo.AccessExpire, 1)
	if err != nil {
		return nil, err
	}
	return &user.UserLoginResponse{Id: 1, Token: token}, nil
}
