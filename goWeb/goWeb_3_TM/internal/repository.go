package internal

import (
	"errors"
	"fmt"
	"time"
)

var (
	trxs []Transaction
)

type Transaction struct {
	Id          int     `json:"id"`
	CodeTrxs    string  `json:"codeTrxs"`
	Coin        string  `json:"coin"`
	Amount      float64 `json:"amount"`
	Transmitter string  `json:"transmitter"`
	Created_at  string  `json:"created_at"`
}

func GenerateMock() {
	coins := []string{"BTC", "ETH", "USDT", "BCH", "LTC", "DOT", "BNB"}
	amounts := []float64{1500.89, 2000.00, 3000.00, 4000.99, 5000.00, 6000.35, 7000.55, 8000.00, 9000.00, 10000.00}

	for i := 0; i < 7; i++ {
		trxs = append(trxs, Transaction{
			Id:          i + 1,
			CodeTrxs:    fmt.Sprintf("code-trxs-%d", i+1),
			Coin:        coins[i],
			Amount:      amounts[i],
			Transmitter: fmt.Sprintf("transmitter-uuid-user-%d", i),
			Created_at:  time.Now().Format("2006-11-02 15:04:05"),
		})
	}
}

type Trxs interface {
	Put(Amount float64, id int, CodeTrxs, Coin, Transmitter string) error
	Patch(Amount float64, id int, CodeTrxs string) error
	Delete(id int) error
}

type repositoryTrxs struct{}

func NewRepository() Trxs {
	return &repositoryTrxs{}
}

func (r *repositoryTrxs) Put(Amount float64, id int, CodeTrxs, Coin, Transmitter string) error {
	if len(trxs) == 0 {
		return errors.New("Data not found")
	}
	findIndex := FindIndexById(id, trxs)
	if findIndex == -1 {
		return errors.New("Data not found")
	}
	trxs[findIndex].Amount = Amount
	trxs[findIndex].CodeTrxs = CodeTrxs
	trxs[findIndex].Coin = Coin
	trxs[findIndex].Transmitter = Transmitter
	fmt.Println(trxs[findIndex], "INDEX")
	fmt.Println(trxs)
	return nil
}

func (r *repositoryTrxs) Patch(Amount float64, id int, CodeTrxs string) error {
	if len(trxs) == 0 {
		return errors.New("Data not found")
	}
	findIndex := FindIndexById(id, trxs)
	if findIndex == -1 {
		return errors.New("Data not found")
	}
	trxs[findIndex].Amount = Amount
	trxs[findIndex].CodeTrxs = CodeTrxs
	fmt.Println(trxs[findIndex], "INDEX")
	fmt.Println(trxs)
	return nil
}

func (r *repositoryTrxs) Delete(id int) error {
	if len(trxs) == 0 {
		return errors.New("Data not found")
	}
	findIndex := FindIndexById(id, trxs)
	if findIndex == -1 {
		return errors.New("Data not found")
	}
	trxs = append(trxs[:findIndex], trxs[findIndex+1:]...)
	fmt.Println(trxs)
	return nil
}

func FindIndexById(id int, slice []Transaction) int {
	for i, v := range slice {
		if v.Id == id {
			return i
		}
	}
	return -1
}
