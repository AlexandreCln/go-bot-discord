package main

import (
    "encoding/json"
    "flag"
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
    "os/signal"
    "strings"
    "syscall"

    "github.com/bwmarrin/discordgo"
)

var (
  Token string
)

func init() {
    flag.StringVar(&Token, "t", "", "Bot Token")
    flag.Parse()
}

func main() {
    dg, err := discordgo.New("Bot " + Token)
    if err != nil {
        fmt.Println("error creating Discord session,", err)
        return
    }

    dg.AddHandler(messageCreate)

    // https://discord.com/developers/docs/topics/gateway#gateway-intents
    dg.Identify.Intents = discordgo.IntentsGuildMessages

    // Open a websocket connection to Discord and begin listening.
    err = dg.Open()
    if err != nil {
        fmt.Println("error opening connection,", err)
        return
    }

    fmt.Println("Bot is now running. Press CTRL-C to exit.")
    // https://pkg.go.dev/os#Signal
    sc := make(chan os.Signal, 1)
    // https://pkg.go.dev/os/signal#Notify
    signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
    <-sc

    // Cleanly close down the Discord session.
    dg.Close()
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

    // Ignore all messages created by the bot itself
    if m.Author.ID == s.State.User.ID {
        return
    }

    if m.Content == "!test" {
      _, err := s.ChannelMessageSend(m.ChannelID, "Hello World")
      if err != nil {
          fmt.Println(err)
      }
    }
}
