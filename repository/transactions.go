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
	result := db.Where(
		"date > ? AND date <= ?",
		time.Now().AddDate(0, 0, -7).Format("2006-01-02"),
		time.Now().Format("2006-01-02"),
	).Find(&transactions)

	return transactions, result.Error
}

func GetWeekSum() (int, error) {
	var sum int
	result := db.Table("transactions").Select("sum(sum)").Where(
		"date > ? AND date <= ?",
		time.Now().AddDate(0, 0, -7).Format("2006-01-02"),
		time.Now().Format("2006-01-02"),
	)
	result.Row().Scan(&sum)

	return sum, result.Error
}

func GetMonthList() ([]Transaction, error) {
	var transactions []Transaction
	result := db.Where(
		"date > ? AND date <= ?",
		time.Now().AddDate(0, -1, 0).Format("2006-01-02"),
		time.Now().Format("2006-01-02"),
	).Find(&transactions)

	return transactions, result.Error
}

func GetMonthSum() (int, error) {
	var sum int
	result := db.Table("transactions").Select("sum(sum)").Where(
		"date > ? AND date <= ?",
		time.Now().AddDate(0, -1, 0).Format("2006-01-02"),
		time.Now().Format("2006-01-02"),
	)
	result.Row().Scan(&sum)

	return sum, result.Error
}

func SaveTransaction(t Transaction) (Transaction, error) {
	result := db.Create(&t)

	return t, result.Error
}
