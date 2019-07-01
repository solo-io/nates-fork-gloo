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
var github_com_solo$io_solo$kit_api_v1_ref_pb = require('../../../../../../../../../github.com/solo-io/solo-kit/api/v1/ref_pb.js');
goog.exportSymbol('proto.azure.plugins.gloo.solo.io.DestinationSpec', null, global);
goog.exportSymbol('proto.azure.plugins.gloo.solo.io.UpstreamSpec', null, global);
goog.exportSymbol('proto.azure.plugins.gloo.solo.io.UpstreamSpec.FunctionSpec', null, global);
goog.exportSymbol('proto.azure.plugins.gloo.solo.io.UpstreamSpec.FunctionSpec.AuthLevel', null, global);

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
proto.azure.plugins.gloo.solo.io.UpstreamSpec = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.azure.plugins.gloo.solo.io.UpstreamSpec.repeatedFields_, null);
};
goog.inherits(proto.azure.plugins.gloo.solo.io.UpstreamSpec, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.azure.plugins.gloo.solo.io.UpstreamSpec.displayName = 'proto.azure.plugins.gloo.solo.io.UpstreamSpec';
}
/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.azure.plugins.gloo.solo.io.UpstreamSpec.repeatedFields_ = [3];



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
proto.azure.plugins.gloo.solo.io.UpstreamSpec.prototype.toObject = function(opt_includeInstance) {
  return proto.azure.plugins.gloo.solo.io.UpstreamSpec.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.azure.plugins.gloo.solo.io.UpstreamSpec} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.azure.plugins.gloo.solo.io.UpstreamSpec.toObject = function(includeInstance, msg) {
  var f, obj = {
    functionAppName: jspb.Message.getFieldWithDefault(msg, 1, ""),
    secretRef: (f = msg.getSecretRef()) && github_com_solo$io_solo$kit_api_v1_ref_pb.ResourceRef.toObject(includeInstance, f),
    functionsList: jspb.Message.toObjectList(msg.getFunctionsList(),
    proto.azure.plugins.gloo.solo.io.UpstreamSpec.FunctionSpec.toObject, includeInstance)
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
 * @return {!proto.azure.plugins.gloo.solo.io.UpstreamSpec}
 */
proto.azure.plugins.gloo.solo.io.UpstreamSpec.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.azure.plugins.gloo.solo.io.UpstreamSpec;
  return proto.azure.plugins.gloo.solo.io.UpstreamSpec.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.azure.plugins.gloo.solo.io.UpstreamSpec} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.azure.plugins.gloo.solo.io.UpstreamSpec}
 */
proto.azure.plugins.gloo.solo.io.UpstreamSpec.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setFunctionAppName(value);
      break;
    case 2:
      var value = new github_com_solo$io_solo$kit_api_v1_ref_pb.ResourceRef;
      reader.readMessage(value,github_com_solo$io_solo$kit_api_v1_ref_pb.ResourceRef.deserializeBinaryFromReader);
      msg.setSecretRef(value);
      break;
    case 3:
      var value = new proto.azure.plugins.gloo.solo.io.UpstreamSpec.FunctionSpec;
      reader.readMessage(value,proto.azure.plugins.gloo.solo.io.UpstreamSpec.FunctionSpec.deserializeBinaryFromReader);
      msg.addFunctions(value);
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
proto.azure.plugins.gloo.solo.io.UpstreamSpec.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.azure.plugins.gloo.solo.io.UpstreamSpec.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.azure.plugins.gloo.solo.io.UpstreamSpec} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.azure.plugins.gloo.solo.io.UpstreamSpec.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getFunctionAppName();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getSecretRef();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      github_com_solo$io_solo$kit_api_v1_ref_pb.ResourceRef.serializeBinaryToWriter
    );
  }
  f = message.getFunctionsList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      3,
      f,
      proto.azure.plugins.gloo.solo.io.UpstreamSpec.FunctionSpec.serializeBinaryToWriter
    );
  }
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
proto.azure.plugins.gloo.solo.io.UpstreamSpec.FunctionSpec = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.azure.plugins.gloo.solo.io.UpstreamSpec.FunctionSpec, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.azure.plugins.gloo.solo.io.UpstreamSpec.FunctionSpec.displayName = 'proto.azure.plugins.gloo.solo.io.UpstreamSpec.FunctionSpec';
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
proto.azure.plugins.gloo.solo.io.UpstreamSpec.FunctionSpec.prototype.toObject = function(opt_includeInstance) {
  return proto.azure.plugins.gloo.solo.io.UpstreamSpec.FunctionSpec.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.azure.plugins.gloo.solo.io.UpstreamSpec.FunctionSpec} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.azure.plugins.gloo.solo.io.UpstreamSpec.FunctionSpec.toObject = function(includeInstance, msg) {
  var f, obj = {
    functionName: jspb.Message.getFieldWithDefault(msg, 1, ""),
    authLevel: jspb.Message.getFieldWithDefault(msg, 2, 0)
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
 * @return {!proto.azure.plugins.gloo.solo.io.UpstreamSpec.FunctionSpec}
 */
proto.azure.plugins.gloo.solo.io.UpstreamSpec.FunctionSpec.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.azure.plugins.gloo.solo.io.UpstreamSpec.FunctionSpec;
  return proto.azure.plugins.gloo.solo.io.UpstreamSpec.FunctionSpec.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.azure.plugins.gloo.solo.io.UpstreamSpec.FunctionSpec} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.azure.plugins.gloo.solo.io.UpstreamSpec.FunctionSpec}
 */
proto.azure.plugins.gloo.solo.io.UpstreamSpec.FunctionSpec.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setFunctionName(value);
      break;
    case 2:
      var value = /** @type {!proto.azure.plugins.gloo.solo.io.UpstreamSpec.FunctionSpec.AuthLevel} */ (reader.readEnum());
      msg.setAuthLevel(value);
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
proto.azure.plugins.gloo.solo.io.UpstreamSpec.FunctionSpec.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.azure.plugins.gloo.solo.io.UpstreamSpec.FunctionSpec.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.azure.plugins.gloo.solo.io.UpstreamSpec.FunctionSpec} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.azure.plugins.gloo.solo.io.UpstreamSpec.FunctionSpec.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getFunctionName();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getAuthLevel();
  if (f !== 0.0) {
    writer.writeEnum(
      2,
      f
    );
  }
};


/**
 * @enum {number}
 */
proto.azure.plugins.gloo.solo.io.UpstreamSpec.FunctionSpec.AuthLevel = {
  ANONYMOUS: 0,
  FUNCTION: 1,
  ADMIN: 2
};

/**
 * optional string function_name = 1;
 * @return {string}
 */
proto.azure.plugins.gloo.solo.io.UpstreamSpec.FunctionSpec.prototype.getFunctionName = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/** @param {string} value */
proto.azure.plugins.gloo.solo.io.UpstreamSpec.FunctionSpec.prototype.setFunctionName = function(value) {
  jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional AuthLevel auth_level = 2;
 * @return {!proto.azure.plugins.gloo.solo.io.UpstreamSpec.FunctionSpec.AuthLevel}
 */
proto.azure.plugins.gloo.solo.io.UpstreamSpec.FunctionSpec.prototype.getAuthLevel = function() {
  return /** @type {!proto.azure.plugins.gloo.solo.io.UpstreamSpec.FunctionSpec.AuthLevel} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/** @param {!proto.azure.plugins.gloo.solo.io.UpstreamSpec.FunctionSpec.AuthLevel} value */
proto.azure.plugins.gloo.solo.io.UpstreamSpec.FunctionSpec.prototype.setAuthLevel = function(value) {
  jspb.Message.setProto3EnumField(this, 2, value);
};


/**
 * optional string function_app_name = 1;
 * @return {string}
 */
proto.azure.plugins.gloo.solo.io.UpstreamSpec.prototype.getFunctionAppName = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/** @param {string} value */
proto.azure.plugins.gloo.solo.io.UpstreamSpec.prototype.setFunctionAppName = function(value) {
  jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional core.solo.io.ResourceRef secret_ref = 2;
 * @return {?proto.core.solo.io.ResourceRef}
 */
proto.azure.plugins.gloo.solo.io.UpstreamSpec.prototype.getSecretRef = function() {
  return /** @type{?proto.core.solo.io.ResourceRef} */ (
    jspb.Message.getWrapperField(this, github_com_solo$io_solo$kit_api_v1_ref_pb.ResourceRef, 2));
};


/** @param {?proto.core.solo.io.ResourceRef|undefined} value */
proto.azure.plugins.gloo.solo.io.UpstreamSpec.prototype.setSecretRef = function(value) {
  jspb.Message.setWrapperField(this, 2, value);
};


proto.azure.plugins.gloo.solo.io.UpstreamSpec.prototype.clearSecretRef = function() {
  this.setSecretRef(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.azure.plugins.gloo.solo.io.UpstreamSpec.prototype.hasSecretRef = function() {
  return jspb.Message.getField(this, 2) != null;
};


/**
 * repeated FunctionSpec functions = 3;
 * @return {!Array<!proto.azure.plugins.gloo.solo.io.UpstreamSpec.FunctionSpec>}
 */
proto.azure.plugins.gloo.solo.io.UpstreamSpec.prototype.getFunctionsList = function() {
  return /** @type{!Array<!proto.azure.plugins.gloo.solo.io.UpstreamSpec.FunctionSpec>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.azure.plugins.gloo.solo.io.UpstreamSpec.FunctionSpec, 3));
};


/** @param {!Array<!proto.azure.plugins.gloo.solo.io.UpstreamSpec.FunctionSpec>} value */
proto.azure.plugins.gloo.solo.io.UpstreamSpec.prototype.setFunctionsList = function(value) {
  jspb.Message.setRepeatedWrapperField(this, 3, value);
};


/**
 * @param {!proto.azure.plugins.gloo.solo.io.UpstreamSpec.FunctionSpec=} opt_value
 * @param {number=} opt_index
 * @return {!proto.azure.plugins.gloo.solo.io.UpstreamSpec.FunctionSpec}
 */
proto.azure.plugins.gloo.solo.io.UpstreamSpec.prototype.addFunctions = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 3, opt_value, proto.azure.plugins.gloo.solo.io.UpstreamSpec.FunctionSpec, opt_index);
};


proto.azure.plugins.gloo.solo.io.UpstreamSpec.prototype.clearFunctionsList = function() {
  this.setFunctionsList([]);
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
proto.azure.plugins.gloo.solo.io.DestinationSpec = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.azure.plugins.gloo.solo.io.DestinationSpec, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.azure.plugins.gloo.solo.io.DestinationSpec.displayName = 'proto.azure.plugins.gloo.solo.io.DestinationSpec';
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
proto.azure.plugins.gloo.solo.io.DestinationSpec.prototype.toObject = function(opt_includeInstance) {
  return proto.azure.plugins.gloo.solo.io.DestinationSpec.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.azure.plugins.gloo.solo.io.DestinationSpec} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.azure.plugins.gloo.solo.io.DestinationSpec.toObject = function(includeInstance, msg) {
  var f, obj = {
    functionName: jspb.Message.getFieldWithDefault(msg, 1, "")
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
 * @return {!proto.azure.plugins.gloo.solo.io.DestinationSpec}
 */
proto.azure.plugins.gloo.solo.io.DestinationSpec.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.azure.plugins.gloo.solo.io.DestinationSpec;
  return proto.azure.plugins.gloo.solo.io.DestinationSpec.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.azure.plugins.gloo.solo.io.DestinationSpec} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.azure.plugins.gloo.solo.io.DestinationSpec}
 */
proto.azure.plugins.gloo.solo.io.DestinationSpec.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setFunctionName(value);
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
proto.azure.plugins.gloo.solo.io.DestinationSpec.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.azure.plugins.gloo.solo.io.DestinationSpec.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.azure.plugins.gloo.solo.io.DestinationSpec} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.azure.plugins.gloo.solo.io.DestinationSpec.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getFunctionName();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
};


/**
 * optional string function_name = 1;
 * @return {string}
 */
proto.azure.plugins.gloo.solo.io.DestinationSpec.prototype.getFunctionName = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/** @param {string} value */
proto.azure.plugins.gloo.solo.io.DestinationSpec.prototype.setFunctionName = function(value) {
  jspb.Message.setProto3StringField(this, 1, value);
};


goog.object.extend(exports, proto.azure.plugins.gloo.solo.io);
