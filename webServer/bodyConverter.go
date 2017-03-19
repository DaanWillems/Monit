package webServer

import (
    "io/ioutil"
    "net/http"
    "encoding/json"
)

func decode(r *http.Request) map[string]interface{} {
    //Read body into byte array
    bodyEnc, _ := ioutil.ReadAll(r.Body)

    //Create variable with empty interface to load arbitrary json
    var body map[string]interface{}
    //Decode json, disgard error(for now)
    _ = json.Unmarshal(bodyEnc, &body)

    return body
}

