package main

import (
  // Native packages
  "encoding/json"
  "fmt"
  "io/ioutil"
  "net/http"

  // External packages
  "github.com/gorilla/websocket"
)

// *** STRUCTS ***
type slackWebsocket struct {
  Url string
}

// *** METHODS **
// Get the Realtime Messaging API Websocket URL
func slackGetWebsocket(apiToken string) (socketAddress string, err error) {
  endpoint := fmt.Sprintf("https://slack.com/api/rtm.start?token=%s", apiToken)
  response, err := http.Get(endpoint)
  if err != nil { return }

  body, err := ioutil.ReadAll(response.Body)
  if err != nil { return }

  var websocketData slackWebsocket
  err = json.Unmarshal(body, &websocketData)

  socketAddress = websocketData.Url
  return
}


// Kick off the websocket listener
func slackInit(apiToken string) (conn *websocket.Conn, err error) {
  websocketUrl, err := slackGetWebsocket(apiToken)
  if err != nil { return }

  // Get a live connection with no special headers
  conn, _, err = websocket.DefaultDialer.Dial(websocketUrl, nil)

  // It's cool if there's an error - it'll be in `err`.
  return
}


