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

var google_protobuf_wrappers_pb = require('google-protobuf/google/protobuf/wrappers_pb.js');
var udpa_annotations_status_pb = require('../../../../udpa/annotations/status_pb.js');
var udpa_annotations_versioning_pb = require('../../../../udpa/annotations/versioning_pb.js');
var validate_validate_pb = require('../../../../validate/validate_pb.js');
var gogoproto_gogo_pb = require('../../../../gogoproto/gogo_pb.js');
goog.exportSymbol('proto.envoy.type.matcher.v3.RegexMatchAndSubstitute', null, global);
goog.exportSymbol('proto.envoy.type.matcher.v3.RegexMatcher', null, global);
goog.exportSymbol('proto.envoy.type.matcher.v3.RegexMatcher.GoogleRE2', null, global);

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
proto.envoy.type.matcher.v3.RegexMatcher = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, proto.envoy.type.matcher.v3.RegexMatcher.oneofGroups_);
};
goog.inherits(proto.envoy.type.matcher.v3.RegexMatcher, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.envoy.type.matcher.v3.RegexMatcher.displayName = 'proto.envoy.type.matcher.v3.RegexMatcher';
}
/**
 * Oneof group definitions for this message. Each group defines the field
 * numbers belonging to that group. When of these fields' value is set, all
 * other fields in the group are cleared. During deserialization, if multiple
 * fields are encountered for a group, only the last value seen will be kept.
 * @private {!Array<!Array<number>>}
 * @const
 */
proto.envoy.type.matcher.v3.RegexMatcher.oneofGroups_ = [[1]];

/**
 * @enum {number}
 */
proto.envoy.type.matcher.v3.RegexMatcher.EngineTypeCase = {
  ENGINE_TYPE_NOT_SET: 0,
  GOOGLE_RE2: 1
};

/**
 * @return {proto.envoy.type.matcher.v3.RegexMatcher.EngineTypeCase}
 */
proto.envoy.type.matcher.v3.RegexMatcher.prototype.getEngineTypeCase = function() {
  return /** @type {proto.envoy.type.matcher.v3.RegexMatcher.EngineTypeCase} */(jspb.Message.computeOneofCase(this, proto.envoy.type.matcher.v3.RegexMatcher.oneofGroups_[0]));
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
proto.envoy.type.matcher.v3.RegexMatcher.prototype.toObject = function(opt_includeInstance) {
  return proto.envoy.type.matcher.v3.RegexMatcher.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.envoy.type.matcher.v3.RegexMatcher} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.envoy.type.matcher.v3.RegexMatcher.toObject = function(includeInstance, msg) {
  var f, obj = {
    googleRe2: (f = msg.getGoogleRe2()) && proto.envoy.type.matcher.v3.RegexMatcher.GoogleRE2.toObject(includeInstance, f),
    regex: jspb.Message.getFieldWithDefault(msg, 2, "")
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
 * @return {!proto.envoy.type.matcher.v3.RegexMatcher}
 */
proto.envoy.type.matcher.v3.RegexMatcher.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.envoy.type.matcher.v3.RegexMatcher;
  return proto.envoy.type.matcher.v3.RegexMatcher.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.envoy.type.matcher.v3.RegexMatcher} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.envoy.type.matcher.v3.RegexMatcher}
 */
proto.envoy.type.matcher.v3.RegexMatcher.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.envoy.type.matcher.v3.RegexMatcher.GoogleRE2;
      reader.readMessage(value,proto.envoy.type.matcher.v3.RegexMatcher.GoogleRE2.deserializeBinaryFromReader);
      msg.setGoogleRe2(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setRegex(value);
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
proto.envoy.type.matcher.v3.RegexMatcher.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.envoy.type.matcher.v3.RegexMatcher.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.envoy.type.matcher.v3.RegexMatcher} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.envoy.type.matcher.v3.RegexMatcher.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getGoogleRe2();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      proto.envoy.type.matcher.v3.RegexMatcher.GoogleRE2.serializeBinaryToWriter
    );
  }
  f = message.getRegex();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
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
proto.envoy.type.matcher.v3.RegexMatcher.GoogleRE2 = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.envoy.type.matcher.v3.RegexMatcher.GoogleRE2, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.envoy.type.matcher.v3.RegexMatcher.GoogleRE2.displayName = 'proto.envoy.type.matcher.v3.RegexMatcher.GoogleRE2';
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
proto.envoy.type.matcher.v3.RegexMatcher.GoogleRE2.prototype.toObject = function(opt_includeInstance) {
  return proto.envoy.type.matcher.v3.RegexMatcher.GoogleRE2.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.envoy.type.matcher.v3.RegexMatcher.GoogleRE2} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.envoy.type.matcher.v3.RegexMatcher.GoogleRE2.toObject = function(includeInstance, msg) {
  var f, obj = {
    maxProgramSize: (f = msg.getMaxProgramSize()) && google_protobuf_wrappers_pb.UInt32Value.toObject(includeInstance, f)
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
 * @return {!proto.envoy.type.matcher.v3.RegexMatcher.GoogleRE2}
 */
proto.envoy.type.matcher.v3.RegexMatcher.GoogleRE2.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.envoy.type.matcher.v3.RegexMatcher.GoogleRE2;
  return proto.envoy.type.matcher.v3.RegexMatcher.GoogleRE2.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.envoy.type.matcher.v3.RegexMatcher.GoogleRE2} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.envoy.type.matcher.v3.RegexMatcher.GoogleRE2}
 */
proto.envoy.type.matcher.v3.RegexMatcher.GoogleRE2.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new google_protobuf_wrappers_pb.UInt32Value;
      reader.readMessage(value,google_protobuf_wrappers_pb.UInt32Value.deserializeBinaryFromReader);
      msg.setMaxProgramSize(value);
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
proto.envoy.type.matcher.v3.RegexMatcher.GoogleRE2.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.envoy.type.matcher.v3.RegexMatcher.GoogleRE2.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.envoy.type.matcher.v3.RegexMatcher.GoogleRE2} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.envoy.type.matcher.v3.RegexMatcher.GoogleRE2.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getMaxProgramSize();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      google_protobuf_wrappers_pb.UInt32Value.serializeBinaryToWriter
    );
  }
};


/**
 * optional google.protobuf.UInt32Value max_program_size = 1;
 * @return {?proto.google.protobuf.UInt32Value}
 */
proto.envoy.type.matcher.v3.RegexMatcher.GoogleRE2.prototype.getMaxProgramSize = function() {
  return /** @type{?proto.google.protobuf.UInt32Value} */ (
    jspb.Message.getWrapperField(this, google_protobuf_wrappers_pb.UInt32Value, 1));
};


/** @param {?proto.google.protobuf.UInt32Value|undefined} value */
proto.envoy.type.matcher.v3.RegexMatcher.GoogleRE2.prototype.setMaxProgramSize = function(value) {
  jspb.Message.setWrapperField(this, 1, value);
};


proto.envoy.type.matcher.v3.RegexMatcher.GoogleRE2.prototype.clearMaxProgramSize = function() {
  this.setMaxProgramSize(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.envoy.type.matcher.v3.RegexMatcher.GoogleRE2.prototype.hasMaxProgramSize = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional GoogleRE2 google_re2 = 1;
 * @return {?proto.envoy.type.matcher.v3.RegexMatcher.GoogleRE2}
 */
proto.envoy.type.matcher.v3.RegexMatcher.prototype.getGoogleRe2 = function() {
  return /** @type{?proto.envoy.type.matcher.v3.RegexMatcher.GoogleRE2} */ (
    jspb.Message.getWrapperField(this, proto.envoy.type.matcher.v3.RegexMatcher.GoogleRE2, 1));
};


/** @param {?proto.envoy.type.matcher.v3.RegexMatcher.GoogleRE2|undefined} value */
proto.envoy.type.matcher.v3.RegexMatcher.prototype.setGoogleRe2 = function(value) {
  jspb.Message.setOneofWrapperField(this, 1, proto.envoy.type.matcher.v3.RegexMatcher.oneofGroups_[0], value);
};


proto.envoy.type.matcher.v3.RegexMatcher.prototype.clearGoogleRe2 = function() {
  this.setGoogleRe2(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.envoy.type.matcher.v3.RegexMatcher.prototype.hasGoogleRe2 = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional string regex = 2;
 * @return {string}
 */
proto.envoy.type.matcher.v3.RegexMatcher.prototype.getRegex = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/** @param {string} value */
proto.envoy.type.matcher.v3.RegexMatcher.prototype.setRegex = function(value) {
  jspb.Message.setProto3StringField(this, 2, value);
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
proto.envoy.type.matcher.v3.RegexMatchAndSubstitute = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.envoy.type.matcher.v3.RegexMatchAndSubstitute, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.envoy.type.matcher.v3.RegexMatchAndSubstitute.displayName = 'proto.envoy.type.matcher.v3.RegexMatchAndSubstitute';
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
proto.envoy.type.matcher.v3.RegexMatchAndSubstitute.prototype.toObject = function(opt_includeInstance) {
  return proto.envoy.type.matcher.v3.RegexMatchAndSubstitute.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.envoy.type.matcher.v3.RegexMatchAndSubstitute} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.envoy.type.matcher.v3.RegexMatchAndSubstitute.toObject = function(includeInstance, msg) {
  var f, obj = {
    pattern: (f = msg.getPattern()) && proto.envoy.type.matcher.v3.RegexMatcher.toObject(includeInstance, f),
    substitution: jspb.Message.getFieldWithDefault(msg, 2, "")
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
 * @return {!proto.envoy.type.matcher.v3.RegexMatchAndSubstitute}
 */
proto.envoy.type.matcher.v3.RegexMatchAndSubstitute.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.envoy.type.matcher.v3.RegexMatchAndSubstitute;
  return proto.envoy.type.matcher.v3.RegexMatchAndSubstitute.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.envoy.type.matcher.v3.RegexMatchAndSubstitute} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.envoy.type.matcher.v3.RegexMatchAndSubstitute}
 */
proto.envoy.type.matcher.v3.RegexMatchAndSubstitute.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.envoy.type.matcher.v3.RegexMatcher;
      reader.readMessage(value,proto.envoy.type.matcher.v3.RegexMatcher.deserializeBinaryFromReader);
      msg.setPattern(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setSubstitution(value);
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
proto.envoy.type.matcher.v3.RegexMatchAndSubstitute.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.envoy.type.matcher.v3.RegexMatchAndSubstitute.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.envoy.type.matcher.v3.RegexMatchAndSubstitute} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.envoy.type.matcher.v3.RegexMatchAndSubstitute.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getPattern();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      proto.envoy.type.matcher.v3.RegexMatcher.serializeBinaryToWriter
    );
  }
  f = message.getSubstitution();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
};


/**
 * optional RegexMatcher pattern = 1;
 * @return {?proto.envoy.type.matcher.v3.RegexMatcher}
 */
proto.envoy.type.matcher.v3.RegexMatchAndSubstitute.prototype.getPattern = function() {
  return /** @type{?proto.envoy.type.matcher.v3.RegexMatcher} */ (
    jspb.Message.getWrapperField(this, proto.envoy.type.matcher.v3.RegexMatcher, 1));
};


/** @param {?proto.envoy.type.matcher.v3.RegexMatcher|undefined} value */
proto.envoy.type.matcher.v3.RegexMatchAndSubstitute.prototype.setPattern = function(value) {
  jspb.Message.setWrapperField(this, 1, value);
};


proto.envoy.type.matcher.v3.RegexMatchAndSubstitute.prototype.clearPattern = function() {
  this.setPattern(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.envoy.type.matcher.v3.RegexMatchAndSubstitute.prototype.hasPattern = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional string substitution = 2;
 * @return {string}
 */
proto.envoy.type.matcher.v3.RegexMatchAndSubstitute.prototype.getSubstitution = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/** @param {string} value */
proto.envoy.type.matcher.v3.RegexMatchAndSubstitute.prototype.setSubstitution = function(value) {
  jspb.Message.setProto3StringField(this, 2, value);
};


goog.object.extend(exports, proto.envoy.type.matcher.v3);
