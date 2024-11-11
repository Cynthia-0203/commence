package rpc

import (
	"fmt"
	"sync"

	"github.com/Cynthia/commence/app/frontend/conf"
	frontendutils "github.com/Cynthia/commence/app/frontend/utils"
	"github.com/Cynthia/commence/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/Cynthia/commence/rpc_gen/kitex_gen/checkout/checkoutservice"
	"github.com/Cynthia/commence/rpc_gen/kitex_gen/order/orderservice"
	"github.com/Cynthia/commence/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/Cynthia/commence/rpc_gen/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
)


var (
	UserClient userservice.Client
	ProductClient productcatalogservice.Client
	CartClient cartservice.Client
	CheckoutClient checkoutservice.Client
	OrderClient orderservice.Client
	once sync.Once
)

func Init() {
	once.Do(func() {
		initUserClient()
		initProductClient()
		initCartClient()
		initOrderClient()
		initCheckoutClient()
	})
}

func initUserClient() {
	var opts []client.Option
	r,err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	
	frontendutils.MustHandleError(err)
	
	opts = append(opts, client.WithResolver(r))
	UserClient,err = userservice.NewClient("user",opts...)
	
	frontendutils.MustHandleError(err)

}

func initProductClient() {
	var opts []client.Option
	r,err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)

	frontendutils.MustHandleError(err)

	opts = append(opts, client.WithResolver(r))
	ProductClient,err = productcatalogservice.NewClient("product", opts...)

	frontendutils.MustHandleError(err)
}

func initCartClient() {
	var opts []client.Option
	r,err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)

	frontendutils.MustHandleError(err)

	opts = append(opts, client.WithResolver(r))
	CartClient,err = cartservice.NewClient("cart", opts...)

	frontendutils.MustHandleError(err)
}

func initCheckoutClient() {
	var opts []client.Option
	r,err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)

	frontendutils.MustHandleError(err)

	opts = append(opts, client.WithResolver(r))
	CheckoutClient,err = checkoutservice.NewClient("checkout", opts...)

	frontendutils.MustHandleError(err)
}

func initOrderClient() {
	var opts []client.Option
	r,err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)

	frontendutils.MustHandleError(err)

	opts = append(opts, client.WithResolver(r))
	OrderClient,err = orderservice.NewClient("order", opts...)

	frontendutils.MustHandleError(err)
	fmt.Println("initOrderClient")
}