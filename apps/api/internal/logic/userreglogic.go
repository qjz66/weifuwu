package logic

import (
	"context"
	"errors"
	"tiny_service/apps/user/user/user"

	"tiny_service/apps/api/internal/svc"
	"tiny_service/apps/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRegLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRegLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegLogic {
	return &UserRegLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRegLogic) UserReg(req *types.UserRegrequest) (resp *types.UserRegresponse, err error) {
	// todo: add your logic here and delete this line
	if l.svcCtx.UserRPC == nil {
		return nil, errors.New("RPC 客户端未初始化")
	}
	if req.Username == "" || req.Password == "" {
		return nil, errors.New("用户名或密码不能为空")
	}

	rpcResp, err := l.svcCtx.UserRPC.Reg(l.ctx, &user.UserRegRequest{Username: req.Username, Password: req.Password})
	if err != nil || rpcResp == nil {
		return nil, err
	}

	resp = &types.UserRegresponse{Id: rpcResp.Id}

	return resp, nil
}
