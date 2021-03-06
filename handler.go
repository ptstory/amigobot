package amigobot

import "github.com/bwmarrin/discordgo"

//Session is an wrapper for discordgo.Session to make testing easier.
//If you need access to more funcs from discordgo.Session, add them here
type Session interface {
	// ChannelMessageSend sends a message to the given channel.
	// channelID : The ID of a Channel.
	// content   : The message to send.
	ChannelMessageSend(channelID string, content string) (*discordgo.Message, error)

	// MessageReactionAdd creates an emoji reaction to a message.
	// channelID : The channel ID.
	// messageID : The message ID.
	// emojiID   : Either the unicode emoji for the reaction, or a guild emoji identifier.
	MessageReactionAdd(channelID, messageID, emojiID string) error

	// ChannelMessagePin pins a message within a given channel.
	// channelID: The ID of a channel.
	// messageID: The ID of a message.
	ChannelMessagePin(channelID, messageID string) (err error)
}

//Handler describes a struct that is able to handle channel messages
type Handler interface {
	//Command is the string that triggers MessageHandle
	//if at the beginning of the message content
	Command() string
	//Handle is the action that is taken once the command has been triggered
	Handle(Session, *discordgo.MessageCreate)
}
