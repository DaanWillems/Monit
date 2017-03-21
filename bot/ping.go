package bot

import "github.com/bwmarrin/discordgo"

func ping(s *discordgo.Session, m *discordgo.MessageCreate, args []string) string {
    return "pong"
}
