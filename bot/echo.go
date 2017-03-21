package bot

import (
    "github.com/bwmarrin/discordgo"
    "strings"
)
func echo(s *discordgo.Session, m *discordgo.MessageCreate, args []string) string {
    DeleteMessage(m.ChannelID, m.ID)
    return strings.Join(args, " ")
}
