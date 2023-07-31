package repository

import (
	"time"
)

const DB_DATE_FORMAT = "2006-01-02"

type Transaction struct {
	ID        uint `gorm:"primarykey"`
	UserId    uint
	Date      string
	Sum       int
	Desc      string
	MessageId uint
}

func GetTodayList(userId uint) ([]Transaction, error) {
	var transactions []Transaction
	result := db.Find(
		&transactions,
		"date = ? AND user_id = ?",
		time.Now().Format(DB_DATE_FORMAT),
		userId,
	)

	return transactions, result.Error
}

func GetTodaySum(userId uint) (int, error) {
	var sum int
	result := db.Table("transactions").Select("sum(sum)").Where(
		"date = ? AND user_id = ?",
		time.Now().Format(DB_DATE_FORMAT),
		userId,
	)
	result.Row().Scan(&sum)

	return sum, result.Error
}

func GetWeekList(userId uint) ([]Transaction, error) {
	var transactions []Transaction
	today := time.Now()
	startOfWeek := getStartOfWeek(today)

	result := db.Where(
		"date >= ? AND date <= ? AND user_id = ?",
		startOfWeek.Format(DB_DATE_FORMAT),
		today.Format(DB_DATE_FORMAT),
		userId,
	).Find(&transactions)

	return transactions, result.Error
}

func GetWeekSum(userId uint) (int, error) {
	var sum int
	today := time.Now()
	startOfWeek := getStartOfWeek(today)

	result := db.Table("transactions").Select("sum(sum)").Where(
		"date >= ? AND date <= ? AND user_id = ?",
		startOfWeek.Format(DB_DATE_FORMAT),
		today.Format(DB_DATE_FORMAT),
		userId,
	)
	result.Row().Scan(&sum)

	return sum, result.Error
}

func GetMonthList(userId uint) ([]Transaction, error) {
	var transactions []Transaction
	today := time.Now()
	startOfMonth := time.Now().AddDate(0, 0, -today.Day())

	result := db.Where(
		"date > ? AND date <= ? AND user_id = ?",
		startOfMonth.Format(DB_DATE_FORMAT),
		today.Format(DB_DATE_FORMAT),
		userId,
	).Find(&transactions)

	return transactions, result.Error
}

func GetMonthSum(userId uint) (int, error) {
	var sum int
	today := time.Now()
	startOfMonth := time.Now().AddDate(0, 0, -today.Day())

	result := db.Table("transactions").Select("sum(sum)").Where(
		"date > ? AND date <= ? AND user_id = ?",
		startOfMonth.Format(DB_DATE_FORMAT),
		today.Format(DB_DATE_FORMAT),
		userId,
	)
	result.Row().Scan(&sum)

	return sum, result.Error
}

func SaveTransaction(t Transaction) (Transaction, error) {
	result := db.Create(&t)

	return t, result.Error
}

func getStartOfWeek(today time.Time) time.Time {
	weekday := today.Weekday()
	var daysFromWeekStart int
	if weekday == 0 {
		daysFromWeekStart = 7 - 1
	} else {
		daysFromWeekStart = int(weekday) - 1
	}
	return today.AddDate(0, 0, -daysFromWeekStart)
}

func UpdateDateByMessageId(messageId uint, date time.Time, userId uint) (bool, error) {
	newDate := date.Format(DB_DATE_FORMAT)
	result := db.Model(&Transaction{}).Where(
		"message_id = ? AND user_id = ?", 
		messageId,
		userId,
	).Update("date", newDate)
	found := result.RowsAffected > 0

	return found, result.Error
}
