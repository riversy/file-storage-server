# File Storage (Test Job)  
Here's the Backend Implementation on Golang. 

## Development

The development pre-requirements are the following. 

### Install Protocol Buffers Compiler

```bash
brew install protobuf
```

### Install Go Protocol Buffers Plugin

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go
```

## Update dependencies

Before update the module, please ensure if you use the latest version of the Proto description from the following repo - https://github.com/riversy/file-storage-proto.  

```bash
git submodule update --init --recursive
```

If there are some changes delivered, please regenerate service definitions for Server application.

```bash
protoc --proto_path=proto \
  --go_out=service \
  --go_opt=paths=source_relative \
  --go-grpc_out=service \
  --go-grpc_opt=paths=source_relative \
  proto/file_storage_service.proto
```

## Test 

```bash
go test ./...
```

