export PATH=$PATH:$(pwd)
go build
go generate
go test ./...
