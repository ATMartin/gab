package main

import (
  // Native packages
  "bytes"
  "encoding/json"
  "fmt"
  "io/ioutil"
  "net/http"

  // External packages
  "github.com/gorilla/websocket"
  "github.com/fatih/structs"
)

// ** GLOBALS **
var slackApiToken string

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
  Channel string
}

type slackRtmAttachment struct {
  Color string  `json:"color"`
  Title string  `json:"title"`
  Text string   `json:"text"`
}

type slackRtmResponse struct {
  Id int                           `json:"id"`
  Type string                      `json:"type"`
  Channel string                   `json:"channel"`
  Text string                      `json:"text"`
  Attachments []slackRtmAttachment `json:"attachments"`
}

type  slackApiResponse struct {
  Channel string                   `json:"channel"`
  Text string                      `json:"text"`
  Attachments []slackRtmAttachment `json:"attachments"`
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

  // We've gotten good data back - stash our API token
  // for future requests that use it!
  slackApiToken = apiToken
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

func slackGetMessage(messageSource []byte) (message slackRtmEvent, err error) {
  err = json.Unmarshal(messageSource, &message)
  return
}

func slackPostMessage(channel string, content map[string]interface{}) (err error) {
  endpoint := fmt.Sprintf("https://slack.com/api/chat.postMessage?token=%s", slackApiToken)
  fmt.Println(string(content))
  resp, err := http.PostForm(endpoint, bytes.NewBuffer(content))
  if err != nil {
    fmt.Println("Error posting message to Slack!")
    return
  }

  body, err := ioutil.ReadAll(resp.Body)
  fmt.Println(string(body))

  return
}
