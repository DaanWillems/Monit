package main

import (
    "github.com/bwmarrin/discordgo"
)

func handleCommand(s *discordgo.Session, m *discordgo.MessageCreate, b Bot) string {
    switch m.Content {
    case "!help":
        return getText("help")
    case "!ping":
        return getText("ping")
    default:
        return getText("unknown")
    }
}
