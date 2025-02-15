package rpc

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/client"
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
	r, err := consul.NewConsulResolver("127.0.0.1:8500")
	if err != nil {
		hlog.Fatal(err)
	}
	UserClient, err = userservice.NewClient("user", client.WithResolver(r))
	if err != nil {
		hlog.Fatal(err)
	}
}
