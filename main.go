package main

import (
	"fmt"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func bot() {
	dat, err := os.ReadFile("data.txt")
	if err != nil {
		log.Panic(err)
	}

	token := string(dat)

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		}
	}
}

func main() {
	db := connectDB()
	defer db.close()
	db.addUser(2)
	// db.printUsers()
	fmt.Println("sus")
	// db.createFanta(1, "pippo")
	db.joinFanta(2, "19ae1f3834aa1da711f5488f653cd0")
}
