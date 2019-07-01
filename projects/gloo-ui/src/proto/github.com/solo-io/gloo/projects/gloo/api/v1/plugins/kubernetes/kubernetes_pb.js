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
var github_com_solo$io_gloo_projects_gloo_api_v1_plugins_subset_spec_pb = require('../../../../../../../../../github.com/solo-io/gloo/projects/gloo/api/v1/plugins/subset_spec_pb.js');
goog.exportSymbol('proto.kubernetes.plugins.gloo.solo.io.UpstreamSpec', null, global);

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
proto.kubernetes.plugins.gloo.solo.io.UpstreamSpec = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.kubernetes.plugins.gloo.solo.io.UpstreamSpec, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.kubernetes.plugins.gloo.solo.io.UpstreamSpec.displayName = 'proto.kubernetes.plugins.gloo.solo.io.UpstreamSpec';
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
proto.kubernetes.plugins.gloo.solo.io.UpstreamSpec.prototype.toObject = function(opt_includeInstance) {
  return proto.kubernetes.plugins.gloo.solo.io.UpstreamSpec.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.kubernetes.plugins.gloo.solo.io.UpstreamSpec} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.kubernetes.plugins.gloo.solo.io.UpstreamSpec.toObject = function(includeInstance, msg) {
  var f, obj = {
    serviceName: jspb.Message.getFieldWithDefault(msg, 1, ""),
    serviceNamespace: jspb.Message.getFieldWithDefault(msg, 2, ""),
    servicePort: jspb.Message.getFieldWithDefault(msg, 3, 0),
    selectorMap: (f = msg.getSelectorMap()) ? f.toObject(includeInstance, undefined) : [],
    serviceSpec: (f = msg.getServiceSpec()) && github_com_solo$io_gloo_projects_gloo_api_v1_plugins_service_spec_pb.ServiceSpec.toObject(includeInstance, f),
    subsetSpec: (f = msg.getSubsetSpec()) && github_com_solo$io_gloo_projects_gloo_api_v1_plugins_subset_spec_pb.SubsetSpec.toObject(includeInstance, f)
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
 * @return {!proto.kubernetes.plugins.gloo.solo.io.UpstreamSpec}
 */
proto.kubernetes.plugins.gloo.solo.io.UpstreamSpec.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.kubernetes.plugins.gloo.solo.io.UpstreamSpec;
  return proto.kubernetes.plugins.gloo.solo.io.UpstreamSpec.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.kubernetes.plugins.gloo.solo.io.UpstreamSpec} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.kubernetes.plugins.gloo.solo.io.UpstreamSpec}
 */
proto.kubernetes.plugins.gloo.solo.io.UpstreamSpec.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setServiceName(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setServiceNamespace(value);
      break;
    case 3:
      var value = /** @type {number} */ (reader.readUint32());
      msg.setServicePort(value);
      break;
    case 4:
      var value = msg.getSelectorMap();
      reader.readMessage(value, function(message, reader) {
        jspb.Map.deserializeBinary(message, reader, jspb.BinaryReader.prototype.readString, jspb.BinaryReader.prototype.readString, null, "");
         });
      break;
    case 5:
      var value = new github_com_solo$io_gloo_projects_gloo_api_v1_plugins_service_spec_pb.ServiceSpec;
      reader.readMessage(value,github_com_solo$io_gloo_projects_gloo_api_v1_plugins_service_spec_pb.ServiceSpec.deserializeBinaryFromReader);
      msg.setServiceSpec(value);
      break;
    case 6:
      var value = new github_com_solo$io_gloo_projects_gloo_api_v1_plugins_subset_spec_pb.SubsetSpec;
      reader.readMessage(value,github_com_solo$io_gloo_projects_gloo_api_v1_plugins_subset_spec_pb.SubsetSpec.deserializeBinaryFromReader);
      msg.setSubsetSpec(value);
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
proto.kubernetes.plugins.gloo.solo.io.UpstreamSpec.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.kubernetes.plugins.gloo.solo.io.UpstreamSpec.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.kubernetes.plugins.gloo.solo.io.UpstreamSpec} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.kubernetes.plugins.gloo.solo.io.UpstreamSpec.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getServiceName();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getServiceNamespace();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getServicePort();
  if (f !== 0) {
    writer.writeUint32(
      3,
      f
    );
  }
  f = message.getSelectorMap(true);
  if (f && f.getLength() > 0) {
    f.serializeBinary(4, writer, jspb.BinaryWriter.prototype.writeString, jspb.BinaryWriter.prototype.writeString);
  }
  f = message.getServiceSpec();
  if (f != null) {
    writer.writeMessage(
      5,
      f,
      github_com_solo$io_gloo_projects_gloo_api_v1_plugins_service_spec_pb.ServiceSpec.serializeBinaryToWriter
    );
  }
  f = message.getSubsetSpec();
  if (f != null) {
    writer.writeMessage(
      6,
      f,
      github_com_solo$io_gloo_projects_gloo_api_v1_plugins_subset_spec_pb.SubsetSpec.serializeBinaryToWriter
    );
  }
};


/**
 * optional string service_name = 1;
 * @return {string}
 */
proto.kubernetes.plugins.gloo.solo.io.UpstreamSpec.prototype.getServiceName = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/** @param {string} value */
proto.kubernetes.plugins.gloo.solo.io.UpstreamSpec.prototype.setServiceName = function(value) {
  jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string service_namespace = 2;
 * @return {string}
 */
proto.kubernetes.plugins.gloo.solo.io.UpstreamSpec.prototype.getServiceNamespace = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/** @param {string} value */
proto.kubernetes.plugins.gloo.solo.io.UpstreamSpec.prototype.setServiceNamespace = function(value) {
  jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional uint32 service_port = 3;
 * @return {number}
 */
proto.kubernetes.plugins.gloo.solo.io.UpstreamSpec.prototype.getServicePort = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 3, 0));
};


/** @param {number} value */
proto.kubernetes.plugins.gloo.solo.io.UpstreamSpec.prototype.setServicePort = function(value) {
  jspb.Message.setProto3IntField(this, 3, value);
};


/**
 * map<string, string> selector = 4;
 * @param {boolean=} opt_noLazyCreate Do not create the map if
 * empty, instead returning `undefined`
 * @return {!jspb.Map<string,string>}
 */
proto.kubernetes.plugins.gloo.solo.io.UpstreamSpec.prototype.getSelectorMap = function(opt_noLazyCreate) {
  return /** @type {!jspb.Map<string,string>} */ (
      jspb.Message.getMapField(this, 4, opt_noLazyCreate,
      null));
};


proto.kubernetes.plugins.gloo.solo.io.UpstreamSpec.prototype.clearSelectorMap = function() {
  this.getSelectorMap().clear();
};


/**
 * optional plugins.gloo.solo.io.ServiceSpec service_spec = 5;
 * @return {?proto.plugins.gloo.solo.io.ServiceSpec}
 */
proto.kubernetes.plugins.gloo.solo.io.UpstreamSpec.prototype.getServiceSpec = function() {
  return /** @type{?proto.plugins.gloo.solo.io.ServiceSpec} */ (
    jspb.Message.getWrapperField(this, github_com_solo$io_gloo_projects_gloo_api_v1_plugins_service_spec_pb.ServiceSpec, 5));
};


/** @param {?proto.plugins.gloo.solo.io.ServiceSpec|undefined} value */
proto.kubernetes.plugins.gloo.solo.io.UpstreamSpec.prototype.setServiceSpec = function(value) {
  jspb.Message.setWrapperField(this, 5, value);
};


proto.kubernetes.plugins.gloo.solo.io.UpstreamSpec.prototype.clearServiceSpec = function() {
  this.setServiceSpec(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.kubernetes.plugins.gloo.solo.io.UpstreamSpec.prototype.hasServiceSpec = function() {
  return jspb.Message.getField(this, 5) != null;
};


/**
 * optional plugins.gloo.solo.io.SubsetSpec subset_spec = 6;
 * @return {?proto.plugins.gloo.solo.io.SubsetSpec}
 */
proto.kubernetes.plugins.gloo.solo.io.UpstreamSpec.prototype.getSubsetSpec = function() {
  return /** @type{?proto.plugins.gloo.solo.io.SubsetSpec} */ (
    jspb.Message.getWrapperField(this, github_com_solo$io_gloo_projects_gloo_api_v1_plugins_subset_spec_pb.SubsetSpec, 6));
};


/** @param {?proto.plugins.gloo.solo.io.SubsetSpec|undefined} value */
proto.kubernetes.plugins.gloo.solo.io.UpstreamSpec.prototype.setSubsetSpec = function(value) {
  jspb.Message.setWrapperField(this, 6, value);
};


proto.kubernetes.plugins.gloo.solo.io.UpstreamSpec.prototype.clearSubsetSpec = function() {
  this.setSubsetSpec(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.kubernetes.plugins.gloo.solo.io.UpstreamSpec.prototype.hasSubsetSpec = function() {
  return jspb.Message.getField(this, 6) != null;
};


goog.object.extend(exports, proto.kubernetes.plugins.gloo.solo.io);
