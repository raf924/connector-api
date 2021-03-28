///
//  Generated code. Do not modify.
//  source: connector.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

class String extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'String', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'connector'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'string')
    ..hasRequiredFields = false
  ;

  String._() : super();
  factory String({
    $core.String? string,
  }) {
    final _result = create();
    if (string != null) {
      _result.string = string;
    }
    return _result;
  }
  factory String.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory String.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  String clone() => String()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  String copyWith(void Function(String) updates) => super.copyWith((message) => updates(message as String)) as String; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static String create() => String._();
  String createEmptyInstance() => create();
  static $pb.PbList<String> createRepeated() => $pb.PbList<String>();
  @$core.pragma('dart2js:noInline')
  static String getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<String>(create);
  static String? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get string => $_getSZ(0);
  @$pb.TagNumber(1)
  set string($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasString() => $_has(0);
  @$pb.TagNumber(1)
  void clearString() => clearField(1);
}

