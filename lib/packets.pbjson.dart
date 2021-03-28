///
//  Generated code. Do not modify.
//  source: packets.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields,deprecated_member_use_from_same_package

import 'dart:core' as $core;
import 'dart:convert' as $convert;
import 'dart:typed_data' as $typed_data;
@$core.Deprecated('Use userEventDescriptor instead')
const UserEvent$json = const {
  '1': 'UserEvent',
  '2': const [
    const {'1': 'JOINED', '2': 0},
    const {'1': 'LEFT', '2': 1},
  ],
};

/// Descriptor for `UserEvent`. Decode as a `google.protobuf.EnumDescriptorProto`.
final $typed_data.Uint8List userEventDescriptor = $convert.base64Decode('CglVc2VyRXZlbnQSCgoGSk9JTkVEEAASCAoETEVGVBAB');
@$core.Deprecated('Use messagePacketDescriptor instead')
const MessagePacket$json = const {
  '1': 'MessagePacket',
  '2': const [
    const {'1': 'timestamp', '3': 1, '4': 1, '5': 11, '6': '.google.protobuf.Timestamp', '10': 'timestamp'},
    const {'1': 'message', '3': 2, '4': 1, '5': 9, '10': 'message'},
    const {'1': 'user', '3': 3, '4': 1, '5': 11, '6': '.user.User', '10': 'user'},
    const {'1': 'private', '3': 4, '4': 1, '5': 8, '10': 'private'},
  ],
};

/// Descriptor for `MessagePacket`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List messagePacketDescriptor = $convert.base64Decode('Cg1NZXNzYWdlUGFja2V0EjgKCXRpbWVzdGFtcBgBIAEoCzIaLmdvb2dsZS5wcm90b2J1Zi5UaW1lc3RhbXBSCXRpbWVzdGFtcBIYCgdtZXNzYWdlGAIgASgJUgdtZXNzYWdlEh4KBHVzZXIYAyABKAsyCi51c2VyLlVzZXJSBHVzZXISGAoHcHJpdmF0ZRgEIAEoCFIHcHJpdmF0ZQ==');
@$core.Deprecated('Use userPacketDescriptor instead')
const UserPacket$json = const {
  '1': 'UserPacket',
  '2': const [
    const {'1': 'timestamp', '3': 1, '4': 1, '5': 11, '6': '.google.protobuf.Timestamp', '10': 'timestamp'},
    const {'1': 'user', '3': 2, '4': 1, '5': 11, '6': '.user.User', '10': 'user'},
    const {'1': 'event', '3': 3, '4': 1, '5': 14, '6': '.packets.UserEvent', '10': 'event'},
  ],
};

/// Descriptor for `UserPacket`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List userPacketDescriptor = $convert.base64Decode('CgpVc2VyUGFja2V0EjgKCXRpbWVzdGFtcBgBIAEoCzIaLmdvb2dsZS5wcm90b2J1Zi5UaW1lc3RhbXBSCXRpbWVzdGFtcBIeCgR1c2VyGAIgASgLMgoudXNlci5Vc2VyUgR1c2VyEigKBWV2ZW50GAMgASgOMhIucGFja2V0cy5Vc2VyRXZlbnRSBWV2ZW50');
@$core.Deprecated('Use commandPacketDescriptor instead')
const CommandPacket$json = const {
  '1': 'CommandPacket',
  '2': const [
    const {'1': 'timestamp', '3': 1, '4': 1, '5': 11, '6': '.google.protobuf.Timestamp', '10': 'timestamp'},
    const {'1': 'command', '3': 2, '4': 1, '5': 9, '10': 'command'},
    const {'1': 'args', '3': 3, '4': 3, '5': 9, '10': 'args'},
    const {'1': 'user', '3': 4, '4': 1, '5': 11, '6': '.user.User', '10': 'user'},
    const {'1': 'private', '3': 5, '4': 1, '5': 8, '10': 'private'},
    const {'1': 'argString', '3': 6, '4': 1, '5': 9, '10': 'argString'},
  ],
};

/// Descriptor for `CommandPacket`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List commandPacketDescriptor = $convert.base64Decode('Cg1Db21tYW5kUGFja2V0EjgKCXRpbWVzdGFtcBgBIAEoCzIaLmdvb2dsZS5wcm90b2J1Zi5UaW1lc3RhbXBSCXRpbWVzdGFtcBIYCgdjb21tYW5kGAIgASgJUgdjb21tYW5kEhIKBGFyZ3MYAyADKAlSBGFyZ3MSHgoEdXNlchgEIAEoCzIKLnVzZXIuVXNlclIEdXNlchIYCgdwcml2YXRlGAUgASgIUgdwcml2YXRlEhwKCWFyZ1N0cmluZxgGIAEoCVIJYXJnU3RyaW5n');
@$core.Deprecated('Use botPacketDescriptor instead')
const BotPacket$json = const {
  '1': 'BotPacket',
  '2': const [
    const {'1': 'timestamp', '3': 1, '4': 1, '5': 11, '6': '.google.protobuf.Timestamp', '10': 'timestamp'},
    const {'1': 'message', '3': 2, '4': 1, '5': 9, '10': 'message'},
    const {'1': 'recipient', '3': 3, '4': 1, '5': 11, '6': '.user.User', '10': 'recipient'},
    const {'1': 'private', '3': 4, '4': 1, '5': 8, '10': 'private'},
  ],
};

/// Descriptor for `BotPacket`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List botPacketDescriptor = $convert.base64Decode('CglCb3RQYWNrZXQSOAoJdGltZXN0YW1wGAEgASgLMhouZ29vZ2xlLnByb3RvYnVmLlRpbWVzdGFtcFIJdGltZXN0YW1wEhgKB21lc3NhZ2UYAiABKAlSB21lc3NhZ2USKAoJcmVjaXBpZW50GAMgASgLMgoudXNlci5Vc2VyUglyZWNpcGllbnQSGAoHcHJpdmF0ZRgEIAEoCFIHcHJpdmF0ZQ==');
@$core.Deprecated('Use confirmationPacketDescriptor instead')
const ConfirmationPacket$json = const {
  '1': 'ConfirmationPacket',
  '2': const [
    const {'1': 'botUser', '3': 1, '4': 1, '5': 11, '6': '.user.User', '10': 'botUser'},
    const {'1': 'users', '3': 2, '4': 3, '5': 11, '6': '.user.User', '10': 'users'},
  ],
};

/// Descriptor for `ConfirmationPacket`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List confirmationPacketDescriptor = $convert.base64Decode('ChJDb25maXJtYXRpb25QYWNrZXQSJAoHYm90VXNlchgBIAEoCzIKLnVzZXIuVXNlclIHYm90VXNlchIgCgV1c2VycxgCIAMoCzIKLnVzZXIuVXNlclIFdXNlcnM=');
@$core.Deprecated('Use registrationPacketDescriptor instead')
const RegistrationPacket$json = const {
  '1': 'RegistrationPacket',
  '2': const [
    const {'1': 'trigger', '3': 1, '4': 1, '5': 9, '10': 'trigger'},
    const {'1': 'commands', '3': 2, '4': 3, '5': 11, '6': '.command.Command', '10': 'commands'},
  ],
};

/// Descriptor for `RegistrationPacket`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List registrationPacketDescriptor = $convert.base64Decode('ChJSZWdpc3RyYXRpb25QYWNrZXQSGAoHdHJpZ2dlchgBIAEoCVIHdHJpZ2dlchIsCghjb21tYW5kcxgCIAMoCzIQLmNvbW1hbmQuQ29tbWFuZFIIY29tbWFuZHM=');
