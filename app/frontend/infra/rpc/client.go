package rpc

import (
	"sync"

	"github.com/Cynthia/commence/app/frontend/conf"
	frontendutils "github.com/Cynthia/commence/app/frontend/utils"
	"github.com/Cynthia/commence/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/Cynthia/commence/rpc_gen/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
)


var (
	UserClient userservice.Client
	ProductClient productcatalogservice.Client
	once sync.Once
)

func Init() {
	once.Do(func() {
		initUserClient()
		initProductClient()
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