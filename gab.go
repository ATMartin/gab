package main

import (
  "encoding/json"
  "fmt"
  "os"
  "strings"
  //"time"

  // "github.com/gorilla/websocket"
)

func makeResponse(messageText string) (responseText string) {
  responseText = "Hey there!"
  return
}

func main() {
  counter := 0
  conn, botId, err := slackInit(os.Getenv("SLACKBOT_TOKEN"))
  if err != nil {
    fmt.Println("Listen harder! Websocket connection failed.")
    return
  }

  fmt.Println("Now you're cooking! Kill process to exit.")

  defer conn.Close()
  for {
    _, event, err := conn.ReadMessage()
    if err != nil {
      fmt.Println("Error processing message:", err)
      return
    }

    fmt.Println(string(event))

    message, err := slackGetMessage(event)
    if strings.Contains(message.Text, botId) {
      fmt.Println("Derp was mentioned!")
      respObj := slackRtmResponse{
        Id: counter,
        Type: "message",
        Channel: message.Channel,
        Text: "You talkin' to me?",
      }
      pkg, err := json.Marshal(respObj)
      fmt.Println(string(pkg))
      if err != nil {
        fmt.Println("Error marshaling!")
        return
      }
      err = conn.WriteMessage(1, pkg)
      if err != nil {
        fmt.Println("Error writing message!")
        return
      }
    }
  }
}
