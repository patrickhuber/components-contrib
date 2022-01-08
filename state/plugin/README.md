# plugin

Provides a GRPC and RPC plugin for specifying DAPR state. 

## reference

For more information on generating go code from proto files see the documentation [here](https://developers.google.com/protocol-buffers/docs/reference/go-generated#package)

## install dependencies

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```

## updating the protocol

If you update the protocol buffers file you can regenerate the server and client proxies using the below commands. 

> go

```bash
# from the 'plugin' directory
$ protoc -I=. --go_out=plugins=grpc:. ./state.proto 
mv github.com/dapr/components-contrib/state/plugin/state.pb.go ./state.pb.go
rm -rf ./github.com

```

or 

```powershell
#? currently broken on windows?
protoc -I="." --go-grpc_out="." ./state.proto
```

```
'protoc-gen-go-grpc' is not recognized as an internal or external command,
operable program or batch file.
--go-grpc_out: protoc-gen-go-grpc: Plugin failed with status code 1.
```

> python

```bash
$ python -m grpc_tools.protoc -I ./proto/ --python_out=./plugin-python/ --grpc_python_out=./plugin-python/ ./proto/kv.proto
```

## configuration

