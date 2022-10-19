package internal

import (
	"errors"
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

type Trxs interface {
	Create(Amount float64, id int, CodeTrxs, Coin, Transmitter string) (Transaction, error)
	Update(Amount float64, id int, CodeTrxs, Coin, Transmitter string) error
	GetAll() []Transaction
	Delete(id int) error
	GetLastId() int
}

type repositoryTrxs struct{}

func NewRepository() Trxs {
	return &repositoryTrxs{}
}

func (r *repositoryTrxs) Create(Amount float64, id int, CodeTrxs, Coin, Transmitter string) (Transaction, error) {
	newTrxs := Transaction{
		Id:          id,
		CodeTrxs:    CodeTrxs,
		Coin:        Coin,
		Amount:      Amount,
		Transmitter: Transmitter,
		Created_at:  time.Now().Format("2006-11-02 15:04:05"),
	}
	trxs = append(trxs, newTrxs)
	return Transaction{}, nil
}

func (r *repositoryTrxs) Update(Amount float64, id int, CodeTrxs, Coin, Transmitter string) error {
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
	return nil
}

func (r *repositoryTrxs) GetAll() []Transaction {
	return trxs
}

func (r *repositoryTrxs) GetLastId() int {
	indexMax := 1
	if len(trxs) == 0 {
		return indexMax
	}
	for _, trx := range trxs {
		if trx.Id > indexMax {
			indexMax = trx.Id
		}
	}
	return indexMax + 1
}

func FindIndexById(id int, slice []Transaction) int {
	for i, v := range slice {
		if v.Id == id {
			return i
		}
	}
	return -1
}