# Bookstore

## Steps to run the service locally
`go run main.go`

## Swagger
Run `swag init` to generate the Swagger documentation from the codes

## Generate Mock
`go generate ./...`

## Run Test
`go test ./...`
### With Coverage
```
go test ./... -coverprofile=cover.out
go tool cover -html=cover.out -o cover.html
```

## Run Linter
`golangci-lint run`