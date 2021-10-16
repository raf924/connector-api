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
}

interface Dispatcher {
    dispatchMessage @0 (message: Packet.IncomingMessagePacket) -> ();
    dispatchCommand @1 (command: Packet.CommandPacket) -> ();
    dispatchEvent   @2 (event: Packet.UserPacket) -> ();
}