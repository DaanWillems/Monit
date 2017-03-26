package bot

import (
    "github.com/bwmarrin/discordgo"
    "fmt"
    "os"
    "monit/config"
)

var session *discordgo.Session
var botId string
var mainChannel string

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
    //Make sure the bot doesnt respond to itself
    if m.Author.ID == botId || m.Author.ID == "293750263351738368" {
        return
    }

    if m.ChannelID == "291912433008640000" {
        return;
    }

    //Make sure the message is a command
    if m.Content[0] != '!' {
        return
    }

    //Send the command to the evaluator
    r := EvaluateCommand(s, m)

    //Send the response
    _, _ = session.ChannelMessageSend(m.ChannelID, r)
}

func SendStandardMessage(m string) *discordgo.Message {
    msg, _ := session.ChannelMessageSend(mainChannel, m)
    return msg
}

func SendMessage(m string, channelID string) {
    _, _ = session.ChannelMessageSend(channelID, m)
}

func DeleteMessage(channelID string, messageID string) {
    session.ChannelMessageDelete(channelID, messageID)
}

func EditMessage(channelID string, messageID string, content string) *discordgo.Message {
    msg, _ := session.ChannelMessageEdit(channelID, messageID, content)
    return msg
}
//Connect to discord server
func Connect() {
    session, _ = discordgo.New("Bot " + config.GetToken())
    mainChannel = config.GetMainChannel() //TODO: Place this in config
    u, err := session.User("@me")
    //If the connection fails the token used to connect is most likely incorrect, thus we exit.
    if err != nil {
        fmt.Println(err.Error())
        fmt.Println("Your token is most likely misconfigured")
        os.Exit(0)
    }

    //Pretty clear
    botId = u.ID

    //Open the actual session and add a handler for messages
    session.Open()
    session.AddHandler(messageHandler)
}
