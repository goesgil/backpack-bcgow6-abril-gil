package db

import (
	"fmt"
	"os"
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
	return DB
}

func Writer(db []Transaction) error {
	file, err := os.OpenFile("./trxs.json", os.O_RDWR|os.O_CREATE, 777)
	if err != nil {
		return err
	}
	return transformTransactionToJson(db, file)
}

func GenerateMock() error {
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

	file, err := os.Create("./trxs.json")
	if err != nil {
		return err
	}
	return transformTransactionToJson(DB, file)
}

func transformTransactionToJson(t []Transaction, file *os.File) error {
	_, err := file.Write([]byte(fmt.Sprintf("[\n")))
	if err != nil {
		return err
	}
	for i, v := range t {
		if i < len(t)-1 {
			_, err = file.Write([]byte(fmt.Sprintf("\t{\n\t\t\"id\": %d,\n \t\t\"codeTrxs\": \"%s\",\n \t\t\"coin\": \"%s\",\n \t\t\"amount\": %f,\n \t\t\"transmitter\": \"%s\",\n \t\t\"created_at\": \"%s\"\n\t},\n", v.Id, v.CodeTrxs, v.Coin, v.Amount, v.Transmitter, v.Created_at)))
			if err != nil {
				return err
			}
		} else {
			_, err = file.Write([]byte(fmt.Sprintf("\t{\n\t\t\"id\": %d,\n \t\t\"codeTrxs\": \"%s\",\n \t\t\"coin\": \"%s\",\n \t\t\"amount\": %f,\n \t\t\"transmitter\": \"%s\",\n \t\t\"created_at\": \"%s\"\n\t}\n", v.Id, v.CodeTrxs, v.Coin, v.Amount, v.Transmitter, v.Created_at)))
			if err != nil {
				return err
			}
		}
	}
	_, err = file.Write([]byte(fmt.Sprintf("]")))
	if err != nil {
		return err
	}
	return nil
}
