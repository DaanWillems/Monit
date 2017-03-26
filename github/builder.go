package github

import (
    "monit/bot"
    "time"
    "fmt"
)

func build() {
    fmt.Println("Building or some shit")
    msg := bot.SendStandardMessage("Change in development detected.")
    time.Sleep(10000)
    msg = bot.EditMessage(msg.ChannelID, msg.ID, msg.Content+"\nDownloading...")
    time.Sleep(10000)
    msg = bot.EditMessage(msg.ChannelID, msg.ID, msg.Content+"\nTesting...")
    time.Sleep(10000)
    msg = bot.EditMessage(msg.ChannelID, msg.ID, msg.Content+"\nDeploying...")
    time.Sleep(10000)
    bot.EditMessage(msg.ChannelID, msg.ID, msg.Content+"\nBuild and deployment succesful....")
}

func deploy() {

}

