package rpc

import (
	"github.com/cloudwego/kitex/client"
	"github.com/feeeeling/eMall/app/frontend/conf"
	frontendUtils "github.com/feeeeling/eMall/app/frontend/utils"
	"github.com/feeeeling/eMall/rpc_gen/kitex_gen/user/userservice"
	consul "github.com/kitex-contrib/registry-consul"
	"sync"
)

var (
	UserClient userservice.Client

	Once sync.Once
)

func Init() {
	Once.Do(func() {
		initUserClient()
	})
}

func initUserClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	frontendUtils.MustHandleError(err)
	UserClient, err = userservice.NewClient("user", client.WithResolver(r))
	frontendUtils.MustHandleError(err)
}
