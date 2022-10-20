package db

import (
	"fmt"
	"os"
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
	file, err := os.OpenFile("./trxs.json", os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		return err
	}
	return transformTransactionToJson(db, file)
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
