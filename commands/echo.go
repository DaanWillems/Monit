package commands 

import (
    "github.com/bwmarrin/discordgo"
    "strings"
)
func echo(s *discordgo.Session, m *discordgo.MessageCreate, args []string) string {
    return strings.Join(args, " ")
}
