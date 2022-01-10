// GENERATED CODE -- DO NOT EDIT!

// Original file comments:
// protobuf specification https://developers.google.com/protocol-buffers/docs/proto3
'use strict';
var grpc = require('@grpc/grpc-js');
var proto_state_pb = require('../proto/state_pb.js');

function serialize_proto_DeleteRequest(arg) {
  if (!(arg instanceof proto_state_pb.DeleteRequest)) {
    throw new Error('Expected argument of type proto.DeleteRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_proto_DeleteRequest(buffer_arg) {
  return proto_state_pb.DeleteRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_proto_Empty(arg) {
  if (!(arg instanceof proto_state_pb.Empty)) {
    throw new Error('Expected argument of type proto.Empty');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_proto_Empty(buffer_arg) {
  return proto_state_pb.Empty.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_proto_GetRequest(arg) {
  if (!(arg instanceof proto_state_pb.GetRequest)) {
    throw new Error('Expected argument of type proto.GetRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_proto_GetRequest(buffer_arg) {
  return proto_state_pb.GetRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_proto_GetResponse(arg) {
  if (!(arg instanceof proto_state_pb.GetResponse)) {
    throw new Error('Expected argument of type proto.GetResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_proto_GetResponse(buffer_arg) {
  return proto_state_pb.GetResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_proto_Metadata(arg) {
  if (!(arg instanceof proto_state_pb.Metadata)) {
    throw new Error('Expected argument of type proto.Metadata');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_proto_Metadata(buffer_arg) {
  return proto_state_pb.Metadata.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_proto_SetRequest(arg) {
  if (!(arg instanceof proto_state_pb.SetRequest)) {
    throw new Error('Expected argument of type proto.SetRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_proto_SetRequest(buffer_arg) {
  return proto_state_pb.SetRequest.deserializeBinary(new Uint8Array(buffer_arg));
}


var StateService = exports.StateService = {
  init: {
    path: '/proto.State/Init',
    requestStream: false,
    responseStream: false,
    requestType: proto_state_pb.Metadata,
    responseType: proto_state_pb.Empty,
    requestSerialize: serialize_proto_Metadata,
    requestDeserialize: deserialize_proto_Metadata,
    responseSerialize: serialize_proto_Empty,
    responseDeserialize: deserialize_proto_Empty,
  },
  get: {
    path: '/proto.State/Get',
    requestStream: false,
    responseStream: false,
    requestType: proto_state_pb.GetRequest,
    responseType: proto_state_pb.GetResponse,
    requestSerialize: serialize_proto_GetRequest,
    requestDeserialize: deserialize_proto_GetRequest,
    responseSerialize: serialize_proto_GetResponse,
    responseDeserialize: deserialize_proto_GetResponse,
  },
  set: {
    path: '/proto.State/Set',
    requestStream: false,
    responseStream: false,
    requestType: proto_state_pb.SetRequest,
    responseType: proto_state_pb.Empty,
    requestSerialize: serialize_proto_SetRequest,
    requestDeserialize: deserialize_proto_SetRequest,
    responseSerialize: serialize_proto_Empty,
    responseDeserialize: deserialize_proto_Empty,
  },
  delete: {
    path: '/proto.State/Delete',
    requestStream: false,
    responseStream: false,
    requestType: proto_state_pb.DeleteRequest,
    responseType: proto_state_pb.Empty,
    requestSerialize: serialize_proto_DeleteRequest,
    requestDeserialize: deserialize_proto_DeleteRequest,
    responseSerialize: serialize_proto_Empty,
    responseDeserialize: deserialize_proto_Empty,
  },
  ping: {
    path: '/proto.State/Ping',
    requestStream: false,
    responseStream: false,
    requestType: proto_state_pb.Empty,
    responseType: proto_state_pb.Empty,
    requestSerialize: serialize_proto_Empty,
    requestDeserialize: deserialize_proto_Empty,
    responseSerialize: serialize_proto_Empty,
    responseDeserialize: deserialize_proto_Empty,
  },
};

exports.StateClient = grpc.makeGenericClientConstructor(StateService);
