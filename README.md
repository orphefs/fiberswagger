# Fiber server with OpenAPI/Swagger documentation

## First steps

To prepare the Swagger middleware,

1. Add comments to your API source code, [See Declarative Comments Format](https://github.com/swaggo/swag#declarative-comments-format).
2. Download [Swag](https://github.com/swaggo/swag) for Go by using:

```sh
go get -u github.com/swaggo/swag/cmd/swag
# 1.16 or newer
go install github.com/swaggo/swag/cmd/swag@latest
```

3. Run the [Swag](https://github.com/swaggo/swag) in your Go project root folder which contains `main.go` file, [Swag](https://github.com/swaggo/swag) will parse comments and generate required files(`docs` folder and `docs/doc.go`).

```sh
swag init
```

Taken from [here](https://raw.githubusercontent.com/gofiber/swagger/main/README.md).

## Compile and run the server

To compile the server, run

```
go install .
```

## Comparison of server frameworks

[here](https://github.com/bradstimpson/fv3vg)
