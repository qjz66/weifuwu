package logic

import (
	"context"
	"tiny_service/apps/user/user/user"

	"tiny_service/apps/api/internal/svc"
	"tiny_service/apps/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.UserLoginrequest) (resp *types.UserLoginresponse, err error) {
	// todo: add your logic here and delete this line
	rpcResp, err := l.svcCtx.UserRPC.Login(l.ctx, &user.UserLoginRequest{Username: req.Username, Password: req.Password})
	if err != nil {
		return nil, err
	}
	resp = &types.UserLoginresponse{Id: rpcResp.Id, Token: rpcResp.Token}
	return resp, nil
}
