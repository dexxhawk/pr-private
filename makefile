codegen:
	go tool ogen -clean -config .ogen.yml -target internal/generated/api api/api.yml

style:
	golangci-lint run