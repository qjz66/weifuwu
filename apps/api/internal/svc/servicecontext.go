package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"tiny_service/apps/api/internal/config"
	"tiny_service/apps/product/product/product"
	"tiny_service/apps/user/user/user"
)

type ServiceContext struct {
	Config     config.Config
	UserRPC    user.UserClient
	ProductRPC product.ProductClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		UserRPC:    user.NewUserClient(zrpc.MustNewClient(c.UserRPC).Conn()),
		ProductRPC: product.NewProductClient(zrpc.MustNewClient(c.ProductRPC).Conn()),
	}
}
