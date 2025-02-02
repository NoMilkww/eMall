.PHONY: gen-server
gen-server:
	@cd app/${svc} && cwgo server --type RPC --service ${svc} --module github.com/feeeeling/eMall/app/${svc} -I ../../idl --idl ../../idl/${svc}.proto

.PHONY: gen-frontend-home
gen-frontend-home:
	@cd app/frontend && cwgo server --type HTTP --service frontend --module github.com/feeeeling/eMall/app/frontend -I ../../idl --idl ../../idl/frontend/home.proto

.PHONY: gen-frontend-auth
gen-frontend-auth:
	@cd app/frontend && cwgo server --type HTTP --service frontend --module github.com/feeeeling/eMall/app/frontend -I ../../idl --idl ../../idl/frontend/auth_page.proto