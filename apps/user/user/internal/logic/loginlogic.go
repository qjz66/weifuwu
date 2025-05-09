package logic

import (
	"context"
	"tiny_service/apps/user/user/dao"
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
	u, err := dao.GetUserByPwd(in.Username, in.Password)
	if err != nil {
		return nil, err
	}
	token, err := common.GetJwtToken(l.svcCtx.Config.AuthInfo.AccessSecret, l.svcCtx.Config.AuthInfo.AccessExpire, int64(u.ID))
	if err != nil {
		return nil, err
	}
	return &user.UserLoginResponse{Id: int64(u.ID), Token: token}, nil
}
