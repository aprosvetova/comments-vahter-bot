package main

import (
	tb "gopkg.in/tucnak/telebot.v2"
	"os"
	"time"
)

func main() {
	b, err := tb.NewBot(tb.Settings{
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

	b.Handle(tb.OnText, func(m *tb.Message) {
		if m.ReplyTo == nil {
			res, err := b.ChatMemberOf(m.Chat, m.Sender)
			if err == nil && res.Role != tb.Creator && res.Role != tb.Administrator {
				b.Delete(m)
			}
		}
	})

	b.Handle(tb.OnUserLeft, func(m *tb.Message) {
		b.Delete(m)
	})

	b.Start()
}
