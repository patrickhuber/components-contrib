var messages = require('./proto/state_pb');
var services = require('./proto/state_grpc_pb');
var grpc = require('@grpc/grpc-js');

var data = new Object();

/**
 * init request handler. Initializes the data in the plugin and returns an Empty response. 
 * @param {EventEmitter} call Call object for the handler to process
 * @param {function(Error, feature)} callback Response callback
 */
function init(_, callback){
  console.log("init called");
  // clear the data
  data = {};
  callback(null, new messages.Empty());
}

/**
 * get request handler. Gets a request with a GetRequest and returns a GetResponse 
 * with the value that matches the given key.
 * @param {EventEmitter} call Call object for the handler to process
 * @param {function(Error, feature)} callback Response callback
 */
function get(call, callback){  
  console.log("get called");
  var key = call.request.getKey();
  var response = new messages.GetResponse();
  if (!(key in data)){
    callback(null, response);
    return
  }
  response.setData(data[key]);
  callback(null, response);
}

/**
 * set request handler. Sets the data in the dictionary with the given value 
 * with the value that matches the given key.
 * @param {EventEmitter} call Call object for the handler to process
 * @param {function(Error, feature)} callback Response callback
 */
function set(call, callback){
  console.log("set called");
  var key = call.request.getKey();
  data[key] = call.request.getValue();
  callback(null, new messages.Empty());
}

/**
 * delete request handler. Initializes the data in the plugin and returns an Empty response. 
 * @param {EventEmitter} call Call object for the handler to process
 * @param {function(Error, feature)} callback Response callback
 */
function del(call, callback){
  console.log("del called");
  delete data[call.request.getKey()];
  callback(null, new messages.Empty());
}


/**
 * ping request handler. A  
 * @param {EventEmitter} call Call object for the handler to process
 * @param {function(Error, feature)} callback Response callback
 */
function ping(_, callback){  
  console.log("ping called");
  callback(null, new messages.Empty());
}

/**
 * Get a new server with the handler functions in this file bound to the methods
 * it serves.
 * @return {Server} The new server object
 */
 function getServer() {
    var server = new grpc.Server();
    server.addService(services.StateService, {
      init: init,
      get: get,
      set: set,
      delete: del,
      ping: ping
    });
    return server;
  }
  
  if (require.main === module) {
    // If this is run as a script, start a server on an unused port
    var stateServer = getServer();
    stateServer.bindAsync('0.0.0.0:0', grpc.ServerCredentials.createInsecure(), (_, port) => {
        stateServer.start();
        console.log("1|1|tcp|127.0.0.1:%d|grpc", port);    
    });
  }
  
  exports.getServer = getServer;