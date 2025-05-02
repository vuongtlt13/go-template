gen-proto:
	protoc \
	  -I=proto \
	  -I=third_party \
	  --proto_path=proto \
	  --go_out=pb \
	  --go_opt paths=source_relative \
	  --go-grpc_out=pb \
	  --go-grpc_opt paths=source_relative \
	  --connect-go_out=pb \
	  --connect-go_opt paths=source_relative \
	  --openapiv2_out=api \
      --openapiv2_opt logtostderr=true \
      --openapiv2_opt generate_unbound_methods=true,allow_merge=true,merge_file_name=api \
	  proto/auth/*.proto proto/health/*.proto

