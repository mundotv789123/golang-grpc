### comandos

ref: https://grpc.io/docs/languages/go/quickstart/
```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
export PATH="$PATH:$(go env GOPATH)/bin"
```

```bash
cd pb/
protoc --go_out=. --go-grpc_out=. user.proto
```