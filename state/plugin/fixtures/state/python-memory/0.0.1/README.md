## setup

All of the following commands will be run from the "components-contrib/state/plugin/fixtures/python-memory/0.0.1" directory.

Full documentation for installation can be found here: https://grpc.io/docs/languages/python/quickstart/

### install pip 

```bash
$ sudo apt-get install python3-pip -y
```

### install dependencies

```bash
$ pip3 install -r requirements.txt
```

** grpcio-tools is only required for code generation

## generate

```bash
$ export COMPONENT_FOLDER=`realpath ../../../..`
$ python3 -m grpc_tools.protoc -I $COMPONENT_FOLDER --python_out=. --grpc_python_out=. $COMPONENTS_FOLDER/proto/state.proto
```