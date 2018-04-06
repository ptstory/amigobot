package greet

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

//Handler handles the ?greet [name] command
type Handler struct{}

//Command is the trigger for the greet message
func (h *Handler) Command() string {
	return "?greet "
}

//Handle greets the person specified
func (h *Handler) Handle(s *discordgo.Session, m *discordgo.MessageCreate) {
	toGreet := strings.TrimPrefix(m.Content, h.Command())
	s.ChannelMessageSend(m.ChannelID, "Ho there, "+toGreet+"!")
}
