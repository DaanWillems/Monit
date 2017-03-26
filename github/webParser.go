package github

import (
    "io/ioutil"
    "net/http"
    "encoding/json"
)


func decode(r *http.Request) map[string]interface{} {
    bodyEnc, _ := ioutil.ReadAll(r.Body)
    var body map[string]interface{}
    _ = json.Unmarshal(bodyEnc, &body)
    return body
}
