FROM golang:1.17-bullseye as builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

# Get dependencies
RUN go mod download -x

# Build
COPY cmd cmd
COPY internal internal
RUN CGO_ENABLED=0 GOOS=linux go build github.com/zigapk/prpo-auth/cmd/auth

#####################
# Running container #
#####################
FROM debian:bullseye

# Get certificates
RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates
RUN update-ca-certificates

WORKDIR /app

COPY configs configs
COPY --from=builder /app/auth .

ENTRYPOINT ["./auth"]
CMD ["serve"]