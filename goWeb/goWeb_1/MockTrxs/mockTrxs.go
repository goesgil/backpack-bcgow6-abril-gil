package mocktrxs

import (
	"fmt"
	"os"
	"time"
)

type Transaction struct {
	Id          int     `json:"id"`
	CodeTrxs    string  `json:"codeTrxs"`
	Coin        string  `json:"coin"`
	Amount      float64 `json:"amount"`
	Transmitter string  `json:"transmitter"`
	Created_at  string  `json:"created_at"`
}

type Transactions struct {
	Transactions []Transaction
}

func Generate() {
	mockSliceData := []Transaction{}
	coins := []string{"BTC", "ETH", "USDT", "BCH", "LTC", "DOT", "BNB"}
	amounts := []float64{1500.89, 2000.00, 3000.00, 4000.99, 5000.00, 6000.35, 7000.55, 8000.00, 9000.00, 10000.00}

	for i := 0; i < 7; i++ {
		mockSliceData = append(mockSliceData, Transaction{
			Id:          i + 1,
			CodeTrxs:    fmt.Sprintf("code-trxs-%d", i+1),
			Coin:        coins[i],
			Amount:      amounts[i],
			Transmitter: fmt.Sprintf("transmitter-uuid-user-%d", i),
			Created_at:  time.Now().Format("2006-11-02 15:04:05"),
		})
	}
	file, err := os.Create("./goWeb/transactions.json")
	if err != nil {
		fmt.Println(err.Error())
	}

	file.Write([]byte(fmt.Sprintf("{\n\t\"Transactions\": [\n")))
	for i, v := range mockSliceData {
		if i < len(mockSliceData)-1 {
			file.Write([]byte(fmt.Sprintf("\t{\n\t\t\"id\": %d,\n \t\t\"codeTrxs\": \"%s\",\n \t\t\"coin\": \"%s\",\n \t\t\"amount\": %f,\n \t\t\"transmitter\": \"%s\",\n \t\t\"created_at\": \"%s\"\n\t},\n", v.Id, v.CodeTrxs, v.Coin, v.Amount, v.Transmitter, v.Created_at)))
		} else {
			file.Write([]byte(fmt.Sprintf("\t{\n\t\t\"id\": %d,\n \t\t\"codeTrxs\": \"%s\",\n \t\t\"coin\": \"%s\",\n \t\t\"amount\": %f,\n \t\t\"transmitter\": \"%s\",\n \t\t\"created_at\": \"%s\"\n\t}\n", v.Id, v.CodeTrxs, v.Coin, v.Amount, v.Transmitter, v.Created_at)))
		}
	}
	file.Write([]byte(fmt.Sprintf("\t]\n}")))
}
