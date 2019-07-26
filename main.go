package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
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
	if err != nil{
		fmt.Println("Error getting channel updates")
	}

	for updates := range chanUpdates {
		if updates.Message == nil {
			continue
		}

		if updates.Message.IsCommand() {
			msg := tgbotapi.NewMessage(updates.Message.Chat.ID, "hello")
			switch updates.Message.Command() {
			case "stats":
				cpu, mem, disk := getStatsFromNode()
				cpuStress := fmt.Sprintf("%f" , cpu)
				memLoad := fmt.Sprintf("%f", mem)
				diskCap := strconv.Itoa(disk)
				msg.Text = cpuStress + "% cpu " + memLoad + "% mem " + diskCap + "% disk "
			default:
				msg.Text = "right command please"
			}
			if _, err :=bot.Send(msg);err != nil {
				log.Panic( err)
			}
		}
	}
}

