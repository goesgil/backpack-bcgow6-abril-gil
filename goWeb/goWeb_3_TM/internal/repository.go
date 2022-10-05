package internal

import (
	"errors"
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

func FindIndexById(id int, slice []Transaction) int {
	for i, v := range slice {
		if v.Id == id {
			return i
		}
	}
	return -1
}
