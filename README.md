# prpo-auth microservice

This repository contains a source code for user management microservice used in a demo project developed under PRPO subject at UNI LJ.

## Setup

Run a development database via Docker:

```shell
docker run -p 5432:5432 \
    -e POSTGRES_USER=prpo \
    -e POSTGRES_PASSWORD=rootroot \
    -v prpo:/var/lib/postgresql/data \
    -d postgres:14
```

Install dependencies and compile:

```shell
go mod download
make bindata
make build
```

Create signing keys:

```shell
./auth genkeys # this creates signing key pair under `conf` directory
```

Create a user from command line:

```shell
./auth createuser --email some.email@mail.asdf --name username123 --password securepwd
```

Run the server:

```shell
./auth serve
```

System behaviour can be configured through [configs/config.ini](configs/config.ini).

## Documentation

Generate docs using
```bash
swag init -d cmd/auth/,internal/handle
```

The docs are then available at `/docs`.

This microservice exposes its functionality via the following http endpoints:

- `/authorize` POST - Obtain JWT refresh and access tokens either via email-password pair or and old refresh token.
- `/change_password` PUT Change user's password by providing an old password.
- `/users` POST Create a new user (used for signup/register).
- `/users/{uid}` GET Retrieve user info by providing UID from JWT token's subject field.
- `/signing_key` GET Obtain server's public signing key to verify token validity.

## Excercise goal equivalents

The orinal excercise is designed for Java and JDBC. Here are some analogus goals fulfiled by this go implementation:

1. A sample [wiki page](https://github.com/zigapk/prpo-auth/wiki).
2. Dependency management implemented using [go modules](go.mod).
3. Http endpoints defined in [`internal/router/routes.go`](internal/router/routes.go) and implemented in [`internal/handle/`](internal/handle).
4. Model management functions in [`internal/models/`](internal/models).
5. Database migrations defined in [`configs/migrations/`](configs/migrations/).
