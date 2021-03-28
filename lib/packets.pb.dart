///
//  Generated code. Do not modify.
//  source: packets.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

import 'google/protobuf/timestamp.pb.dart' as $2;
import 'user.pb.dart' as $3;
import 'command.pb.dart' as $4;

import 'packets.pbenum.dart';

export 'packets.pbenum.dart';

class MessagePacket extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'MessagePacket', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'packets'), createEmptyInstance: create)
    ..aOM<$2.Timestamp>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'timestamp', subBuilder: $2.Timestamp.create)
    ..aOS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'message')
    ..aOM<$3.User>(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'user', subBuilder: $3.User.create)
    ..aOB(4, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'private')
    ..hasRequiredFields = false
  ;

  MessagePacket._() : super();
  factory MessagePacket({
    $2.Timestamp? timestamp,
    $core.String? message,
    $3.User? user,
    $core.bool? private,
  }) {
    final _result = create();
    if (timestamp != null) {
      _result.timestamp = timestamp;
    }
    if (message != null) {
      _result.message = message;
    }
    if (user != null) {
      _result.user = user;
    }
    if (private != null) {
      _result.private = private;
    }
    return _result;
  }
  factory MessagePacket.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory MessagePacket.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  MessagePacket clone() => MessagePacket()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  MessagePacket copyWith(void Function(MessagePacket) updates) => super.copyWith((message) => updates(message as MessagePacket)) as MessagePacket; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static MessagePacket create() => MessagePacket._();
  MessagePacket createEmptyInstance() => create();
  static $pb.PbList<MessagePacket> createRepeated() => $pb.PbList<MessagePacket>();
  @$core.pragma('dart2js:noInline')
  static MessagePacket getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<MessagePacket>(create);
  static MessagePacket? _defaultInstance;

  @$pb.TagNumber(1)
  $2.Timestamp get timestamp => $_getN(0);
  @$pb.TagNumber(1)
  set timestamp($2.Timestamp v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasTimestamp() => $_has(0);
  @$pb.TagNumber(1)
  void clearTimestamp() => clearField(1);
  @$pb.TagNumber(1)
  $2.Timestamp ensureTimestamp() => $_ensure(0);

  @$pb.TagNumber(2)
  $core.String get message => $_getSZ(1);
  @$pb.TagNumber(2)
  set message($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasMessage() => $_has(1);
  @$pb.TagNumber(2)
  void clearMessage() => clearField(2);

  @$pb.TagNumber(3)
  $3.User get user => $_getN(2);
  @$pb.TagNumber(3)
  set user($3.User v) { setField(3, v); }
  @$pb.TagNumber(3)
  $core.bool hasUser() => $_has(2);
  @$pb.TagNumber(3)
  void clearUser() => clearField(3);
  @$pb.TagNumber(3)
  $3.User ensureUser() => $_ensure(2);

  @$pb.TagNumber(4)
  $core.bool get private => $_getBF(3);
  @$pb.TagNumber(4)
  set private($core.bool v) { $_setBool(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasPrivate() => $_has(3);
  @$pb.TagNumber(4)
  void clearPrivate() => clearField(4);
}

class UserPacket extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'UserPacket', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'packets'), createEmptyInstance: create)
    ..aOM<$2.Timestamp>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'timestamp', subBuilder: $2.Timestamp.create)
    ..aOM<$3.User>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'user', subBuilder: $3.User.create)
    ..e<UserEvent>(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'event', $pb.PbFieldType.OE, defaultOrMaker: UserEvent.JOINED, valueOf: UserEvent.valueOf, enumValues: UserEvent.values)
    ..hasRequiredFields = false
  ;

  UserPacket._() : super();
  factory UserPacket({
    $2.Timestamp? timestamp,
    $3.User? user,
    UserEvent? event,
  }) {
    final _result = create();
    if (timestamp != null) {
      _result.timestamp = timestamp;
    }
    if (user != null) {
      _result.user = user;
    }
    if (event != null) {
      _result.event = event;
    }
    return _result;
  }
  factory UserPacket.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory UserPacket.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  UserPacket clone() => UserPacket()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  UserPacket copyWith(void Function(UserPacket) updates) => super.copyWith((message) => updates(message as UserPacket)) as UserPacket; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static UserPacket create() => UserPacket._();
  UserPacket createEmptyInstance() => create();
  static $pb.PbList<UserPacket> createRepeated() => $pb.PbList<UserPacket>();
  @$core.pragma('dart2js:noInline')
  static UserPacket getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<UserPacket>(create);
  static UserPacket? _defaultInstance;

  @$pb.TagNumber(1)
  $2.Timestamp get timestamp => $_getN(0);
  @$pb.TagNumber(1)
  set timestamp($2.Timestamp v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasTimestamp() => $_has(0);
  @$pb.TagNumber(1)
  void clearTimestamp() => clearField(1);
  @$pb.TagNumber(1)
  $2.Timestamp ensureTimestamp() => $_ensure(0);

  @$pb.TagNumber(2)
  $3.User get user => $_getN(1);
  @$pb.TagNumber(2)
  set user($3.User v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasUser() => $_has(1);
  @$pb.TagNumber(2)
  void clearUser() => clearField(2);
  @$pb.TagNumber(2)
  $3.User ensureUser() => $_ensure(1);

  @$pb.TagNumber(3)
  UserEvent get event => $_getN(2);
  @$pb.TagNumber(3)
  set event(UserEvent v) { setField(3, v); }
  @$pb.TagNumber(3)
  $core.bool hasEvent() => $_has(2);
  @$pb.TagNumber(3)
  void clearEvent() => clearField(3);
}

class CommandPacket extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'CommandPacket', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'packets'), createEmptyInstance: create)
    ..aOM<$2.Timestamp>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'timestamp', subBuilder: $2.Timestamp.create)
    ..aOS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'command')
    ..pPS(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'args')
    ..aOM<$3.User>(4, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'user', subBuilder: $3.User.create)
    ..aOB(5, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'private')
    ..aOS(6, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'argString', protoName: 'argString')
    ..hasRequiredFields = false
  ;

  CommandPacket._() : super();
  factory CommandPacket({
    $2.Timestamp? timestamp,
    $core.String? command,
    $core.Iterable<$core.String>? args,
    $3.User? user,
    $core.bool? private,
    $core.String? argString,
  }) {
    final _result = create();
    if (timestamp != null) {
      _result.timestamp = timestamp;
    }
    if (command != null) {
      _result.command = command;
    }
    if (args != null) {
      _result.args.addAll(args);
    }
    if (user != null) {
      _result.user = user;
    }
    if (private != null) {
      _result.private = private;
    }
    if (argString != null) {
      _result.argString = argString;
    }
    return _result;
  }
  factory CommandPacket.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory CommandPacket.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  CommandPacket clone() => CommandPacket()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  CommandPacket copyWith(void Function(CommandPacket) updates) => super.copyWith((message) => updates(message as CommandPacket)) as CommandPacket; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static CommandPacket create() => CommandPacket._();
  CommandPacket createEmptyInstance() => create();
  static $pb.PbList<CommandPacket> createRepeated() => $pb.PbList<CommandPacket>();
  @$core.pragma('dart2js:noInline')
  static CommandPacket getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<CommandPacket>(create);
  static CommandPacket? _defaultInstance;

  @$pb.TagNumber(1)
  $2.Timestamp get timestamp => $_getN(0);
  @$pb.TagNumber(1)
  set timestamp($2.Timestamp v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasTimestamp() => $_has(0);
  @$pb.TagNumber(1)
  void clearTimestamp() => clearField(1);
  @$pb.TagNumber(1)
  $2.Timestamp ensureTimestamp() => $_ensure(0);

  @$pb.TagNumber(2)
  $core.String get command => $_getSZ(1);
  @$pb.TagNumber(2)
  set command($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasCommand() => $_has(1);
  @$pb.TagNumber(2)
  void clearCommand() => clearField(2);

  @$pb.TagNumber(3)
  $core.List<$core.String> get args => $_getList(2);

  @$pb.TagNumber(4)
  $3.User get user => $_getN(3);
  @$pb.TagNumber(4)
  set user($3.User v) { setField(4, v); }
  @$pb.TagNumber(4)
  $core.bool hasUser() => $_has(3);
  @$pb.TagNumber(4)
  void clearUser() => clearField(4);
  @$pb.TagNumber(4)
  $3.User ensureUser() => $_ensure(3);

  @$pb.TagNumber(5)
  $core.bool get private => $_getBF(4);
  @$pb.TagNumber(5)
  set private($core.bool v) { $_setBool(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasPrivate() => $_has(4);
  @$pb.TagNumber(5)
  void clearPrivate() => clearField(5);

  @$pb.TagNumber(6)
  $core.String get argString => $_getSZ(5);
  @$pb.TagNumber(6)
  set argString($core.String v) { $_setString(5, v); }
  @$pb.TagNumber(6)
  $core.bool hasArgString() => $_has(5);
  @$pb.TagNumber(6)
  void clearArgString() => clearField(6);
}

class BotPacket extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'BotPacket', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'packets'), createEmptyInstance: create)
    ..aOM<$2.Timestamp>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'timestamp', subBuilder: $2.Timestamp.create)
    ..aOS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'message')
    ..aOM<$3.User>(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'recipient', subBuilder: $3.User.create)
    ..aOB(4, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'private')
    ..hasRequiredFields = false
  ;

  BotPacket._() : super();
  factory BotPacket({
    $2.Timestamp? timestamp,
    $core.String? message,
    $3.User? recipient,
    $core.bool? private,
  }) {
    final _result = create();
    if (timestamp != null) {
      _result.timestamp = timestamp;
    }
    if (message != null) {
      _result.message = message;
    }
    if (recipient != null) {
      _result.recipient = recipient;
    }
    if (private != null) {
      _result.private = private;
    }
    return _result;
  }
  factory BotPacket.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory BotPacket.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  BotPacket clone() => BotPacket()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  BotPacket copyWith(void Function(BotPacket) updates) => super.copyWith((message) => updates(message as BotPacket)) as BotPacket; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static BotPacket create() => BotPacket._();
  BotPacket createEmptyInstance() => create();
  static $pb.PbList<BotPacket> createRepeated() => $pb.PbList<BotPacket>();
  @$core.pragma('dart2js:noInline')
  static BotPacket getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<BotPacket>(create);
  static BotPacket? _defaultInstance;

  @$pb.TagNumber(1)
  $2.Timestamp get timestamp => $_getN(0);
  @$pb.TagNumber(1)
  set timestamp($2.Timestamp v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasTimestamp() => $_has(0);
  @$pb.TagNumber(1)
  void clearTimestamp() => clearField(1);
  @$pb.TagNumber(1)
  $2.Timestamp ensureTimestamp() => $_ensure(0);

  @$pb.TagNumber(2)
  $core.String get message => $_getSZ(1);
  @$pb.TagNumber(2)
  set message($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasMessage() => $_has(1);
  @$pb.TagNumber(2)
  void clearMessage() => clearField(2);

  @$pb.TagNumber(3)
  $3.User get recipient => $_getN(2);
  @$pb.TagNumber(3)
  set recipient($3.User v) { setField(3, v); }
  @$pb.TagNumber(3)
  $core.bool hasRecipient() => $_has(2);
  @$pb.TagNumber(3)
  void clearRecipient() => clearField(3);
  @$pb.TagNumber(3)
  $3.User ensureRecipient() => $_ensure(2);

  @$pb.TagNumber(4)
  $core.bool get private => $_getBF(3);
  @$pb.TagNumber(4)
  set private($core.bool v) { $_setBool(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasPrivate() => $_has(3);
  @$pb.TagNumber(4)
  void clearPrivate() => clearField(4);
}

class ConfirmationPacket extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'ConfirmationPacket', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'packets'), createEmptyInstance: create)
    ..aOM<$3.User>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'botUser', protoName: 'botUser', subBuilder: $3.User.create)
    ..pc<$3.User>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'users', $pb.PbFieldType.PM, subBuilder: $3.User.create)
    ..hasRequiredFields = false
  ;

  ConfirmationPacket._() : super();
  factory ConfirmationPacket({
    $3.User? botUser,
    $core.Iterable<$3.User>? users,
  }) {
    final _result = create();
    if (botUser != null) {
      _result.botUser = botUser;
    }
    if (users != null) {
      _result.users.addAll(users);
    }
    return _result;
  }
  factory ConfirmationPacket.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ConfirmationPacket.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  ConfirmationPacket clone() => ConfirmationPacket()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  ConfirmationPacket copyWith(void Function(ConfirmationPacket) updates) => super.copyWith((message) => updates(message as ConfirmationPacket)) as ConfirmationPacket; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static ConfirmationPacket create() => ConfirmationPacket._();
  ConfirmationPacket createEmptyInstance() => create();
  static $pb.PbList<ConfirmationPacket> createRepeated() => $pb.PbList<ConfirmationPacket>();
  @$core.pragma('dart2js:noInline')
  static ConfirmationPacket getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ConfirmationPacket>(create);
  static ConfirmationPacket? _defaultInstance;

  @$pb.TagNumber(1)
  $3.User get botUser => $_getN(0);
  @$pb.TagNumber(1)
  set botUser($3.User v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasBotUser() => $_has(0);
  @$pb.TagNumber(1)
  void clearBotUser() => clearField(1);
  @$pb.TagNumber(1)
  $3.User ensureBotUser() => $_ensure(0);

  @$pb.TagNumber(2)
  $core.List<$3.User> get users => $_getList(1);
}

class RegistrationPacket extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'RegistrationPacket', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'packets'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'trigger')
    ..pc<$4.Command>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'commands', $pb.PbFieldType.PM, subBuilder: $4.Command.create)
    ..hasRequiredFields = false
  ;

  RegistrationPacket._() : super();
  factory RegistrationPacket({
    $core.String? trigger,
    $core.Iterable<$4.Command>? commands,
  }) {
    final _result = create();
    if (trigger != null) {
      _result.trigger = trigger;
    }
    if (commands != null) {
      _result.commands.addAll(commands);
    }
    return _result;
  }
  factory RegistrationPacket.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory RegistrationPacket.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  RegistrationPacket clone() => RegistrationPacket()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  RegistrationPacket copyWith(void Function(RegistrationPacket) updates) => super.copyWith((message) => updates(message as RegistrationPacket)) as RegistrationPacket; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static RegistrationPacket create() => RegistrationPacket._();
  RegistrationPacket createEmptyInstance() => create();
  static $pb.PbList<RegistrationPacket> createRepeated() => $pb.PbList<RegistrationPacket>();
  @$core.pragma('dart2js:noInline')
  static RegistrationPacket getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<RegistrationPacket>(create);
  static RegistrationPacket? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get trigger => $_getSZ(0);
  @$pb.TagNumber(1)
  set trigger($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasTrigger() => $_has(0);
  @$pb.TagNumber(1)
  void clearTrigger() => clearField(1);

  @$pb.TagNumber(2)
  $core.List<$4.Command> get commands => $_getList(1);
}

