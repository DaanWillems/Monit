package bot

import (
    "os"
    "github.com/bwmarrin/discordgo"
)
func disconnect(s *discordgo.Session, m *discordgo.MessageCreate, args []string) string {
    if(m.Author.Username == "Solawep") {
        SendMessage("Shutting down bot...", m.ChannelID)
        os.Exit(0)
        return "Shutting down bot...!"
    } else {
        return "You are allowed to use this command."
    }
}
