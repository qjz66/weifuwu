package logic

import (
	"context"
	"rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
	"rpc/internal/svc"
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

func (l *LoginLogic) Login(in *user.UserLoginReq) (*user.UserLoginReqResp, error) {
	// todo: add your logic here and delete this line

	return &user.UserLoginReqResp{}, nil
}
