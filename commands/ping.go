package commands 

import "github.com/bwmarrin/discordgo"

func ping(s *discordgo.Session, m *discordgo.MessageCreate) string {
    return "pong"
}
