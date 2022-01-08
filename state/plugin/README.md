# plugin

Provides a GRPC and RPC plugin for specifying DAPR state. The implementation utilizes the [go-plugin](https://github.com/hashicorp/go-plugin) package to manage the grpc client and server. 

The proto/state.proto file contains the protocol buffer definition for the grpc plugins.

The (shared) folder contains components that are shared between client and plugin. These could be moved into a sdk along with the (proto) folder at a later date. 

## reference

For more information on generating go code from proto files see the documentation [here](https://developers.google.com/protocol-buffers/docs/reference/go-generated#package)

## updating the protocol

You *MUST* have the following protoc plugins installed, along with `protoc`. Installation instructions are located [here](https://grpc.io/docs/languages/go/quickstart/)

```bash
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
```

If you update the protocol buffers file you can regenerate the server and client proxies using the below commands. 

> go

```bash
# from the 'plugin' directory
$ go generate
```

## configuration

```yaml
apiVersion: dapr.io/v1alpha1
kind: Component
metadata:
  name: statestore
spec:
  type: state.plugin
  version: v1
  metadata:
  - name: plugin.basedir
    value: ./plugins
  - name: plugin.version
    value: 0.0.1
  - name: plugin.name
    value: go-memory
  - name: plugin.runner
    value: exec
  - name: message
    value: hello
```