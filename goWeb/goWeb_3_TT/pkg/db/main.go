package db

import (
	"fmt"
	"time"
)

var (
	DB []Transaction
)

type Transaction struct {
	Id          int     `json:"id"`
	CodeTrxs    string  `json:"codeTrxs"`
	Coin        string  `json:"coin"`
	Amount      float64 `json:"amount"`
	Transmitter string  `json:"transmitter"`
	Created_at  string  `json:"created_at"`
}

func NewDB() []Transaction {
	return []Transaction{}
}

func GenerateMock() {
	coins := []string{"BTC", "ETH", "USDT", "BCH", "LTC", "DOT", "BNB"}
	amounts := []float64{1500.89, 2000.00, 3000.00, 4000.99, 5000.00, 6000.35, 7000.55, 8000.00, 9000.00, 10000.00}

	for i := 0; i < 7; i++ {
		DB = append(DB, Transaction{
			Id:          i + 1,
			CodeTrxs:    fmt.Sprintf("code-trxs-%d", i+1),
			Coin:        coins[i],
			Amount:      amounts[i],
			Transmitter: fmt.Sprintf("transmitter-uuid-user-%d", i),
			Created_at:  time.Now().Format("2006-11-02 15:04:05"),
		})
	}
}
