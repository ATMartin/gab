package main

import (
  "fmt"
  "os"
)

func main() {
  conn, err := slackInit(os.Getenv("SLACKBOT_TOKEN"))
  if err != nil {
    fmt.Println("Listen harder! Websocket connection failed.")
    return
  }

  fmt.Println("Now you're cooking! Kill process to exit.")

  for {
    _, message, err := conn.ReadMessage()
    if err != nil { return }
    fmt.Println(string(message))
  }
}
