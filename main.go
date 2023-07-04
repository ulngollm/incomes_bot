package main

import (
	repo "cost-bot/repository"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
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
		transactions := repo.GetTodayList()
		messageTitle := "Today's transactions"
		message := fmt.Sprintf("%s: \n\n%v", messageTitle, formatTransactions(transactions))

		return c.Send(message)
	})

	b.Handle("/week", func(c tele.Context) error {
		transactions := repo.GetWeekList()

		messageTitle := "Weeks's transactions"
		message := fmt.Sprintf("%s: \n\n%v", messageTitle, formatTransactions(transactions))

		return c.Send(message)
	})

	b.Handle("/month", func(c tele.Context) error {
		transactions := repo.GetMonthList()
		messageTitle := "Month's transactions"
		message := fmt.Sprintf("%s: \n\n%v", messageTitle, formatTransactions(transactions))

		return c.Send(message)
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

	t := repo.Transaction{
		Sum: sum,
		Description: found[0][2],
		Date: time.Now(),
	}
	repo.SaveTransaction(t)

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

func formatTransactions(transactions []repo.Transaction) string {
	fmtTransactions := make([]string, len(transactions))
	for i, t := range transactions {
		fmtTransactions[i] = fmt.Sprintf("%d %s", t.Sum, t.Description)
	}

	return strings.Join(fmtTransactions, "\n")
}