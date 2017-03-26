package main

import (
    "os"
    "bufio"
    "fmt"
    "monit/bot"
    "monit/config"
    "monit/github"
)

func main() {
    config.LoadConfig()
 //   go getInput()
    bot.LoadCommands()
    bot.Connect()
    github.StartServer()
    <-make(chan struct{})

    return
}

func getInput() {
    for {
        reader := bufio.NewReader(os.Stdin)
        text, err := reader.ReadString('\n')
        if err != nil {
            fmt.Println(err)
        }
        fmt.Println(text)
    }
}

