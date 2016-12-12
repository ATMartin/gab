package main

import (
  "encoding/json"
  "fmt"
  "os"
  "strings"
  "github.com/gorilla/websocket"
)

func makeResponse(messageText string) (responseText string) {
  responseText = "Hey there!"
  return
}

func simpleResponse(conn *websocket.Conn, responseData []byte) (err error) {
  err = conn.WriteMessage(1, responseData)
  if err != nil {
    fmt.Println("Error writing message!")
    return
  }
  return
}

func main() {
  //counter := 0
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

      respAttachment := slackRtmAttachment{
        Color: "red",
        Title: "Attached Title",
        Text: "This is some awesome text yo.",
      }

      respObj := slackApiResponse{
        Channel: message.Channel,
        Text: "You talkin' to me?",
        Attachments: []slackRtmAttachment{respAttachment},
      }

/*
      pkg, _ := json.Marshal(respObj)
      if err != nil {
        fmt.Println("Error marshaling!")
        return
      }
      */

      err := slackPostMessage(message.Channel, structs.Map(respObj))
      if err != nil {
        fmt.Println("Error attempting to post message!")
        return
      }


    }
  }
}
