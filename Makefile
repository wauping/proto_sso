sso_gen_go:
	protoc -I proto proto/sso/sso.proto --go_out=./gen/sso/go --go_opt=paths=source_relative --go-grpc_out=./gen/sso/go/ --go-grpc_opt=paths=source_relative

sso_gen_ts:
	npx protoc --ts_out ./gen/sso/ts --proto_path proto proto/sso/sso.proto

mdts_gen_go:
	protoc -I proto proto/mdts/mdts.proto --go_out=./gen/mdts/go --go_opt=paths=source_relative --go-grpc_out=./gen/mdts/go/ --go-grpc_opt=paths=source_relative

mdts_gen_ts:
	npx protoc --ts_out ./gen/mdts/ts --proto_path proto proto/mdts/mdts.proto


.PHONY: sso_gen_go sso_gen_ts mdts_gen_go mdts_gen_ts
