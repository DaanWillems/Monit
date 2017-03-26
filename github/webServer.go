package github

import (
    "net/http"
 //   "monit/bot"
    "fmt"
)


func handler(w http.ResponseWriter, r *http.Request) {
 //   bot.SendStandardMessage(r.Host +" "+r.URL.Query().Get("cookie")) 
 //   Get the type of event
    var gitType string = r.Header.Get("X-GitHub-Event")

    b := decode(r)

    if gitType == "pull_request" && b["action"].(string) == "opened" {
        go build()
    }
}

func StartServer() {
    fmt.Println("test")
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8380", nil)
}
