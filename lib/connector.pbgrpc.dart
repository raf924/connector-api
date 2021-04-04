///
//  Generated code. Do not modify.
//  source: connector.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields

import 'dart:async' as $async;

import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'packets.pb.dart' as $0;
import 'google/protobuf/empty.pb.dart' as $1;
export 'connector.pb.dart';

class ConnectorClient extends $grpc.Client {
  static final _$register =
      $grpc.ClientMethod<$0.RegistrationPacket, $0.ConfirmationPacket>(
          '/connector.Connector/Register',
          ($0.RegistrationPacket value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $0.ConfirmationPacket.fromBuffer(value));
  static final _$readMessages = $grpc.ClientMethod<$1.Empty, $0.MessagePacket>(
      '/connector.Connector/ReadMessages',
      ($1.Empty value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.MessagePacket.fromBuffer(value));
  static final _$readCommands = $grpc.ClientMethod<$1.Empty, $0.CommandPacket>(
      '/connector.Connector/ReadCommands',
      ($1.Empty value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.CommandPacket.fromBuffer(value));
  static final _$readUserEvents = $grpc.ClientMethod<$1.Empty, $0.UserPacket>(
      '/connector.Connector/ReadUserEvents',
      ($1.Empty value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.UserPacket.fromBuffer(value));
  static final _$sendMessage = $grpc.ClientMethod<$0.BotPacket, $1.Empty>(
      '/connector.Connector/SendMessage',
      ($0.BotPacket value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $1.Empty.fromBuffer(value));
  static final _$ping = $grpc.ClientMethod<$1.Empty, $1.Empty>(
      '/connector.Connector/Ping',
      ($1.Empty value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $1.Empty.fromBuffer(value));

  ConnectorClient($grpc.ClientChannel channel,
      {$grpc.CallOptions? options,
      $core.Iterable<$grpc.ClientInterceptor>? interceptors})
      : super(channel, options: options, interceptors: interceptors);

  $grpc.ResponseFuture<$0.ConfirmationPacket> register(
      $0.RegistrationPacket request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$register, request, options: options);
  }

  $grpc.ResponseStream<$0.MessagePacket> readMessages($1.Empty request,
      {$grpc.CallOptions? options}) {
    return $createStreamingCall(
        _$readMessages, $async.Stream.fromIterable([request]),
        options: options);
  }

  $grpc.ResponseStream<$0.CommandPacket> readCommands($1.Empty request,
      {$grpc.CallOptions? options}) {
    return $createStreamingCall(
        _$readCommands, $async.Stream.fromIterable([request]),
        options: options);
  }

  $grpc.ResponseStream<$0.UserPacket> readUserEvents($1.Empty request,
      {$grpc.CallOptions? options}) {
    return $createStreamingCall(
        _$readUserEvents, $async.Stream.fromIterable([request]),
        options: options);
  }

  $grpc.ResponseFuture<$1.Empty> sendMessage($0.BotPacket request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$sendMessage, request, options: options);
  }

  $grpc.ResponseFuture<$1.Empty> ping($1.Empty request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$ping, request, options: options);
  }
}

abstract class ConnectorServiceBase extends $grpc.Service {
  $core.String get $name => 'connector.Connector';

  ConnectorServiceBase() {
    $addMethod(
        $grpc.ServiceMethod<$0.RegistrationPacket, $0.ConfirmationPacket>(
            'Register',
            register_Pre,
            false,
            false,
            ($core.List<$core.int> value) =>
                $0.RegistrationPacket.fromBuffer(value),
            ($0.ConfirmationPacket value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$1.Empty, $0.MessagePacket>(
        'ReadMessages',
        readMessages_Pre,
        false,
        true,
        ($core.List<$core.int> value) => $1.Empty.fromBuffer(value),
        ($0.MessagePacket value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$1.Empty, $0.CommandPacket>(
        'ReadCommands',
        readCommands_Pre,
        false,
        true,
        ($core.List<$core.int> value) => $1.Empty.fromBuffer(value),
        ($0.CommandPacket value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$1.Empty, $0.UserPacket>(
        'ReadUserEvents',
        readUserEvents_Pre,
        false,
        true,
        ($core.List<$core.int> value) => $1.Empty.fromBuffer(value),
        ($0.UserPacket value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.BotPacket, $1.Empty>(
        'SendMessage',
        sendMessage_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.BotPacket.fromBuffer(value),
        ($1.Empty value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$1.Empty, $1.Empty>(
        'Ping',
        ping_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $1.Empty.fromBuffer(value),
        ($1.Empty value) => value.writeToBuffer()));
  }

  $async.Future<$0.ConfirmationPacket> register_Pre($grpc.ServiceCall call,
      $async.Future<$0.RegistrationPacket> request) async {
    return register(call, await request);
  }

  $async.Stream<$0.MessagePacket> readMessages_Pre(
      $grpc.ServiceCall call, $async.Future<$1.Empty> request) async* {
    yield* readMessages(call, await request);
  }

  $async.Stream<$0.CommandPacket> readCommands_Pre(
      $grpc.ServiceCall call, $async.Future<$1.Empty> request) async* {
    yield* readCommands(call, await request);
  }

  $async.Stream<$0.UserPacket> readUserEvents_Pre(
      $grpc.ServiceCall call, $async.Future<$1.Empty> request) async* {
    yield* readUserEvents(call, await request);
  }

  $async.Future<$1.Empty> sendMessage_Pre(
      $grpc.ServiceCall call, $async.Future<$0.BotPacket> request) async {
    return sendMessage(call, await request);
  }

  $async.Future<$1.Empty> ping_Pre(
      $grpc.ServiceCall call, $async.Future<$1.Empty> request) async {
    return ping(call, await request);
  }

  $async.Future<$0.ConfirmationPacket> register(
      $grpc.ServiceCall call, $0.RegistrationPacket request);
  $async.Stream<$0.MessagePacket> readMessages(
      $grpc.ServiceCall call, $1.Empty request);
  $async.Stream<$0.CommandPacket> readCommands(
      $grpc.ServiceCall call, $1.Empty request);
  $async.Stream<$0.UserPacket> readUserEvents(
      $grpc.ServiceCall call, $1.Empty request);
  $async.Future<$1.Empty> sendMessage(
      $grpc.ServiceCall call, $0.BotPacket request);
  $async.Future<$1.Empty> ping($grpc.ServiceCall call, $1.Empty request);
}
