package main

import (
	tb "gopkg.in/tucnak/telebot.v2"
	"os"
	"time"
)

var b *tb.Bot

func main() {
	var err error

	b, err = tb.NewBot(tb.Settings{
		Token:  os.Getenv("TG_TOKEN"),
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		panic(err)
	}

	b.Handle(tb.OnUserJoined, func(m *tb.Message) {
		b.Delete(m)
		if m.UserJoined.ID == m.Sender.ID {
			b.Ban(m.Chat, &tb.ChatMember{
				User: m.Sender,
			})
			b.Unban(m.Chat, m.Sender)
		}
	})

	// WHAT THE FUCK, TUCNAK?!
	b.Handle(tb.OnText, deleteMessage)
	b.Handle(tb.OnPhoto, deleteMessage)
	b.Handle(tb.OnAudio, deleteMessage)
	b.Handle(tb.OnAnimation, deleteMessage)
	b.Handle(tb.OnDocument, deleteMessage)
	b.Handle(tb.OnSticker, deleteMessage)
	b.Handle(tb.OnVideo, deleteMessage)
	b.Handle(tb.OnVoice, deleteMessage)
	b.Handle(tb.OnVideoNote, deleteMessage)
	b.Handle(tb.OnContact, deleteMessage)
	b.Handle(tb.OnLocation, deleteMessage)
	b.Handle(tb.OnVenue, deleteMessage)
	b.Handle(tb.OnDice, deleteMessage)

	b.Handle(tb.OnUserLeft, func(m *tb.Message) {
		b.Delete(m)
	})

	b.Start()
}

func deleteMessage(m *tb.Message) {
	if m.Sender.ID == 777000 {
		return
	}

	if m.ReplyTo == nil {
		res, err := b.ChatMemberOf(m.Chat, m.Sender)
		if err == nil && res.Role != tb.Creator && res.Role != tb.Administrator {
			b.Delete(m)
		}
	}
}
