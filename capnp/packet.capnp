using Go = import "/go.capnp";
using Java = import "/java.capnp";
@0x824b0dcb3e553a2a;

$Go.package("connector");
$Go.import("github.com/raf924/connector-api");

$Java.package("tech.raf924.connectorapi");
$Java.outerClassname("Messages");

struct Timestamp {
    milliseconds @0 :UInt64;
}

struct Command {
    name    @0 :Text;
    aliases @1 :List(Text);
    usage   @2 :Text;
}

enum UserRole {
    regular @0;
    mod     @1;
    admin   @2;
}

struct User {
    nick    @0  :Text;
    id      @1  :Text;
    role    @2  :UserRole;
    joinedAt @3 :Timestamp;
}

struct IncomingMessagePacket {
    timestamp   @0 :Timestamp;
    message     @1 :Text;
    sender      @2 :User;
    private     @3 :Bool;
    recipients  @4 :List(User);
    mentionsConnectorUser @5 :Bool;
    incoming @6: Bool;
}

enum UserEvent {
    joined  @0;
    left    @1;
}

struct UserPacket {
    timestamp   @0 :Timestamp;
    event       @1 :UserEvent;
    user        @2 :User;
}

struct CommandPacket {
    timestamp   @0 :Timestamp;
    command     @1 :Text;
    sender      @2 :User;
    private     @3 :Bool;
    args        @4 :List(Text);
    argString  @5 :Text;
}

struct OutgoingMessagePacket {
    message     @0 :Text;
    recipient   @1 :User;
    private     @2 :Bool;
}

struct RegistrationPacket {
    commands @0 :List(Command);
}

struct ConfirmationPacket {
    botUser @0 :User;
    trigger @1 :Text;
    users   @2 :List(User);
}