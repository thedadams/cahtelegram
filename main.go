package main

import (
	"cahbot/secrets"
	"cahbot/tgbotapi"
	"log"
	"time"
)

func main() {
	bot, err := NewCAHBot(secrets.Token)
	if err != nil {
		log.Panic(err)
	}
	defer bot.db_conn.Close()

	// Remove when deployed
	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	err = bot.UpdatesChan(u)
	c := time.Tick(1 * time.Minute)
	go func() {
		for now := range c {
			log.Printf("%v %s\n", now, "This is when we would clean up the database.")
		}
	}()

	for update := range bot.Updates {
		go bot.HandleUpdate(&update)
	}
}
