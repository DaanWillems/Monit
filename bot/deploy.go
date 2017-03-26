package bot

import (
    "github.com/bwmarrin/discordgo"
)
func deploy(s *discordgo.Session, m *discordgo.MessageCreate, args []string) string {
    if(len(args) != 1) {
        return "Incorrect usage of command. Please refer to manual('!help deploy')"
    }
    //os.Exec("/usr/local/scripts/changebranch"+args[1])
    return "Im a bot tralalala and i do bottie things trrallalallala by name is bot"
}
