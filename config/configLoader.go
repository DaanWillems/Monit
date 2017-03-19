package config

type config struct {
    Token string
    Mainchannel string
}

var c config

func LoadConfig() {
    parseConfig(&c)
}

func initConfig() config {
    co := config{Token: "0", Mainchannel: "0"}
    return co
}

func GetToken() string {
    return c.Token
}

func GetMainChannel() string {
    return c.Mainchannel
}

