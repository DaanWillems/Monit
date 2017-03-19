package main

import (
    "monit/bot"
    "monit/config"
    "monit/webServer"
    "os"
    "bufio"
    "fmt"
)

func main() {
    config.LoadConfig()

    bot.Connect()

    go getInput()

    webServer.StartServer()
 //   <-make(chan struct{})
    return
}

func getInput() {
    for {
        reader := bufio.NewReader(os.Stdin)
        text, err := reader.ReadString('\n')
        if err != nil {
            fmt.Println(err)
        }
        bot.SendMessage(text)
    }
}

