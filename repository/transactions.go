package repository

import (
	"time"
)

type Transaction struct {
	ID          uint `gorm:"primarykey"`
	Date        string
	Sum         int
	Description string
}

func GetTodayList() ([]Transaction, error) {
	var transactions []Transaction
	result := db.Find(
		&transactions,
		"date = ?",
		time.Now().Format("2006-01-02"),
	)

	return transactions, result.Error
}

func GetTodaySum() (int, error) {
	var sum int
	result := db.Table("transactions").Select("sum(sum)").Where(
		"date = ?",
		time.Now().Format("2006-01-02"),
	)
	result.Row().Scan(&sum)

	return sum, result.Error
}

func GetWeekList() ([]Transaction, error) {
	var transactions []Transaction
	today := time.Now()
	startOfWeek := getStartOfWeek(today)

	result := db.Where(
		"date >= ? AND date <= ?",
		startOfWeek.Format("2006-01-02"),
		today.Format("2006-01-02"),
	).Find(&transactions)

	return transactions, result.Error
}

func GetWeekSum() (int, error) {
	var sum int
	today := time.Now()
	startOfWeek := getStartOfWeek(today)

	result := db.Table("transactions").Select("sum(sum)").Where(
		"date >= ? AND date <= ?",
		startOfWeek.Format("2006-01-02"),
		today.Format("2006-01-02"),
	)
	result.Row().Scan(&sum)

	return sum, result.Error
}

func GetMonthList() ([]Transaction, error) {
	var transactions []Transaction
	today := time.Now()
	startOfMonth := time.Now().AddDate(0, 0, -today.Day())

	result := db.Where(
		"date > ? AND date <= ?",
		startOfMonth.Format("2006-01-02"),
		today.Format("2006-01-02"),
	).Find(&transactions)

	return transactions, result.Error
}

func GetMonthSum() (int, error) {
	var sum int
	today := time.Now()
	startOfMonth := time.Now().AddDate(0, 0, -today.Day())

	result := db.Table("transactions").Select("sum(sum)").Where(
		"date > ? AND date <= ?",
		startOfMonth.Format("2006-01-02"),
		today.Format("2006-01-02"),
	)
	result.Row().Scan(&sum)

	return sum, result.Error
}

func SaveTransaction(t Transaction) (Transaction, error) {
	result := db.Create(&t)

	return t, result.Error
}

func getStartOfWeek (today time.Time) time.Time {
	weekday := today.Weekday()
	var daysFromWeekStart int
	if weekday == 0 {
		daysFromWeekStart = 7 - 1
	} else {
		daysFromWeekStart = int(weekday) - 1
	}
	return today.AddDate(0, 0, -daysFromWeekStart)
}