.PHONY: gen-server
gen-server:
	@cd app/${svc} && cwgo server --type RPC --service ${svc} --module github.com/feeeeling/eMall/app/${svc} -I ../../idl --idl ../../idl/${svc}.proto 