package main

import (
    "monit/bot"
    "monit/config"
)

func main() {
    config.LoadConfig()

    bot := new(bot.Bot)
    bot.Connect()

    <-make(chan struct{})
    return
}

