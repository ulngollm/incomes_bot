package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	tele "gopkg.in/telebot.v3"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
		return
	}
}

type Transaction struct {
	Sum int
	Description string
	Date time.Time
}


func main() {
	botToken := os.Getenv("TOKEN")
	pref := tele.Settings{
		Token: botToken,
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/today", func(c tele.Context) error {
		transactions := []Transaction{}

		return c.Send(fmt.Sprintf("Today's transactions: %v", transactions))
	})

	b.Handle("/week", func(c tele.Context) error {
		transactions := []Transaction{}

		return c.Send(fmt.Sprintf("Weeks's transactions: %v", transactions))
	})

	b.Handle("/month", func(c tele.Context) error {
		transactions := []Transaction{}

		return c.Send(fmt.Sprintf("Month's transactions: %v", transactions))
	})

	b.Handle(tele.OnText, saveTransaction, CheckFormat)

	b.Start()
}

func saveTransaction(c tele.Context) error {
	re := regexp.MustCompile(`^(?P<sum>[+-]?\d+)?\s+(?P<desc>.*)$`)
	found := re.FindAllStringSubmatch(c.Text(), -1)

	if found == nil {
		return nil
	}

	sum, _ := strconv.Atoi(found[0][1])

	t := Transaction{
		Sum: sum,
		Description: found[0][2],
		Date: time.Now(),
	}
	fmt.Printf("%v", t)

	return c.Send("Save!")
}

func CheckFormat(next tele.HandlerFunc) tele.HandlerFunc {
	return func(c tele.Context) error {
		found, err := regexp.MatchString("^(?P<sum>[+-]?\\d+)?\\s+(?P<desc>.*)$", c.Text())

		if err != nil {
			log.Println(err)
		}

		if !found {
			c.Send("Unrecognized format")
			return nil
		}

		return next(c)
	}
}
