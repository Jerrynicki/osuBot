package commands

import "github.com/bwmarrin/discordgo"

// Ping Command
func Ping(session *discordgo.Session, message *discordgo.MessageCreate, args []string) {
	session.ChannelMessageSend(message.ChannelID, "Pong!")
}
