package commands

import (
    "github.com/bwmarrin/discordgo"
)
var commandMap map[string]func(s *discordgo.Session, m *discordgo.MessageCreate) string

func LoadCommands() {
    commandMap = map[string]func(s *discordgo.Session, m *discordgo.MessageCreate) string {
        "!ping":     ping,
        "!echo":     echo,
        "!deploy":   deploy,
        "!help":     help,
        "!source":   source,
    }
}

func EvaluateCommand(cmd string, s *discordgo.Session, m *discordgo.MessageCreate) string {
    if _, ok := commandMap[cmd]; !ok{  
        return commandMap["!help"](s, m)
    }
    return commandMap[cmd](s, m)
}
