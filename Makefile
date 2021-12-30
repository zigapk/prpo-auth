bindata:
	@go-bindata -o internal/config/bindata.go -prefix="configs" -pkg=config configs/...

docs:
	@swag fmt -d cmd/auth/,internal/handle && swag init -d cmd/auth/,internal/handle

build:
	@go build github.com/zigapk/prpo-auth/cmd/auth