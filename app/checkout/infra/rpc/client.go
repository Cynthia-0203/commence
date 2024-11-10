package rpc

import (
	"sync"

	"github.com/Cynthia/commence/app/checkout/conf"
	"github.com/Cynthia/commence/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/Cynthia/commence/rpc_gen/kitex_gen/payment/paymentservice"
	"github.com/Cynthia/commence/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/kitex/client"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/transport"
	consul "github.com/kitex-contrib/registry-consul"
)

var (
	CartClient cartservice.Client
	ProductClient productcatalogservice.Client
	PaymentClient paymentservice.Client
	once sync.Once
)

func Init() {
	once.Do(func() {
		initCartClient()
		initPaymentClient()
		initProductClient()
	})
}

func initCartClient(){
	var opts []client.Option
	r,err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	if err != nil {
		panic(err)
	}

	opts = append(opts, client.WithResolver(r))
	opts = append(opts, 
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: conf.GetConf().Kitex.Service,
		}),
		client.WithTransportProtocol(transport.GRPC),
		client.WithMetaHandler(transmeta.ClientHTTP2Handler),
	)
	CartClient,err = cartservice.NewClient("cart", opts...)
	if err != nil {
		panic(err)
	}
}

func initProductClient(){
	var opts []client.Option
	r,err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	if err != nil {
		panic(err)
	}

	opts = append(opts, client.WithResolver(r))
	opts = append(opts, 
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: conf.GetConf().Kitex.Service,
		}),
		client.WithTransportProtocol(transport.GRPC),
		client.WithMetaHandler(transmeta.ClientHTTP2Handler),
	)
	ProductClient,err = productcatalogservice.NewClient("product", opts...)
	if err != nil {
		panic(err)
	}
}

func initPaymentClient(){
	var opts []client.Option
	r,err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	if err != nil {
		panic(err)
	}

	opts = append(opts, client.WithResolver(r))
	opts = append(opts, 
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: conf.GetConf().Kitex.Service,
		}),
		client.WithTransportProtocol(transport.GRPC),
		client.WithMetaHandler(transmeta.ClientHTTP2Handler),
	)
	PaymentClient,err = paymentservice.NewClient("payment", opts...)
	if err != nil {
		panic(err)
	}
}