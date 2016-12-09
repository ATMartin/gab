package main

import (
  "fmt"
  "os"
  "strings"
  "time"

  // "github.com/gorilla/websocket"
)

func makeResponse(messageText string) (responseText string) {
  responseText = "Hey there!"
  return
}

func main() {
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
    if strings.Contains(message, botId) {
      go func() {
        fmt.Println("Derp was mentioned!")
        resp := []byte("You talkin' to me?")
        err = conn.WriteMessage(1, resp)
        if err != nil {
          fmt.Println("Error writing message:", string(resp))
          return
        }
      }
    }
  }
}
