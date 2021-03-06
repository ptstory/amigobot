package mock

import "github.com/bwmarrin/discordgo"

//Session is a mock amigobot.Session for testing purposes
type Session struct {
	ChannelMessageSendFn func(channelID string, content string) (*discordgo.Message, error)
	MessageReactionAddFn func(channelID string, messageID string, emojiID string) error
	ChannelMessagePinFn  func(channelID, messageID string) (err error)
}

//ChannelMessageSend calls the provided mock implementation
func (s *Session) ChannelMessageSend(channelID string, content string) (*discordgo.Message, error) {
	return s.ChannelMessageSendFn(channelID, content)
}

//MessageReactionAdd calls the provided mock implementation
func (s *Session) MessageReactionAdd(channelID string, messageID string, emojiID string) error {
	return s.MessageReactionAddFn(channelID, messageID, emojiID)
}

//ChannelMessagePin calls the provided mock implementation
func (s *Session) ChannelMessagePin(channelID, messageID string) (err error) {
	return s.ChannelMessagePinFn(channelID, messageID)
}
