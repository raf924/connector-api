package connector

import (
	"capnproto.org/go/capnp/v3"
	"fmt"
	"github.com/raf924/bot/pkg/domain"
	"log"
	"time"
)

func MapCommandMessageToDTO(message *domain.CommandMessage, messageDTO CommandPacket) error {
	err := messageDTO.SetCommand(message.Command())
	sender, err := messageDTO.NewSender()
	if err != nil {
		return err
	}
	err = MapUserToDTO(message.Sender(), sender)
	if err != nil {
		return err
	}
	timestamp, err := messageDTO.NewTimestamp()
	if err != nil {
		return err
	}
	err = mapTimeToTimestamp(message.Timestamp(), &timestamp)
	if err != nil {
		return err
	}
	messageDTO.SetPrivate(message.Private())
	args, err := messageDTO.NewArgs(int32(len(message.Args())))
	if err != nil {
		return err
	}
	for i, arg := range message.Args() {
		err := args.Set(i, arg)
		if err != nil {
			return err
		}
	}
	err = messageDTO.SetArgString(message.ArgString())
	if err != nil {
		return err
	}
	return nil
}

func MapChatMessageToDTO(message *domain.ChatMessage, messageDTO IncomingMessagePacket) error {
	err := messageDTO.SetMessage(message.Message())
	if err != nil {
		return err
	}
	messageDTO.SetPrivate(message.Private())
	sender, err := messageDTO.NewSender()
	if err != nil {
		return err
	}
	err = MapUserToDTO(message.Sender(), sender)
	if err != nil {
		return err
	}
	timestamp, err := messageDTO.NewTimestamp()
	if err != nil {
		return err
	}
	err = mapTimeToTimestamp(message.Timestamp(), &timestamp)
	if err != nil {
		return err
	}
	recipients, err := messageDTO.NewRecipients(int32(len(message.Recipients())))
	for i, recipient := range message.Recipients() {
		recipientDTO, err := NewUser(messageDTO.Segment())
		if err != nil {
			return err
		}
		err = MapUserToDTO(recipient, recipientDTO)
		if err != nil {
			return err
		}
		err = recipients.Set(i, recipientDTO)
		if err != nil {
			return err
		}
	}
	messageDTO.SetMentionsConnectorUser(message.MentionsConnectorUser())
	messageDTO.SetIncoming(message.Incoming())
	return nil
}

func MapUserEventToDTO(event *domain.UserEvent, userEventDTO UserPacket) error {
	user, err := userEventDTO.NewUser()
	if err != nil {
		return err
	}
	err = MapUserToDTO(event.User(), user)
	if err != nil {
		return err
	}
	timestamp, err := userEventDTO.NewTimestamp()
	if err != nil {
		return err
	}
	err = mapTimeToTimestamp(event.Timestamp(), &timestamp)
	if err != nil {
		return err
	}
	userEventDTO.SetEvent(MapEventTypeToDTO(event.EventType()))
	return nil
}

func MapUserToDTO(user *domain.User, userDTO User) error {
	err := userDTO.SetNick(user.Nick())
	if err != nil {
		return err
	}
	err = userDTO.SetId(user.Id())
	if err != nil {
		return err
	}
	userDTO.SetRole(MapRoleToDTO(user.Role()))
	if user.JoinedAt() != nil {
		timestamp, err := userDTO.NewJoinedAt()
		if err != nil {
			return err
		}
		err = mapTimeToTimestamp(*user.JoinedAt(), &timestamp)
		if err != nil {
			return err
		}
	}
	return nil
}

func MapEventTypeToDTO(eventType domain.UserEventType) UserEvent {
	switch eventType {
	case domain.UserJoined:
		return UserEvent_joined
	case domain.UserLeft:
		return UserEvent_left
	}
	panic(fmt.Errorf("unknown eventType %s", eventType))
}

func MapRoleToDTO(userRole domain.UserRole) UserRole {
	switch userRole {
	case domain.RegularUser:
		return UserRole_regular
	case domain.Moderator:
		return UserRole_mod
	case domain.Admin:
		return UserRole_admin
	}
	return UserRole_regular
}

func MapToRole(userRoleDTO UserRole) domain.UserRole {
	switch userRoleDTO {
	case UserRole_regular:
		return domain.RegularUser
	case UserRole_mod:
		return domain.Moderator
	case UserRole_admin:
		return domain.Admin
	}
	return domain.RegularUser
}

func mapTimeToTimestamp(time time.Time, timestamp *Timestamp) error {
	timestamp.SetMilliseconds(uint64(time.UTC().UnixMilli()))
	return nil
}

func CreateConfirmationPacket(botUser *domain.User, trigger string, users []*domain.User, confirmationPacket ConfirmationPacket) error {
	botUserDTO, _ := confirmationPacket.NewBotUser()
	err := MapUserToDTO(botUser, botUserDTO)
	if err != nil {
		return err
	}
	err = confirmationPacket.SetTrigger(trigger)
	if err != nil {
		return err
	}
	userList, err := confirmationPacket.NewUsers(int32(len(users)))
	if err != nil {
		return err
	}
	for i, user := range users {
		userDTO, _ := NewUser(confirmationPacket.Segment())
		err := MapUserToDTO(user, userDTO)
		if err != nil {
			return err
		}
		err = userList.Set(i, userDTO)
		if err != nil {
			return err
		}
	}
	return nil
}

func MapDTOToCommand(command Command) (*domain.Command, error) {
	name, err := command.Name()
	if err != nil {
		return nil, err
	}
	aliasesDTO, err := command.Aliases()
	if err != nil {
		return nil, err
	}
	usage, err := command.Usage()
	if err != nil {
		return nil, err
	}
	aliases := make([]string, aliasesDTO.Len())
	for i := 0; i < aliasesDTO.Len(); i++ {
		aliases[i], err = aliasesDTO.At(i)
		if err != nil {
			return nil, err
		}
	}
	return domain.NewCommand(name, aliases, usage), nil
}

func MapDTOToCommandList(commands Command_List) domain.CommandList {
	commandList := domain.NewCommandList()
	for i := 0; i < commands.Len(); i++ {
		command, err := MapDTOToCommand(commands.At(i))
		if err != nil {
			log.Println(err)
			continue
		}
		commandList.Add(command)
	}
	return commandList
}

func MapClientMessageToDTO(message *domain.ClientMessage, packet OutgoingMessagePacket) error {
	err := packet.SetMessage(message.Message())
	if err != nil {
		return err
	}
	if message.Recipient() != nil {
		recipient, err := packet.NewRecipient()
		if err != nil {
			return err
		}
		err = MapUserToDTO(message.Recipient(), recipient)
		if err != nil {
			return err
		}
	}
	packet.SetPrivate(message.Private())
	return nil
}

func MapDTOToChatMessage(packet IncomingMessagePacket) *domain.ChatMessage {
	message, err := packet.Message()
	if err != nil {
		return nil
	}
	recipients, err := packet.Recipients()
	if err != nil {
		return nil
	}
	sender, err := packet.Sender()
	if err != nil {
		return nil
	}

	timestamp, err := packet.Timestamp()
	if err != nil {
		return nil
	}
	return domain.NewChatMessage(message, MapDTOToUser(sender), MapDTOToUsers(recipients), packet.MentionsConnectorUser(), packet.Private(), MapDTOToTime(timestamp), packet.Incoming())
}

func MapDTOToCommandMessage(packet CommandPacket) *domain.CommandMessage {
	args, err := packet.Args()
	if err != nil {
		return nil
	}
	argString, err := packet.ArgString()
	if err != nil {
		return nil
	}
	command, err := packet.Command()
	if err != nil {
		return nil
	}
	timestamp, err := packet.Timestamp()
	if err != nil {
		return nil
	}
	sender, err := packet.Sender()
	if err != nil {
		return nil
	}
	return domain.NewCommandMessage(command, MapDTOToStrings(args), argString, MapDTOToUser(sender), packet.Private(), MapDTOToTime(timestamp))
}

func MapDTOToUserEvent(packet UserPacket) *domain.UserEvent {
	user, err := packet.User()
	if err != nil {
		return nil
	}
	timestamp, err := packet.Timestamp()
	if err != nil {
		return nil
	}
	return domain.NewUserEvent(MapDTOToUser(user), MapDTOToUserEventType(packet.Event()), MapDTOToTime(timestamp))
}

func MapDTOToUserEventType(event UserEvent) domain.UserEventType {
	switch event {
	case UserEvent_joined:
		return domain.UserJoined
	case UserEvent_left:
		return domain.UserLeft
	}
	return ""
}

func MapDTOToStrings(list capnp.TextList) []string {
	ss := make([]string, list.Len())
	for i := 0; i < list.Len(); i++ {
		var err error
		ss[i], err = list.At(i)
		if err != nil {
			return nil
		}
	}
	return ss
}

func MapDTOToClientMessage(packet OutgoingMessagePacket) *domain.ClientMessage {
	if !packet.IsValid() {
		return nil
	}
	message, err := packet.Message()
	if err != nil {
		return nil
	}
	var user *domain.User
	if packet.HasRecipient() {
		recipient, err := packet.Recipient()
		if err != nil {
			return nil
		}
		user = MapDTOToUser(recipient)
	}
	return domain.NewClientMessage(message, user, packet.Private())
}

func MapDTOToTime(packet Timestamp) time.Time {
	return time.UnixMilli(int64(packet.Milliseconds()))
}

func MapCommandToDTO(command *domain.Command, dto *Command) error {
	err := dto.SetName(command.Name())
	if err != nil {
		return err
	}
	err = dto.SetUsage(command.Usage())
	if err != nil {
		return err
	}
	aliases, err := dto.NewAliases(int32(len(command.Aliases())))
	if err != nil {
		return err
	}
	for i := 0; i < aliases.Len(); i++ {
		err := aliases.Set(i, command.Aliases()[i])
		if err != nil {
			return err
		}
	}
	return nil
}

func MapDTOToUsers(list User_List) []*domain.User {
	users := make([]*domain.User, list.Len())
	for i := 0; i < list.Len(); i++ {
		users[i] = MapDTOToUser(list.At(i))
	}
	return users
}

func MapCommandsToDTO(commands []*domain.Command, list *Command_List) error {
	for i := 0; i < len(commands); i++ {
		command := list.At(i)
		err := MapCommandToDTO(commands[i], &command)
		if err != nil {
			return err
		}
	}
	return nil
}

func MapRegistrationToDTO(message *domain.RegistrationMessage, dto *RegistrationPacket) error {
	commands, err := dto.NewCommands(int32(len(message.Commands())))
	if err != nil {
		return err
	}
	err = MapCommandsToDTO(message.Commands(), &commands)
	if err != nil {
		return err
	}
	return nil
}

func MapDTOToUser(recipient User) *domain.User {
	nick, err := recipient.Nick()
	if err != nil {
		return nil
	}
	id, err := recipient.Id()
	if err != nil {
		return nil
	}
	return domain.NewUser(nick, id, MapToRole(recipient.Role()))
}

func MapDTOToConfirmationMessage(packet ConfirmationPacket) *domain.ConfirmationMessage {
	user, err := packet.BotUser()
	if err != nil {
		return nil
	}
	users, err := packet.Users()
	if err != nil {
		return nil
	}
	trigger, err := packet.Trigger()
	if err != nil {
		return nil
	}
	return domain.NewConfirmationMessage(MapDTOToUser(user), trigger, MapDTOToUsers(users))
}
