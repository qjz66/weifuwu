package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	UserRPC  zrpc.RpcClientConf
	Mysql    *MysqlConf
	AuthInfo struct {
		AccessSecret string
		AccessExpire int64
	}
	//UserRPC zrpc.RpcClientConf
}
type MysqlConf struct {
	DataSource string // mysql链接地址，满足 $user:$password@tcp($ip:$port)/$db?$queries 格式即可
}
