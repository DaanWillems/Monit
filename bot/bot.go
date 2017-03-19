package bot

import (
    "github.com/bwmarrin/discordgo"
    "fmt"
    "monit/config"
    "os"
)

var Session *discordgo.Session
var botId string
var mainChannel string

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
    //Make sure the bot doesnt respond to itself
    if m.Author.ID == botId {
        return
    }

    //Make sure the message is a command
    if m.Content[0] != '!' {
        return
    }

    r := handleCommand(s, m)
    fmt.Println(m.ChannelID)
    _, _ = s.ChannelMessageSend(m.ChannelID, r)
}

//Send message defaults to mainchannel unless specified otherwise in the constructor
func SendMessage(m string) {
    _, _ = Session.ChannelMessageSend(mainChannel, m)
}

func Connect() {
    Session, _ = discordgo.New("Bot " +  config.GetToken()) 
    mainChannel = config.GetMainChannel()//TODO: Place this in config
    u, err := Session.User("@me")

    if err != nil {
        fmt.Println(err.Error())
        fmt.Println("Your token is most likely misconfigured")
        os.Exit(0)
    }
    botId = u.ID

    Session.Open()
    Session.AddHandler(messageHandler)
}

