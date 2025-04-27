package logic

import (
	"context"
	"tiny_service/apps/user/user/dao"
	"tiny_service/apps/user/user/internal/models"

	"tiny_service/apps/user/user/internal/svc"
	"tiny_service/apps/user/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegLogic {
	return &RegLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegLogic) Reg(in *user.UserRegRequest) (*user.UserRegResponse, error) {
	// todo: add your logic here and delete this line
	u := &models.User{Username: in.Username, Password: in.Password}
	err := dao.AddUser(u)
	if err != nil {
		return nil, err
	}
	return &user.UserRegResponse{Id: int64(u.ID)}, nil
}
