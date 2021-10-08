using Go = import "/go.capnp";
using Java = import "/java.capnp";
@0x962b7a64e96c7802;

$Go.package("connector");
$Go.import("github.com/raf924/connector-api");

$Java.package("tech.raf924.connectorapi");


using Packet = import "/packet.capnp";

interface Connector {
    register        @0 (registration :Packet.RegistrationPacket) -> (confirmation :Packet.ConfirmationPacket);
    send            @1 (message  :Packet.OutgoingMessagePacket) -> ();
    messageStream   @2 (receiver :Receiver(Packet.IncomingMessagePacket)) -> ();
    commandStream   @3 (receiver :Receiver(Packet.CommandPacket)) -> ();
    eventStream     @4 (receiver :Receiver(Packet.UserPacket)) -> ();

    interface Receiver(Message) {
        receive @0 (message: Message) -> ();
    }
}

interface Server {
    send @0 (receiver :Receiver(Packet.IncomingMessagePacket)) -> ();
    interface Receiver(Message) {
        receive @0 (message: Message) -> ();
    }
}