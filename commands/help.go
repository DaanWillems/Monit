package commands 

import "github.com/bwmarrin/discordgo"

func help(s *discordgo.Session, m *discordgo.MessageCreate, args []string) string {
    return "Commands: !help, !ping, !echo, !deploy, !source"
}
