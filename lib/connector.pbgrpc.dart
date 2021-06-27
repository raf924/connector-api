///
//  Generated code. Do not modify.
//  source: connector.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields

import 'dart:async' as $async;

import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'google/protobuf/empty.pb.dart' as $0;
import 'packets.pb.dart' as $1;
export 'connector.pb.dart';

class ConnectorClient extends $grpc.Client {
  static final _$connect = $grpc.ClientMethod<$0.Empty, $0.Empty>(
      '/connector.Connector/Connect',
      ($0.Empty value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.Empty.fromBuffer(value));
  static final _$register =
      $grpc.ClientMethod<$1.RegistrationPacket, $1.ConfirmationPacket>(
          '/connector.Connector/Register',
          ($1.RegistrationPacket value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $1.ConfirmationPacket.fromBuffer(value));
  static final _$readMessages = $grpc.ClientMethod<$0.Empty, $1.MessagePacket>(
      '/connector.Connector/ReadMessages',
      ($0.Empty value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $1.MessagePacket.fromBuffer(value));
  static final _$readCommands = $grpc.ClientMethod<$0.Empty, $1.CommandPacket>(
      '/connector.Connector/ReadCommands',
      ($0.Empty value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $1.CommandPacket.fromBuffer(value));
  static final _$readUserEvents = $grpc.ClientMethod<$0.Empty, $1.UserPacket>(
      '/connector.Connector/ReadUserEvents',
      ($0.Empty value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $1.UserPacket.fromBuffer(value));
  static final _$sendMessage = $grpc.ClientMethod<$1.BotPacket, $0.Empty>(
      '/connector.Connector/SendMessage',
      ($1.BotPacket value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.Empty.fromBuffer(value));
  static final _$ping = $grpc.ClientMethod<$0.Empty, $0.Empty>(
      '/connector.Connector/Ping',
      ($0.Empty value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.Empty.fromBuffer(value));

  ConnectorClient($grpc.ClientChannel channel,
      {$grpc.CallOptions? options,
      $core.Iterable<$grpc.ClientInterceptor>? interceptors})
      : super(channel, options: options, interceptors: interceptors);

  $grpc.ResponseFuture<$0.Empty> connect($0.Empty request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$connect, request, options: options);
  }

  $grpc.ResponseFuture<$1.ConfirmationPacket> register(
      $1.RegistrationPacket request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$register, request, options: options);
  }

  $grpc.ResponseStream<$1.MessagePacket> readMessages($0.Empty request,
      {$grpc.CallOptions? options}) {
    return $createStreamingCall(
        _$readMessages, $async.Stream.fromIterable([request]),
        options: options);
  }

  $grpc.ResponseStream<$1.CommandPacket> readCommands($0.Empty request,
      {$grpc.CallOptions? options}) {
    return $createStreamingCall(
        _$readCommands, $async.Stream.fromIterable([request]),
        options: options);
  }

  $grpc.ResponseStream<$1.UserPacket> readUserEvents($0.Empty request,
      {$grpc.CallOptions? options}) {
    return $createStreamingCall(
        _$readUserEvents, $async.Stream.fromIterable([request]),
        options: options);
  }

  $grpc.ResponseFuture<$0.Empty> sendMessage($1.BotPacket request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$sendMessage, request, options: options);
  }

  $grpc.ResponseFuture<$0.Empty> ping($0.Empty request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$ping, request, options: options);
  }
}

abstract class ConnectorServiceBase extends $grpc.Service {
  $core.String get $name => 'connector.Connector';

  ConnectorServiceBase() {
    $addMethod($grpc.ServiceMethod<$0.Empty, $0.Empty>(
        'Connect',
        connect_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.Empty.fromBuffer(value),
        ($0.Empty value) => value.writeToBuffer()));
    $addMethod(
        $grpc.ServiceMethod<$1.RegistrationPacket, $1.ConfirmationPacket>(
            'Register',
            register_Pre,
            false,
            false,
            ($core.List<$core.int> value) =>
                $1.RegistrationPacket.fromBuffer(value),
            ($1.ConfirmationPacket value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.Empty, $1.MessagePacket>(
        'ReadMessages',
        readMessages_Pre,
        false,
        true,
        ($core.List<$core.int> value) => $0.Empty.fromBuffer(value),
        ($1.MessagePacket value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.Empty, $1.CommandPacket>(
        'ReadCommands',
        readCommands_Pre,
        false,
        true,
        ($core.List<$core.int> value) => $0.Empty.fromBuffer(value),
        ($1.CommandPacket value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.Empty, $1.UserPacket>(
        'ReadUserEvents',
        readUserEvents_Pre,
        false,
        true,
        ($core.List<$core.int> value) => $0.Empty.fromBuffer(value),
        ($1.UserPacket value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$1.BotPacket, $0.Empty>(
        'SendMessage',
        sendMessage_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $1.BotPacket.fromBuffer(value),
        ($0.Empty value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.Empty, $0.Empty>(
        'Ping',
        ping_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.Empty.fromBuffer(value),
        ($0.Empty value) => value.writeToBuffer()));
  }

  $async.Future<$0.Empty> connect_Pre(
      $grpc.ServiceCall call, $async.Future<$0.Empty> request) async {
    return connect(call, await request);
  }

  $async.Future<$1.ConfirmationPacket> register_Pre($grpc.ServiceCall call,
      $async.Future<$1.RegistrationPacket> request) async {
    return register(call, await request);
  }

  $async.Stream<$1.MessagePacket> readMessages_Pre(
      $grpc.ServiceCall call, $async.Future<$0.Empty> request) async* {
    yield* readMessages(call, await request);
  }

  $async.Stream<$1.CommandPacket> readCommands_Pre(
      $grpc.ServiceCall call, $async.Future<$0.Empty> request) async* {
    yield* readCommands(call, await request);
  }

  $async.Stream<$1.UserPacket> readUserEvents_Pre(
      $grpc.ServiceCall call, $async.Future<$0.Empty> request) async* {
    yield* readUserEvents(call, await request);
  }

  $async.Future<$0.Empty> sendMessage_Pre(
      $grpc.ServiceCall call, $async.Future<$1.BotPacket> request) async {
    return sendMessage(call, await request);
  }

  $async.Future<$0.Empty> ping_Pre(
      $grpc.ServiceCall call, $async.Future<$0.Empty> request) async {
    return ping(call, await request);
  }

  $async.Future<$0.Empty> connect($grpc.ServiceCall call, $0.Empty request);
  $async.Future<$1.ConfirmationPacket> register(
      $grpc.ServiceCall call, $1.RegistrationPacket request);
  $async.Stream<$1.MessagePacket> readMessages(
      $grpc.ServiceCall call, $0.Empty request);
  $async.Stream<$1.CommandPacket> readCommands(
      $grpc.ServiceCall call, $0.Empty request);
  $async.Stream<$1.UserPacket> readUserEvents(
      $grpc.ServiceCall call, $0.Empty request);
  $async.Future<$0.Empty> sendMessage(
      $grpc.ServiceCall call, $1.BotPacket request);
  $async.Future<$0.Empty> ping($grpc.ServiceCall call, $0.Empty request);
}
