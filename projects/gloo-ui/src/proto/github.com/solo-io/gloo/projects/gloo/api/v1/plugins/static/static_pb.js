/* eslint-disable */
/**
 * @fileoverview
 * @enhanceable
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!

var jspb = require('google-protobuf');
var goog = jspb;
var global = Function('return this')();

var gogoproto_gogo_pb = require('../../../../../../../../gogo/protobuf/gogoproto/gogo_pb.js');
var github_com_solo$io_gloo_projects_gloo_api_v1_plugins_service_spec_pb = require('../../../../../../../../../github.com/solo-io/gloo/projects/gloo/api/v1/plugins/service_spec_pb.js');
goog.exportSymbol('proto.static.plugins.gloo.solo.io.Host', null, global);
goog.exportSymbol('proto.static.plugins.gloo.solo.io.UpstreamSpec', null, global);

/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.static.plugins.gloo.solo.io.UpstreamSpec = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.static.plugins.gloo.solo.io.UpstreamSpec.repeatedFields_, null);
};
goog.inherits(proto.static.plugins.gloo.solo.io.UpstreamSpec, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.static.plugins.gloo.solo.io.UpstreamSpec.displayName = 'proto.static.plugins.gloo.solo.io.UpstreamSpec';
}
/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.static.plugins.gloo.solo.io.UpstreamSpec.repeatedFields_ = [1];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.static.plugins.gloo.solo.io.UpstreamSpec.prototype.toObject = function(opt_includeInstance) {
  return proto.static.plugins.gloo.solo.io.UpstreamSpec.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.static.plugins.gloo.solo.io.UpstreamSpec} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.static.plugins.gloo.solo.io.UpstreamSpec.toObject = function(includeInstance, msg) {
  var f, obj = {
    hostsList: jspb.Message.toObjectList(msg.getHostsList(),
    proto.static.plugins.gloo.solo.io.Host.toObject, includeInstance),
    useTls: jspb.Message.getFieldWithDefault(msg, 3, false),
    serviceSpec: (f = msg.getServiceSpec()) && github_com_solo$io_gloo_projects_gloo_api_v1_plugins_service_spec_pb.ServiceSpec.toObject(includeInstance, f),
    useHttp2: jspb.Message.getFieldWithDefault(msg, 6, false)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.static.plugins.gloo.solo.io.UpstreamSpec}
 */
proto.static.plugins.gloo.solo.io.UpstreamSpec.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.static.plugins.gloo.solo.io.UpstreamSpec;
  return proto.static.plugins.gloo.solo.io.UpstreamSpec.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.static.plugins.gloo.solo.io.UpstreamSpec} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.static.plugins.gloo.solo.io.UpstreamSpec}
 */
proto.static.plugins.gloo.solo.io.UpstreamSpec.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.static.plugins.gloo.solo.io.Host;
      reader.readMessage(value,proto.static.plugins.gloo.solo.io.Host.deserializeBinaryFromReader);
      msg.addHosts(value);
      break;
    case 3:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setUseTls(value);
      break;
    case 5:
      var value = new github_com_solo$io_gloo_projects_gloo_api_v1_plugins_service_spec_pb.ServiceSpec;
      reader.readMessage(value,github_com_solo$io_gloo_projects_gloo_api_v1_plugins_service_spec_pb.ServiceSpec.deserializeBinaryFromReader);
      msg.setServiceSpec(value);
      break;
    case 6:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setUseHttp2(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.static.plugins.gloo.solo.io.UpstreamSpec.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.static.plugins.gloo.solo.io.UpstreamSpec.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.static.plugins.gloo.solo.io.UpstreamSpec} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.static.plugins.gloo.solo.io.UpstreamSpec.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getHostsList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      1,
      f,
      proto.static.plugins.gloo.solo.io.Host.serializeBinaryToWriter
    );
  }
  f = message.getUseTls();
  if (f) {
    writer.writeBool(
      3,
      f
    );
  }
  f = message.getServiceSpec();
  if (f != null) {
    writer.writeMessage(
      5,
      f,
      github_com_solo$io_gloo_projects_gloo_api_v1_plugins_service_spec_pb.ServiceSpec.serializeBinaryToWriter
    );
  }
  f = message.getUseHttp2();
  if (f) {
    writer.writeBool(
      6,
      f
    );
  }
};


/**
 * repeated Host hosts = 1;
 * @return {!Array<!proto.static.plugins.gloo.solo.io.Host>}
 */
proto.static.plugins.gloo.solo.io.UpstreamSpec.prototype.getHostsList = function() {
  return /** @type{!Array<!proto.static.plugins.gloo.solo.io.Host>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.static.plugins.gloo.solo.io.Host, 1));
};


/** @param {!Array<!proto.static.plugins.gloo.solo.io.Host>} value */
proto.static.plugins.gloo.solo.io.UpstreamSpec.prototype.setHostsList = function(value) {
  jspb.Message.setRepeatedWrapperField(this, 1, value);
};


/**
 * @param {!proto.static.plugins.gloo.solo.io.Host=} opt_value
 * @param {number=} opt_index
 * @return {!proto.static.plugins.gloo.solo.io.Host}
 */
proto.static.plugins.gloo.solo.io.UpstreamSpec.prototype.addHosts = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 1, opt_value, proto.static.plugins.gloo.solo.io.Host, opt_index);
};


proto.static.plugins.gloo.solo.io.UpstreamSpec.prototype.clearHostsList = function() {
  this.setHostsList([]);
};


/**
 * optional bool use_tls = 3;
 * Note that Boolean fields may be set to 0/1 when serialized from a Java server.
 * You should avoid comparisons like {@code val === true/false} in those cases.
 * @return {boolean}
 */
proto.static.plugins.gloo.solo.io.UpstreamSpec.prototype.getUseTls = function() {
  return /** @type {boolean} */ (jspb.Message.getFieldWithDefault(this, 3, false));
};


/** @param {boolean} value */
proto.static.plugins.gloo.solo.io.UpstreamSpec.prototype.setUseTls = function(value) {
  jspb.Message.setProto3BooleanField(this, 3, value);
};


/**
 * optional plugins.gloo.solo.io.ServiceSpec service_spec = 5;
 * @return {?proto.plugins.gloo.solo.io.ServiceSpec}
 */
proto.static.plugins.gloo.solo.io.UpstreamSpec.prototype.getServiceSpec = function() {
  return /** @type{?proto.plugins.gloo.solo.io.ServiceSpec} */ (
    jspb.Message.getWrapperField(this, github_com_solo$io_gloo_projects_gloo_api_v1_plugins_service_spec_pb.ServiceSpec, 5));
};


/** @param {?proto.plugins.gloo.solo.io.ServiceSpec|undefined} value */
proto.static.plugins.gloo.solo.io.UpstreamSpec.prototype.setServiceSpec = function(value) {
  jspb.Message.setWrapperField(this, 5, value);
};


proto.static.plugins.gloo.solo.io.UpstreamSpec.prototype.clearServiceSpec = function() {
  this.setServiceSpec(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.static.plugins.gloo.solo.io.UpstreamSpec.prototype.hasServiceSpec = function() {
  return jspb.Message.getField(this, 5) != null;
};


/**
 * optional bool use_http2 = 6;
 * Note that Boolean fields may be set to 0/1 when serialized from a Java server.
 * You should avoid comparisons like {@code val === true/false} in those cases.
 * @return {boolean}
 */
proto.static.plugins.gloo.solo.io.UpstreamSpec.prototype.getUseHttp2 = function() {
  return /** @type {boolean} */ (jspb.Message.getFieldWithDefault(this, 6, false));
};


/** @param {boolean} value */
proto.static.plugins.gloo.solo.io.UpstreamSpec.prototype.setUseHttp2 = function(value) {
  jspb.Message.setProto3BooleanField(this, 6, value);
};



/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.static.plugins.gloo.solo.io.Host = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.static.plugins.gloo.solo.io.Host, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.static.plugins.gloo.solo.io.Host.displayName = 'proto.static.plugins.gloo.solo.io.Host';
}


if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.static.plugins.gloo.solo.io.Host.prototype.toObject = function(opt_includeInstance) {
  return proto.static.plugins.gloo.solo.io.Host.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.static.plugins.gloo.solo.io.Host} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.static.plugins.gloo.solo.io.Host.toObject = function(includeInstance, msg) {
  var f, obj = {
    addr: jspb.Message.getFieldWithDefault(msg, 1, ""),
    port: jspb.Message.getFieldWithDefault(msg, 2, 0)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.static.plugins.gloo.solo.io.Host}
 */
proto.static.plugins.gloo.solo.io.Host.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.static.plugins.gloo.solo.io.Host;
  return proto.static.plugins.gloo.solo.io.Host.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.static.plugins.gloo.solo.io.Host} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.static.plugins.gloo.solo.io.Host}
 */
proto.static.plugins.gloo.solo.io.Host.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setAddr(value);
      break;
    case 2:
      var value = /** @type {number} */ (reader.readUint32());
      msg.setPort(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.static.plugins.gloo.solo.io.Host.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.static.plugins.gloo.solo.io.Host.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.static.plugins.gloo.solo.io.Host} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.static.plugins.gloo.solo.io.Host.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getAddr();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getPort();
  if (f !== 0) {
    writer.writeUint32(
      2,
      f
    );
  }
};


/**
 * optional string addr = 1;
 * @return {string}
 */
proto.static.plugins.gloo.solo.io.Host.prototype.getAddr = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/** @param {string} value */
proto.static.plugins.gloo.solo.io.Host.prototype.setAddr = function(value) {
  jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional uint32 port = 2;
 * @return {number}
 */
proto.static.plugins.gloo.solo.io.Host.prototype.getPort = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/** @param {number} value */
proto.static.plugins.gloo.solo.io.Host.prototype.setPort = function(value) {
  jspb.Message.setProto3IntField(this, 2, value);
};


goog.object.extend(exports, proto.static.plugins.gloo.solo.io);
