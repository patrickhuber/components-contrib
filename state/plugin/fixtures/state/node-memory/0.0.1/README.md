# node-memory state plugin

The node-memory state plugin provides an example nodejs plugin that implements the state api with a nodejs application. 

## run instructions

To run this plugin, you will need nodejs installed. The base image for dapr-dev development container doesn't contain this depdendency so the test will fail if you run without node installed. 

To install node follow these steps:

Install nvm

```bash
sudo curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.1/install.sh | bash
```

Install latest version of node

```bash
nvm install node
```

You need to install the dependencies. The `npm` commands must be run from the directory that contains the package.json file. Switch to the node-memory/0.0.1 folder. 

```
cd state/plugin/fixtures/state/node-memory/0.0.1
npm install
```

## generate

To regenerate the node proxies run the following command from state/plugins/fixtures/state/node-memory/0.0.1 folder. For more information about javascript code generation, see the following [protoc-gen-grpc](https://www.npmjs.com/package/protoc-gen-grpc) link. 

There are two grpc packages for node. grpc-js and grpc. They are binary compatible, but grpc-js is pure js and doesn't have c++ dependencies that decrease its portability. 

Additional information about protocol bufferes in javascript can be found here: [protocol buffers link](https://developers.google.com/protocol-buffers/docs/reference/javascript-generated#invocation).

```bash
export COMPONENT_FOLDER=`realpath ../../../..`

npx grpc_tools_node_protoc \
  --js_out=import_style=commonjs,binary:./ \
  --grpc_out=grpc_js:./ \
  -I=${COMPONENT_FOLDER} ${COMPONENT_FOLDER}/proto/state.proto
```

This will produce a proto folder under the version root. 