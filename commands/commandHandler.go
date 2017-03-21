package commands

import (
    "github.com/bwmarrin/discordgo"
)
var commandMap map[string]func(s *discordgo.Session, m *discordgo.MessageCreate, args []string) string

func LoadCommands() {
    commandMap = map[string]func(s *discordgo.Session, m *discordgo.MessageCreate, args []string) string {
        "!ping":     ping,
        "!echo":     echo,
        "!deploy":   deploy,
        "!help":     help,
        "!source":   source,
    }
}

func EvaluateCommand(s *discordgo.Session, m *discordgo.MessageCreate) string {
    args, cmd := parseCommand(m.Content)

    if _, ok := commandMap[cmd]; !ok{
        return commandMap["!help"](s, m, args)
    }

    return commandMap[cmd](s, m, args)
}
