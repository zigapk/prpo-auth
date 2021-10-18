bindata:
	@go-bindata -o internal/config/bindata.go -prefix="configs" -pkg=config configs/...

build:
	@go build github.com/zigapk/prpo-auth/cmd/auth