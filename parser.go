package main

import (
    "os"
    "io/ioutil"
    "encoding/json"
    "fmt"
)

const ConfigLocation string = "/etc/opt/monit/"
const ConfigName string = "config.json"
const Template string = "{\n \"mainchannel\": \"0\",\n \"token\": \"0\",\n \"text\": {\n   \"example\": \"example\"\n } \n}"

var data map[string]interface{}
var text map[string]interface{}

func loadConfig() {

    var mustConfigure bool = false

    if _, err := os.Stat(ConfigLocation); os.IsNotExist(err) {
        os.Mkdir(ConfigLocation, os.ModePerm)
        mustConfigure = true
    }

    var _, err = os.Stat(ConfigLocation+ConfigName)

    if os.IsNotExist(err) {
        var file, _ = os.Create(ConfigLocation+ConfigName)
        _, _ = file.WriteString(Template)
        defer file.Close()
        mustConfigure = true
    }

    if mustConfigure {
        fmt.Println("Created conf files, please configure your token correctly and then start Monit again.")
        os.Exit(0)
    }

    dat, _ := ioutil.ReadFile(ConfigLocation+ConfigName)
    _ = json.Unmarshal(dat, &data)

    text = data["text"].(map[string]interface{})
}

func getToken() string {
    return data["token"].(string)
}

func getMainChannel() string {
    return data["mainchannel"].(string)
}

func getText(s string) string {
    return text[s].(string)
}
