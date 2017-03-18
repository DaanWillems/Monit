package config

type config struct {
    Token string
    Mainchannel string
    Text string
}

var c config

func LoadConfig() {
    parseConfig(&c)
}

func initConfig() config {
    co := config{Token: "0", Mainchannel: "0", Text: "test"}
    return co
}

func GetToken() string {
    return c.Token
}

func GetMainChannel() string {
    return c.Mainchannel
}

func GetText(s string) string {
    return c.Text
}
