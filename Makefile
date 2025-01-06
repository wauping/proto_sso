gen_go:
	protoc -I proto proto/sso/sso.proto --go_out=./gen/go --go_opt=paths=source_relative --go-grpc_out=./gen/go/ --go-grpc_opt=paths=source_relative

gen_ts:
	protoc -I proto proto/sso/sso.proto --ts_out=./gen/ts
.PHONY: gen_go gen_ts
