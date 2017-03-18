package main

import (
    "github.com/bwmarrin/discordgo"
)

const Token string = getToken()

type Bot struct {
    dg *discordgo.Session
    botId string
    mainChannel string
}

func (b Bot) messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
    //Make sure the bot doesnt respond to itself
    if m.Author.ID == b.botId {
        return
    }

    //Make sure the message is a command
    if m.Content[0] != '!' {
        return
    }

    r := handleCommand(s, m, b)
    _, _ = s.ChannelMessageSend(m.ChannelID, r)
}

func (b Bot) Connect()  {
    b.dg, _ = discordgo.New("Bot " + Token) 
    b.mainChannel = getMainChannel()//TODO: Place this in config
    u, _ := b.dg.User("@me")
    b.botId = u.ID
    b.dg.Open()
    b.dg.AddHandler(b.messageHandler)
}

