package repository

import "time"


type Transaction struct {
	Sum int
	Description string
	Date time.Time
}

var transactions = []Transaction{
	{Sum: -10, Description: "Hoa"},
	{Sum: 10, Description: "Hadd"},
	{Sum: 10, Description: "Uoa"},
	{Sum: +120, Description: "Haw"},
	{Sum: 10, Description: "Opdd"},
	{Sum: -150, Description: "Has"},
}

func GetTodayList() []Transaction {
	return transactions
}

func GetWeekList() []Transaction {
	return transactions
}

func GetMonthList() []Transaction {
	return transactions
}

func SaveTransaction(t Transaction) error {
	transactions = append(transactions, t)
	return nil
}