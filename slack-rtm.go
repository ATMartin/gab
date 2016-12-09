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
type slackRtmStart struct {
  Url string
  Self *slackRtmSelf
}

type slackRtmSelf struct {
  Id string
}

type slackRtmEvent struct {
  Type string
  Text string
}

// *** METHODS **
// Get the Realtime Messaging API Websocket URL
func slackGetWebsocket(apiToken string) (socketAddress, botId string, err error) {
  endpoint := fmt.Sprintf("https://slack.com/api/rtm.start?token=%s", apiToken)
  response, err := http.Get(endpoint)
  if err != nil { return }

  body, err := ioutil.ReadAll(response.Body)
  if err != nil { return }

  var websocketData slackRtmStart
  err = json.Unmarshal(body, &websocketData)

  socketAddress = websocketData.Url
  botId = websocketData.Self.Id
  return
}


// Kick off the websocket listener
func slackInit(apiToken string) (conn *websocket.Conn, botId string, err error) {
  websocketUrl, botId, err := slackGetWebsocket(apiToken)
  if err != nil { return }

  // Get a live connection with no special headers
  conn, _, err = websocket.DefaultDialer.Dial(websocketUrl, nil)

  // It's cool if there's an error - it'll be in `err`.
  return
}

func slackGetMessage(messageSource []byte) (message string, err error) {
  message = ""
  var event slackRtmEvent
  err = json.Unmarshal(messageSource, &event)
  if err != nil { return }
  if event.Type == "message" {
    message = event.Text
  }
  return
}
