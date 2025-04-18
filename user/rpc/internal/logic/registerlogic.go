package logic

import (
	"context"
	"rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
	"rpc/internal/svc"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.UserRegisterReq) (*user.UserRegisterResp, error) {
	// todo: add your logic here and delete this line
	
	return &user.UserRegisterResp{}, nil
}
