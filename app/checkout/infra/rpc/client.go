package rpc

import (
	"github.com/cloudwego/kitex/client"
	"github.com/feeeeling/eMall/app/checkout/conf"
	checkoututils "github.com/feeeeling/eMall/app/checkout/utils"
	"github.com/feeeeling/eMall/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/feeeeling/eMall/rpc_gen/kitex_gen/payment/paymentservice"
	"github.com/feeeeling/eMall/rpc_gen/kitex_gen/order/orderservice"
	"github.com/feeeeling/eMall/rpc_gen/kitex_gen/product/productcatalogservice"
	consul "github.com/kitex-contrib/registry-consul"
	"sync"
)

var (
	CartClient    cartservice.Client
	ProductClient productcatalogservice.Client
	PaymentClient paymentservice.Client
	OrderClient orderservice.Client

	err  error
	once sync.Once
)

func InitClient() {
	once.Do(func() {
		initCartClient()
		initProductClient()
		initPaymentClient()
		initOrderClient()
	})
}

func initCartClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	checkoututils.MustHandleError(err)
	CartClient, err = cartservice.NewClient("cart", client.WithResolver(r))
	checkoututils.MustHandleError(err)
}

func initProductClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	checkoututils.MustHandleError(err)
	ProductClient, err = productcatalogservice.NewClient("product", client.WithResolver(r))
	checkoututils.MustHandleError(err)
}

func initPaymentClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	checkoututils.MustHandleError(err)
	PaymentClient, err = paymentservice.NewClient("payment", client.WithResolver(r))
	checkoututils.MustHandleError(err)
}

func initOrderClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	checkoututils.MustHandleError(err)
	OrderClient, err = orderservice.NewClient("order", client.WithResolver(r))
	checkoututils.MustHandleError(err)
}
