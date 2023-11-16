.PHONY: generate
generate:
	protoc --proto_path api/proto \
		--go_out=pkg/pb --go_opt=paths=source_relative \
		--plugin=protoc-gen-go=$(GOPATH)/bin/protoc-gen-go \
		api/proto/message.proto