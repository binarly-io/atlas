// source: data_chunk_properties.proto
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

var digest_pb = require('./digest_pb.js');
goog.object.extend(proto, digest_pb);
goog.exportSymbol('proto.atlas.DataChunkProperties', null, global);
goog.exportSymbol('proto.atlas.DataChunkProperties.DigestOptionalCase', null, global);
goog.exportSymbol('proto.atlas.DataChunkProperties.LastOptionalCase', null, global);
goog.exportSymbol('proto.atlas.DataChunkProperties.LenOptionalCase', null, global);
goog.exportSymbol('proto.atlas.DataChunkProperties.OffsetOptionalCase', null, global);
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
proto.atlas.DataChunkProperties = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, proto.atlas.DataChunkProperties.oneofGroups_);
};
goog.inherits(proto.atlas.DataChunkProperties, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.atlas.DataChunkProperties.displayName = 'proto.atlas.DataChunkProperties';
}

/**
 * Oneof group definitions for this message. Each group defines the field
 * numbers belonging to that group. When of these fields' value is set, all
 * other fields in the group are cleared. During deserialization, if multiple
 * fields are encountered for a group, only the last value seen will be kept.
 * @private {!Array<!Array<number>>}
 * @const
 */
proto.atlas.DataChunkProperties.oneofGroups_ = [[100],[200],[300],[400]];

/**
 * @enum {number}
 */
proto.atlas.DataChunkProperties.DigestOptionalCase = {
  DIGEST_OPTIONAL_NOT_SET: 0,
  DIGEST: 100
};

/**
 * @return {proto.atlas.DataChunkProperties.DigestOptionalCase}
 */
proto.atlas.DataChunkProperties.prototype.getDigestOptionalCase = function() {
  return /** @type {proto.atlas.DataChunkProperties.DigestOptionalCase} */(jspb.Message.computeOneofCase(this, proto.atlas.DataChunkProperties.oneofGroups_[0]));
};

/**
 * @enum {number}
 */
proto.atlas.DataChunkProperties.LenOptionalCase = {
  LEN_OPTIONAL_NOT_SET: 0,
  LEN: 200
};

/**
 * @return {proto.atlas.DataChunkProperties.LenOptionalCase}
 */
proto.atlas.DataChunkProperties.prototype.getLenOptionalCase = function() {
  return /** @type {proto.atlas.DataChunkProperties.LenOptionalCase} */(jspb.Message.computeOneofCase(this, proto.atlas.DataChunkProperties.oneofGroups_[1]));
};

/**
 * @enum {number}
 */
proto.atlas.DataChunkProperties.OffsetOptionalCase = {
  OFFSET_OPTIONAL_NOT_SET: 0,
  OFFSET: 300
};

/**
 * @return {proto.atlas.DataChunkProperties.OffsetOptionalCase}
 */
proto.atlas.DataChunkProperties.prototype.getOffsetOptionalCase = function() {
  return /** @type {proto.atlas.DataChunkProperties.OffsetOptionalCase} */(jspb.Message.computeOneofCase(this, proto.atlas.DataChunkProperties.oneofGroups_[2]));
};

/**
 * @enum {number}
 */
proto.atlas.DataChunkProperties.LastOptionalCase = {
  LAST_OPTIONAL_NOT_SET: 0,
  LAST: 400
};

/**
 * @return {proto.atlas.DataChunkProperties.LastOptionalCase}
 */
proto.atlas.DataChunkProperties.prototype.getLastOptionalCase = function() {
  return /** @type {proto.atlas.DataChunkProperties.LastOptionalCase} */(jspb.Message.computeOneofCase(this, proto.atlas.DataChunkProperties.oneofGroups_[3]));
};



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.atlas.DataChunkProperties.prototype.toObject = function(opt_includeInstance) {
  return proto.atlas.DataChunkProperties.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.atlas.DataChunkProperties} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.atlas.DataChunkProperties.toObject = function(includeInstance, msg) {
  var f, obj = {
    digest: (f = msg.getDigest()) && digest_pb.Digest.toObject(includeInstance, f),
    len: jspb.Message.getFieldWithDefault(msg, 200, 0),
    offset: jspb.Message.getFieldWithDefault(msg, 300, 0),
    last: jspb.Message.getBooleanFieldWithDefault(msg, 400, false)
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
 * @return {!proto.atlas.DataChunkProperties}
 */
proto.atlas.DataChunkProperties.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.atlas.DataChunkProperties;
  return proto.atlas.DataChunkProperties.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.atlas.DataChunkProperties} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.atlas.DataChunkProperties}
 */
proto.atlas.DataChunkProperties.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 100:
      var value = new digest_pb.Digest;
      reader.readMessage(value,digest_pb.Digest.deserializeBinaryFromReader);
      msg.setDigest(value);
      break;
    case 200:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setLen(value);
      break;
    case 300:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setOffset(value);
      break;
    case 400:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setLast(value);
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
proto.atlas.DataChunkProperties.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.atlas.DataChunkProperties.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.atlas.DataChunkProperties} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.atlas.DataChunkProperties.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getDigest();
  if (f != null) {
    writer.writeMessage(
      100,
      f,
      digest_pb.Digest.serializeBinaryToWriter
    );
  }
  f = /** @type {number} */ (jspb.Message.getField(message, 200));
  if (f != null) {
    writer.writeInt64(
      200,
      f
    );
  }
  f = /** @type {number} */ (jspb.Message.getField(message, 300));
  if (f != null) {
    writer.writeInt64(
      300,
      f
    );
  }
  f = /** @type {boolean} */ (jspb.Message.getField(message, 400));
  if (f != null) {
    writer.writeBool(
      400,
      f
    );
  }
};


/**
 * optional Digest digest = 100;
 * @return {?proto.atlas.Digest}
 */
proto.atlas.DataChunkProperties.prototype.getDigest = function() {
  return /** @type{?proto.atlas.Digest} */ (
    jspb.Message.getWrapperField(this, digest_pb.Digest, 100));
};


/**
 * @param {?proto.atlas.Digest|undefined} value
 * @return {!proto.atlas.DataChunkProperties} returns this
*/
proto.atlas.DataChunkProperties.prototype.setDigest = function(value) {
  return jspb.Message.setOneofWrapperField(this, 100, proto.atlas.DataChunkProperties.oneofGroups_[0], value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.atlas.DataChunkProperties} returns this
 */
proto.atlas.DataChunkProperties.prototype.clearDigest = function() {
  return this.setDigest(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.atlas.DataChunkProperties.prototype.hasDigest = function() {
  return jspb.Message.getField(this, 100) != null;
};


/**
 * optional int64 len = 200;
 * @return {number}
 */
proto.atlas.DataChunkProperties.prototype.getLen = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 200, 0));
};


/**
 * @param {number} value
 * @return {!proto.atlas.DataChunkProperties} returns this
 */
proto.atlas.DataChunkProperties.prototype.setLen = function(value) {
  return jspb.Message.setOneofField(this, 200, proto.atlas.DataChunkProperties.oneofGroups_[1], value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.atlas.DataChunkProperties} returns this
 */
proto.atlas.DataChunkProperties.prototype.clearLen = function() {
  return jspb.Message.setOneofField(this, 200, proto.atlas.DataChunkProperties.oneofGroups_[1], undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.atlas.DataChunkProperties.prototype.hasLen = function() {
  return jspb.Message.getField(this, 200) != null;
};


/**
 * optional int64 offset = 300;
 * @return {number}
 */
proto.atlas.DataChunkProperties.prototype.getOffset = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 300, 0));
};


/**
 * @param {number} value
 * @return {!proto.atlas.DataChunkProperties} returns this
 */
proto.atlas.DataChunkProperties.prototype.setOffset = function(value) {
  return jspb.Message.setOneofField(this, 300, proto.atlas.DataChunkProperties.oneofGroups_[2], value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.atlas.DataChunkProperties} returns this
 */
proto.atlas.DataChunkProperties.prototype.clearOffset = function() {
  return jspb.Message.setOneofField(this, 300, proto.atlas.DataChunkProperties.oneofGroups_[2], undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.atlas.DataChunkProperties.prototype.hasOffset = function() {
  return jspb.Message.getField(this, 300) != null;
};


/**
 * optional bool last = 400;
 * @return {boolean}
 */
proto.atlas.DataChunkProperties.prototype.getLast = function() {
  return /** @type {boolean} */ (jspb.Message.getBooleanFieldWithDefault(this, 400, false));
};


/**
 * @param {boolean} value
 * @return {!proto.atlas.DataChunkProperties} returns this
 */
proto.atlas.DataChunkProperties.prototype.setLast = function(value) {
  return jspb.Message.setOneofField(this, 400, proto.atlas.DataChunkProperties.oneofGroups_[3], value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.atlas.DataChunkProperties} returns this
 */
proto.atlas.DataChunkProperties.prototype.clearLast = function() {
  return jspb.Message.setOneofField(this, 400, proto.atlas.DataChunkProperties.oneofGroups_[3], undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.atlas.DataChunkProperties.prototype.hasLast = function() {
  return jspb.Message.getField(this, 400) != null;
};


goog.object.extend(exports, proto.atlas);
