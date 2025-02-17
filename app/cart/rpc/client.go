package rpc

import (
	"github.com/cloudwego/kitex/client"
	"github.com/feeeeling/eMall/app/cart/conf"
	cartUtils "github.com/feeeeling/eMall/app/cart/utils"
	"github.com/feeeeling/eMall/rpc_gen/kitex_gen/product/productcatalogservice"
	consul "github.com/kitex-contrib/registry-consul"
	"sync"
)

var (
	ProductClient productcatalogservice.Client
	Once          sync.Once
)

func InitClient() {
	Once.Do(func() {
		initProductClient()
	})
}

func initProductClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	cartUtils.MustHandleError(err)
	opts = append(opts, client.WithResolver(r))
	ProductClient, err = productcatalogservice.NewClient("product", opts...)
	cartUtils.MustHandleError(err)
}
