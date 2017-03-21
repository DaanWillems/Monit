package bot

import "github.com/bwmarrin/discordgo"

func source(s *discordgo.Session, m *discordgo.MessageCreate, args []string) string {
    return "https://github.com/DaanWillems/Monit"
}
