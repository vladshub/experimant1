// package: pb
// file: keywee.proto

var keywee_pb = require("./keywee_pb");
var grpc = require("grpc-web-client").grpc;

var API = (function () {
  function API() {}
  API.serviceName = "pb.API";
  return API;
}());

API.Add = {
  methodName: "Add",
  service: API,
  requestStream: false,
  responseStream: false,
  requestType: keywee_pb.Item,
  responseType: keywee_pb.Status
};

API.Read = {
  methodName: "Read",
  service: API,
  requestStream: false,
  responseStream: false,
  requestType: keywee_pb.Item,
  responseType: keywee_pb.Item
};

API.Index = {
  methodName: "Index",
  service: API,
  requestStream: false,
  responseStream: true,
  requestType: keywee_pb.Empty,
  responseType: keywee_pb.Item
};

exports.API = API;

function APIClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

APIClient.prototype.add = function add(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  grpc.unary(API.Add, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          callback(Object.assign(new Error(response.statusMessage), { code: response.status, metadata: response.trailers }), null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
};

APIClient.prototype.read = function read(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  grpc.unary(API.Read, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          callback(Object.assign(new Error(response.statusMessage), { code: response.status, metadata: response.trailers }), null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
};

APIClient.prototype.index = function index(requestMessage, metadata) {
  var listeners = {
    data: [],
    end: [],
    status: []
  };
  var client = grpc.invoke(API.Index, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    onMessage: function (responseMessage) {
      listeners.data.forEach(function (handler) {
        handler(responseMessage);
      });
    },
    onEnd: function (status, statusMessage, trailers) {
      listeners.end.forEach(function (handler) {
        handler();
      });
      listeners.status.forEach(function (handler) {
        handler({ code: status, details: statusMessage, metadata: trailers });
      });
      listeners = null;
    }
  });
  return {
    on: function (type, handler) {
      listeners[type].push(handler);
      return this;
    },
    cancel: function () {
      listeners = null;
      client.close();
    }
  };
};

exports.APIClient = APIClient;

