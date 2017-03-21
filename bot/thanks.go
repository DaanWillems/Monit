package bot

import (
    "github.com/bwmarrin/discordgo"
)
func thanks(s *discordgo.Session, m *discordgo.MessageCreate, args []string) string {
    return "You're welcome!"
}
