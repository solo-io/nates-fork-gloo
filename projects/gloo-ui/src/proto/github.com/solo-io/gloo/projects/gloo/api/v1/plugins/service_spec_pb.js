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

var gogoproto_gogo_pb = require('../../../../../../../gogo/protobuf/gogoproto/gogo_pb.js');
var github_com_solo$io_gloo_projects_gloo_api_v1_plugins_rest_rest_pb = require('../../../../../../../../github.com/solo-io/gloo/projects/gloo/api/v1/plugins/rest/rest_pb.js');
var github_com_solo$io_gloo_projects_gloo_api_v1_plugins_grpc_grpc_pb = require('../../../../../../../../github.com/solo-io/gloo/projects/gloo/api/v1/plugins/grpc/grpc_pb.js');
goog.exportSymbol('proto.plugins.gloo.solo.io.ServiceSpec', null, global);

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
proto.plugins.gloo.solo.io.ServiceSpec = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, proto.plugins.gloo.solo.io.ServiceSpec.oneofGroups_);
};
goog.inherits(proto.plugins.gloo.solo.io.ServiceSpec, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.plugins.gloo.solo.io.ServiceSpec.displayName = 'proto.plugins.gloo.solo.io.ServiceSpec';
}
/**
 * Oneof group definitions for this message. Each group defines the field
 * numbers belonging to that group. When of these fields' value is set, all
 * other fields in the group are cleared. During deserialization, if multiple
 * fields are encountered for a group, only the last value seen will be kept.
 * @private {!Array<!Array<number>>}
 * @const
 */
proto.plugins.gloo.solo.io.ServiceSpec.oneofGroups_ = [[1,2]];

/**
 * @enum {number}
 */
proto.plugins.gloo.solo.io.ServiceSpec.PluginTypeCase = {
  PLUGIN_TYPE_NOT_SET: 0,
  REST: 1,
  GRPC: 2
};

/**
 * @return {proto.plugins.gloo.solo.io.ServiceSpec.PluginTypeCase}
 */
proto.plugins.gloo.solo.io.ServiceSpec.prototype.getPluginTypeCase = function() {
  return /** @type {proto.plugins.gloo.solo.io.ServiceSpec.PluginTypeCase} */(jspb.Message.computeOneofCase(this, proto.plugins.gloo.solo.io.ServiceSpec.oneofGroups_[0]));
};



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
proto.plugins.gloo.solo.io.ServiceSpec.prototype.toObject = function(opt_includeInstance) {
  return proto.plugins.gloo.solo.io.ServiceSpec.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.plugins.gloo.solo.io.ServiceSpec} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.plugins.gloo.solo.io.ServiceSpec.toObject = function(includeInstance, msg) {
  var f, obj = {
    rest: (f = msg.getRest()) && github_com_solo$io_gloo_projects_gloo_api_v1_plugins_rest_rest_pb.ServiceSpec.toObject(includeInstance, f),
    grpc: (f = msg.getGrpc()) && github_com_solo$io_gloo_projects_gloo_api_v1_plugins_grpc_grpc_pb.ServiceSpec.toObject(includeInstance, f)
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
 * @return {!proto.plugins.gloo.solo.io.ServiceSpec}
 */
proto.plugins.gloo.solo.io.ServiceSpec.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.plugins.gloo.solo.io.ServiceSpec;
  return proto.plugins.gloo.solo.io.ServiceSpec.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.plugins.gloo.solo.io.ServiceSpec} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.plugins.gloo.solo.io.ServiceSpec}
 */
proto.plugins.gloo.solo.io.ServiceSpec.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new github_com_solo$io_gloo_projects_gloo_api_v1_plugins_rest_rest_pb.ServiceSpec;
      reader.readMessage(value,github_com_solo$io_gloo_projects_gloo_api_v1_plugins_rest_rest_pb.ServiceSpec.deserializeBinaryFromReader);
      msg.setRest(value);
      break;
    case 2:
      var value = new github_com_solo$io_gloo_projects_gloo_api_v1_plugins_grpc_grpc_pb.ServiceSpec;
      reader.readMessage(value,github_com_solo$io_gloo_projects_gloo_api_v1_plugins_grpc_grpc_pb.ServiceSpec.deserializeBinaryFromReader);
      msg.setGrpc(value);
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
proto.plugins.gloo.solo.io.ServiceSpec.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.plugins.gloo.solo.io.ServiceSpec.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.plugins.gloo.solo.io.ServiceSpec} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.plugins.gloo.solo.io.ServiceSpec.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getRest();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      github_com_solo$io_gloo_projects_gloo_api_v1_plugins_rest_rest_pb.ServiceSpec.serializeBinaryToWriter
    );
  }
  f = message.getGrpc();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      github_com_solo$io_gloo_projects_gloo_api_v1_plugins_grpc_grpc_pb.ServiceSpec.serializeBinaryToWriter
    );
  }
};


/**
 * optional rest.plugins.gloo.solo.io.ServiceSpec rest = 1;
 * @return {?proto.rest.plugins.gloo.solo.io.ServiceSpec}
 */
proto.plugins.gloo.solo.io.ServiceSpec.prototype.getRest = function() {
  return /** @type{?proto.rest.plugins.gloo.solo.io.ServiceSpec} */ (
    jspb.Message.getWrapperField(this, github_com_solo$io_gloo_projects_gloo_api_v1_plugins_rest_rest_pb.ServiceSpec, 1));
};


/** @param {?proto.rest.plugins.gloo.solo.io.ServiceSpec|undefined} value */
proto.plugins.gloo.solo.io.ServiceSpec.prototype.setRest = function(value) {
  jspb.Message.setOneofWrapperField(this, 1, proto.plugins.gloo.solo.io.ServiceSpec.oneofGroups_[0], value);
};


proto.plugins.gloo.solo.io.ServiceSpec.prototype.clearRest = function() {
  this.setRest(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.plugins.gloo.solo.io.ServiceSpec.prototype.hasRest = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional grpc.plugins.gloo.solo.io.ServiceSpec grpc = 2;
 * @return {?proto.grpc.plugins.gloo.solo.io.ServiceSpec}
 */
proto.plugins.gloo.solo.io.ServiceSpec.prototype.getGrpc = function() {
  return /** @type{?proto.grpc.plugins.gloo.solo.io.ServiceSpec} */ (
    jspb.Message.getWrapperField(this, github_com_solo$io_gloo_projects_gloo_api_v1_plugins_grpc_grpc_pb.ServiceSpec, 2));
};


/** @param {?proto.grpc.plugins.gloo.solo.io.ServiceSpec|undefined} value */
proto.plugins.gloo.solo.io.ServiceSpec.prototype.setGrpc = function(value) {
  jspb.Message.setOneofWrapperField(this, 2, proto.plugins.gloo.solo.io.ServiceSpec.oneofGroups_[0], value);
};


proto.plugins.gloo.solo.io.ServiceSpec.prototype.clearGrpc = function() {
  this.setGrpc(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.plugins.gloo.solo.io.ServiceSpec.prototype.hasGrpc = function() {
  return jspb.Message.getField(this, 2) != null;
};


goog.object.extend(exports, proto.plugins.gloo.solo.io);
