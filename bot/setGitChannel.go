package bot

import (
    "github.com/bwmarrin/discordgo"
)

func setGitChannel(s *discordgo.Session, m *discordgo.MessageCreate, args []string) string {
    if len(args) == 1 {
        mainChannel = args[0]
        return "Git channel set to: "+args[0]+"."
    }
    mainChannel = m.ChannelID
    return "Git channel set!"
}
