package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	if err := godotenv.Load();err != nil {
		fmt.Println("env file not found")
	}
}
func main() {
	token, exists := os.LookupEnv("BOT_TOKEN")
	if exists {
		fmt.Println("%s", token)
	}

	bot,err := tgbotapi.NewBotAPI(token)
	if err != nil{
		fmt.Println("Bot not found ")
		log.Panic(err)
	}
	bot.Debug = true

	update := tgbotapi.NewUpdate(0)
	update.Timeout = 60

	chanUpdates, err := bot.GetUpdatesChan(update)

	for updates := range chanUpdates {
		if updates.Message == nil {
			continue
		}

		if updates.Message.IsCommand() {
			msg := tgbotapi.NewMessage(updates.Message.Chat.ID, "hello")
			switch updates.Message.Command() {
			case "hello":
				msg.Text = "hello"
			default:
				msg.Text = "right command please"
			}
			if _, err :=bot.Send(msg);err != nil {
				log.Panic( err)
			}
		}
	}
	}

