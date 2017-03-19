package bot

import (
    "github.com/bwmarrin/discordgo"
    "monit/config"
)

func handleCommand(s *discordgo.Session, m *discordgo.MessageCreate) string {
    switch m.Content {
    case "!echo":
        return m.Content[1:len(m.Content)]
    case "!help":
        return config.GetText("help")
    case "!ping":
        return config.GetText("ping")
    default:
        return config.GetText("unknown")
    }
}
