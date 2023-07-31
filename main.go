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

	menu := &tele.ReplyMarkup{}
	btnSummaryDaily := menu.Data("More", "daily")

	b.Handle("/today", func(c tele.Context) error {
		sum, err := repo.GetTodaySum()
		if err != nil {
			log.Fatal(err)
		}

		message := fmt.Sprintf("Total: %d", sum)
		menu.Inline(
			menu.Row(btnSummaryDaily),
		)

		return c.Send(message, menu)
	}, CheckAccess)

	b.Handle(&btnSummaryDaily, func(c tele.Context) error {
		transactions, err := repo.GetTodayList()
		if err != nil {
			log.Fatal(err)
		}

		messageTitle := "Today's transactions"
		message := fmt.Sprintf(
			"%s: \n\n%v",
			messageTitle,
			formatTransactions(transactions),
		)

		return c.Send(message)
	}, CheckAccess)

	b.Handle("/week", func(c tele.Context) error {
		sum, err := repo.GetWeekSum()
		if err != nil {
			log.Fatal(err)
		}

		messageTitle := "Weeks's transactions"
		message := fmt.Sprintf(
			"%s: \n\nTotal: %d",
			messageTitle,
			sum,
		)

		return c.Send(message)
	}, CheckAccess)

	b.Handle("/month", func(c tele.Context) error {
		sum, err := repo.GetMonthSum()
		if err != nil {
			log.Fatal(err)
		}

		messageTitle := "Month's transactions"
		message := fmt.Sprintf(
			"%s: \n\nTotal: %d",
			messageTitle,
			sum,
		)

		return c.Send(message)
	}, CheckAccess)

	b.Handle(tele.OnText, saveTransaction, CheckAccess, CheckChangeDateRequest, CheckFormat)

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
		Sum:       sum,
		Desc:      found[0][2],
		Date:      time.Now().Format("2006-01-02"),
		MessageId: uint(c.Message().ID),
	}
	repo.SaveTransaction(t)

	return c.Send("Save!")
}

func changeDate(c tele.Context) error {
	messageId := uint(c.Message().ReplyTo.ID)

	dateRegex := regexp.MustCompile(`(\d{2}\.\d{2})`)
	dateFound := dateRegex.FindAllString(c.Text(), -1)

	date := time.Now()
	if dateFound != nil {
		year := date.Year()
		date, _ = time.Parse("02.01", dateFound[0])
		date = date.AddDate(year, 0, 0)
	}

	transactionExists, err := repo.UpdateDateByMessageId(messageId, date)
	if !transactionExists {
		return c.Send("Transaction not found")
	}

	if err != nil {
		log.Fatal(err)
	}

	return c.Send("Date updated!")
}

func CheckChangeDateRequest(next tele.HandlerFunc) tele.HandlerFunc {
	return func(c tele.Context) error {
		if c.Message().ReplyTo != nil {
			found, err := regexp.MatchString(`^\d{2}\.\d{2}$`, c.Text())
			if err != nil {
				log.Println(err)
			}

			if !found {
				c.Send("Unrecognized format of date")
				return nil
			}

			return changeDate(c)
		}
		return next(c)
	}
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

func CheckAccess(next tele.HandlerFunc) tele.HandlerFunc {
	return func(c tele.Context) error {
		allowedUserId := os.Getenv("USER_ID")
		userId := int(c.Message().Chat.ID)
		if fmt.Sprintf("%d", userId) != allowedUserId {
			c.Send("Access denied!")
			return nil
		}

		return next(c)
	}
}

func formatTransactions(transactions []repo.Transaction) string {
	if len(transactions) == 0 {
		return "no transactions"
	}

	fmtTransactions := make([]string, len(transactions))
	for i, t := range transactions {
		fmtTransactions[i] = fmt.Sprintf("%d %s", t.Sum, t.Desc)
	}

	return strings.Join(fmtTransactions, "\n")
}
