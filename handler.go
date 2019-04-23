package main

import (
	"fmt"
	"log"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

const (
	GenerateMessage    = "generate random"
	ForTodayMessage    = "for today"
	ForTomorrowMessage = "for tomorrow"
)

type UpdateHandler interface {
	Handle(update tgbotapi.Update)
}

func NewMessageHandler(kgc *KudaGoClient, b *tgbotapi.BotAPI) UpdateHandler {
	return &handler{
		kgc: *kgc,
		bot: b,
	}
}

type handler struct {
	kgc KudaGoClient
	bot *tgbotapi.BotAPI
}

func (h *handler) Handle(update tgbotapi.Update) {
	log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

	var m string
	switch update.Message.Text {
	case "/start":
		m = "Please generate new event"
	case GenerateMessage:
		event, err := h.kgc.GetRandomEvent("msk", update.Message.From.UserName)
		if err != nil {
			log.Print(err)
			return
		}
		m = fmt.Sprintf("%s \n %s", event.Title, event.Site)
	case ForTodayMessage:
		event, err := h.kgc.GetEventForToday("msk", update.Message.From.UserName)
		if err != nil {
			log.Print(err)
			return
		}
		m = fmt.Sprintf("Event for today \n %s \n %s", event.Title, event.Site)
	case ForTomorrowMessage:
		event, err := h.kgc.GetEventForTomorrow("msk", update.Message.From.UserName)
		if err != nil {
			log.Print(err)
			return
		}
		m = fmt.Sprintf("Event for tomorrow \n %s \n %s", event.Title, event.Site)
	case "/help":
		m = "Generate random event"
	default:
		m = "Fuck off -_-"
	}
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, m)
	msg.ReplyToMessageID = update.Message.MessageID

	attachKeybord(&msg)

	_, err := h.bot.Send(msg)
	if err != nil {
		log.Print(err)
		return
	}
}

func attachKeybord(msg *tgbotapi.MessageConfig) {
	msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
		[]tgbotapi.KeyboardButton{
			tgbotapi.NewKeyboardButton(GenerateMessage),
		},
		[]tgbotapi.KeyboardButton{
			tgbotapi.NewKeyboardButton(ForTodayMessage),
			tgbotapi.NewKeyboardButton(ForTomorrowMessage),
		},
	)
}
