package main

import (
	"flag"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	channelName   = flag.String("channel", "anders14_", "The main channel to connect to")
	commandPrefix = flag.String("prefix", "!", "What prefix to put in front of commands")
)

func main() {
	flag.Parse()

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	botName := os.Getenv("BOT_NAME")
	botOauth := os.Getenv("BOT_OAUTH")

	client := NewClient(botName, botOauth, "#"+*channelName)

	client.Connect()
	defer client.Close()
	client.Login()

	client.HandleChat()
}
