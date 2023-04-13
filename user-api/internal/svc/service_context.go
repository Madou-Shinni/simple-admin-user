package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"simple-admin-user-api/internal/config"
	i18n2 "simple-admin-user-api/internal/i18n"
	"simple-admin-user-api/internal/middleware"
	"simple-admin-user-rpc/userclient"

	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/casbin/casbin/v2"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config    config.Config
	Casbin    *casbin.Enforcer
	Authority rest.Middleware
	Trans     *i18n.Translator
	UserRpc   userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {

	rds := redis.MustNewRedis(c.RedisConf)

	cbn := c.CasbinConf.MustNewCasbinWithRedisWatcher(c.DatabaseConf.Type, c.DatabaseConf.GetDSN(), c.RedisConf)

	trans := i18n.NewTranslator(i18n2.LocaleFS)

	return &ServiceContext{
		Config:    c,
		Authority: middleware.NewAuthorityMiddleware(cbn, rds, trans).Handle,
		Trans:     trans,
		UserRpc:   userclient.NewUser(zrpc.NewClientIfEnable(c.UserRpc)),
	}
}
