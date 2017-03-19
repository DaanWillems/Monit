package commands 

import "github.com/bwmarrin/discordgo"

func echo(s *discordgo.Session, m *discordgo.MessageCreate) string {
    return "pong"
}
