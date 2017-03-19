package config

import (
    "os"
    "io/ioutil"
    "encoding/json"
    "fmt"
)

const ConfigLocation string = "/etc/opt/monit/"
const ConfigName string = "config.json"

func parseConfig(c *config) {
    loadFile(ConfigLocation, ConfigName, c)
}

func loadFile(p string, f string, c *config) {

    var mustConfigure bool = false

    if _, err := os.Stat(p); os.IsNotExist(err) {
        os.Mkdir(p, os.ModePerm)
        mustConfigure = true
    }

    var _, err = os.Stat(p+f)

    if os.IsNotExist(err) {
        var file, _ = os.Create(p+f)
        var co config = initConfig()
        configString, _ := json.Marshal(co)
        fmt.Println(configString)
        _, _ = file.WriteString(string(configString))
        defer file.Close()
        mustConfigure = true
    }

    if mustConfigure {
        fmt.Println("Created conf files, please configure your token correctly and then start Monit again.")
        os.Exit(0)
    }

    dat, _ := ioutil.ReadFile(p+f)

    _ = json.Unmarshal(dat, c)
    fmt.Println("test: "+c.Mainchannel)
}

