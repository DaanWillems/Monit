package main

import (
    "os"
    "io/ioutil"
    "encoding/json"
)

const ConfigLocation string = "/etc/opt/monit/"
const ConfigName string = "config.json"
const Template string = "{\n \"mainchannel\": \"0\",\n \"token\": \"0\" \n}"

var data map[string]interface{}

func loadConfig() {
    if _, err := os.Stat(ConfigLocation); os.IsNotExist(err) {
        os.Mkdir(ConfigLocation, os.ModePerm)
    }

    var _, err = os.Stat(ConfigLocation+ConfigName)

    if os.IsNotExist(err) {
        var file, _ = os.Create(ConfigLocation+ConfigName)
        _, _ = file.WriteString(Template)
        defer file.Close()
    }

    dat, _ := ioutil.ReadFile(ConfigLocation+ConfigName)
    _ = json.Unmarshal(dat, &data)
}

func getToken() string {
    return data["token"].(string)
}

func getMainChannel() string {
    return data["mainchannel"].(string)
}
