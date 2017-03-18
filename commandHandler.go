package main

import (
    "github.com/bwmarrin/discordgo"
)

func handleCommand(s *discordgo.Session, m *discordgo.MessageCreate, b Bot) string {
    switch m.Content {
    case "!help":
        return "Commands: !ping, !help, !deploy"
    case "!ping":
        return "pong"
    default:
        return "Unknown command, type \"!help\" for help."
    }
}
