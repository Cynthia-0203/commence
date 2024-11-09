package rpc

import (
	"sync"

	"github.com/Cynthia/commence/app/cart/conf"
	cartutils "github.com/Cynthia/commence/app/cart/utils"
	"github.com/Cynthia/commence/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
)


var (
	ProductClient productcatalogservice.Client
	once sync.Once
)

func Init() {
	once.Do(func() {
		initProductClient()
	})
}


func initProductClient() {
	var opts []client.Option
	r,err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])

	cartutils.MustHandleError(err)

	opts = append(opts, client.WithResolver(r))
	ProductClient,err = productcatalogservice.NewClient("product", opts...)

	cartutils.MustHandleError(err)
}