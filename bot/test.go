package bot

import "github.com/bwmarrin/discordgo"

func test(s *discordgo.Session, m *discordgo.MessageCreate, args []string) string {
    c, _ := s.UserChannelCreate(m.Author.ID)
    SendMessage("hoi", c.ID)
    return "Message Send!"
}
