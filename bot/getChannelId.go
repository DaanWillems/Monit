package bot

import (
    "github.com/bwmarrin/discordgo"
)
func getChannelId(s *discordgo.Session, m *discordgo.MessageCreate, args []string) string {
    return m.ChannelID
}
