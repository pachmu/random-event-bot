package main

import (
	"flag"
	"log"
	"time"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

var configPath = flag.String("config", "", "Path to config file")

func main() {
	flag.Parse()
	config, err := GetConfig(*configPath)
	if err != nil {
		log.Panic(err)
		return
	}
	kclient := NewKudaGoClient("https://kudago.com/public-api/v1.4")
	bot, err := tgbotapi.NewBotAPI(config.Token)
	if err != nil {
		log.Panic(err)
		return
	}
	handler := NewMessageHandler(&kclient, bot)

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Panic(err)
		return
	}

	// Do not handle a large backlog of old messages
	time.Sleep(time.Millisecond * 500)
	updates.Clear()

	for update := range updates {
		if update.Message == nil {
			continue
		}
		go handler.Handle(update)
	}
}
