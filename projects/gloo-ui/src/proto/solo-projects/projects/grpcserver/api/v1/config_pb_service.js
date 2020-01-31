// package: glooeeapi.solo.io
// file: solo-projects/projects/grpcserver/api/v1/config.proto

var solo_projects_projects_grpcserver_api_v1_config_pb = require("../../../../../solo-projects/projects/grpcserver/api/v1/config_pb");
var grpc = require("@improbable-eng/grpc-web").grpc;

var ConfigApi = (function () {
  function ConfigApi() {}
  ConfigApi.serviceName = "glooeeapi.solo.io.ConfigApi";
  return ConfigApi;
}());

ConfigApi.GetVersion = {
  methodName: "GetVersion",
  service: ConfigApi,
  requestStream: false,
  responseStream: false,
  requestType: solo_projects_projects_grpcserver_api_v1_config_pb.GetVersionRequest,
  responseType: solo_projects_projects_grpcserver_api_v1_config_pb.GetVersionResponse
};

ConfigApi.GetOAuthEndpoint = {
  methodName: "GetOAuthEndpoint",
  service: ConfigApi,
  requestStream: false,
  responseStream: false,
  requestType: solo_projects_projects_grpcserver_api_v1_config_pb.GetOAuthEndpointRequest,
  responseType: solo_projects_projects_grpcserver_api_v1_config_pb.GetOAuthEndpointResponse
};

ConfigApi.GetIsLicenseValid = {
  methodName: "GetIsLicenseValid",
  service: ConfigApi,
  requestStream: false,
  responseStream: false,
  requestType: solo_projects_projects_grpcserver_api_v1_config_pb.GetIsLicenseValidRequest,
  responseType: solo_projects_projects_grpcserver_api_v1_config_pb.GetIsLicenseValidResponse
};

ConfigApi.GetSettings = {
  methodName: "GetSettings",
  service: ConfigApi,
  requestStream: false,
  responseStream: false,
  requestType: solo_projects_projects_grpcserver_api_v1_config_pb.GetSettingsRequest,
  responseType: solo_projects_projects_grpcserver_api_v1_config_pb.GetSettingsResponse
};

ConfigApi.UpdateSettings = {
  methodName: "UpdateSettings",
  service: ConfigApi,
  requestStream: false,
  responseStream: false,
  requestType: solo_projects_projects_grpcserver_api_v1_config_pb.UpdateSettingsRequest,
  responseType: solo_projects_projects_grpcserver_api_v1_config_pb.UpdateSettingsResponse
};

ConfigApi.UpdateSettingsYaml = {
  methodName: "UpdateSettingsYaml",
  service: ConfigApi,
  requestStream: false,
  responseStream: false,
  requestType: solo_projects_projects_grpcserver_api_v1_config_pb.UpdateSettingsYamlRequest,
  responseType: solo_projects_projects_grpcserver_api_v1_config_pb.UpdateSettingsResponse
};

ConfigApi.ListNamespaces = {
  methodName: "ListNamespaces",
  service: ConfigApi,
  requestStream: false,
  responseStream: false,
  requestType: solo_projects_projects_grpcserver_api_v1_config_pb.ListNamespacesRequest,
  responseType: solo_projects_projects_grpcserver_api_v1_config_pb.ListNamespacesResponse
};

ConfigApi.GetPodNamespace = {
  methodName: "GetPodNamespace",
  service: ConfigApi,
  requestStream: false,
  responseStream: false,
  requestType: solo_projects_projects_grpcserver_api_v1_config_pb.GetPodNamespaceRequest,
  responseType: solo_projects_projects_grpcserver_api_v1_config_pb.GetPodNamespaceResponse
};

exports.ConfigApi = ConfigApi;

function ConfigApiClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

ConfigApiClient.prototype.getVersion = function getVersion(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(ConfigApi.GetVersion, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

ConfigApiClient.prototype.getOAuthEndpoint = function getOAuthEndpoint(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(ConfigApi.GetOAuthEndpoint, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

ConfigApiClient.prototype.getIsLicenseValid = function getIsLicenseValid(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(ConfigApi.GetIsLicenseValid, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

ConfigApiClient.prototype.getSettings = function getSettings(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(ConfigApi.GetSettings, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

ConfigApiClient.prototype.updateSettings = function updateSettings(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(ConfigApi.UpdateSettings, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

ConfigApiClient.prototype.updateSettingsYaml = function updateSettingsYaml(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(ConfigApi.UpdateSettingsYaml, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

ConfigApiClient.prototype.listNamespaces = function listNamespaces(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(ConfigApi.ListNamespaces, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

ConfigApiClient.prototype.getPodNamespace = function getPodNamespace(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(ConfigApi.GetPodNamespace, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

exports.ConfigApiClient = ConfigApiClient;

