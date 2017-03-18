package main

func main() {
    loadConfig()

    bot := new(Bot)
    bot.Connect()

    <-make(chan struct{})
    return
}

