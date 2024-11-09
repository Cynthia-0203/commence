# .PHONY: gen-demo-proto
gen-home-proto:
	@cd app/frontend && cwgo server --type HTTP --idl ../../idl/frontend/home.proto --service frontend --module github.com/Cynthia/commence/app/frontend -I ../../
gen-auth_page-proto:
	@cd app/frontend && cwgo server --type HTTP --idl ../../idl/frontend/auth_page.proto --service frontend --module github.com/Cynthia/commence/app/frontend -I ../../
gen-product-proto:
	@cd app/frontend && cwgo server --type HTTP --idl ../../idl/frontend/product_page.proto --service frontend --module github.com/Cynthia/commence/app/frontend -I ../../
gen-category-proto:
	@cd app/frontend && cwgo server --type HTTP --idl ../../idl/frontend/category_page.proto --service frontend --module github.com/Cynthia/commence/app/frontend -I ../../
gen-cart_page-proto:
	@cd app/frontend && cwgo server --type HTTP --idl ../../idl/frontend/cart_page.proto --service frontend --module github.com/Cynthia/commence/app/frontend -I ../../



gen-user-proto-client:
	@cd rpc_gen && cwgo client --type RPC  --service user --module github.com/Cynthia/commence/rpc_gen -I ../idl --idl ../idl/user.proto

gen-user-proto-server:
	@cd app/user && cwgo server --type RPC  --service user --module github.com/Cynthia/commence/app/user --pass "-use github.com/Cynthia/commence/rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/user.proto

gen-product-proto-client:
	@cd rpc_gen && cwgo client --type RPC  --service product --module github.com/Cynthia/commence/rpc_gen -I ../idl --idl ../idl/product.proto

gen-product-proto-server:
	@cd app/product && cwgo server --type RPC  --service product --module github.com/Cynthia/commence/app/product --pass "-use github.com/Cynthia/commence/rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/product.proto

gen-cart-proto:
	@cd rpc_gen && cwgo client --type RPC  --service cart --module github.com/Cynthia/commence/rpc_gen -I ../idl --idl ../idl/cart.proto
	@cd app/cart && cwgo server --type RPC  --service cart --module github.com/Cynthia/commence/app/cart --pass "-use github.com/Cynthia/commence/rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/cart.proto
