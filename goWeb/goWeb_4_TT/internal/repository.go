package internal

import (
	"errors"
	"time"

	"github.com/goesgil/backpack-bcgow6-abril-gil/goWeb/goWeb_4_TT/pkg/db"
)

type Trxs interface {
	Create(Amount float64, id int, CodeTrxs, Coin, Transmitter string) (db.Transaction, error)
	Put(Amount float64, id int, CodeTrxs, Coin, Transmitter string) error
	Patch(Amount float64, id int, CodeTrxs string) error
	Delete(id int) error
	GetLastId() int
	GetAll() []db.Transaction
}

type repositoryTrxs struct {
	db []db.Transaction
}

func NewRepository(newDB []db.Transaction) Trxs {
	return &repositoryTrxs{
		db: newDB,
	}
}

func (r *repositoryTrxs) Create(Amount float64, id int, CodeTrxs, Coin, Transmitter string) (db.Transaction, error) {
	newTrxs := db.Transaction{
		Id:          id,
		CodeTrxs:    CodeTrxs,
		Coin:        Coin,
		Amount:      Amount,
		Transmitter: Transmitter,
		Created_at:  time.Now().Format("2006-11-02 15:04:05"),
	}

	r.db = append(r.db, newTrxs)
	err := db.Writer(r.db)
	if err != nil {
		return db.Transaction{}, err
	}
	return newTrxs, nil
}

func (r *repositoryTrxs) Put(Amount float64, id int, CodeTrxs, Coin, Transmitter string) error {
	if len(r.db) == 0 {
		return errors.New("Data not found")
	}
	findIndex := FindIndexById(id, r.db)
	if findIndex == -1 {
		return errors.New("Data not found")
	}
	r.db[findIndex].Amount = Amount
	r.db[findIndex].CodeTrxs = CodeTrxs
	r.db[findIndex].Coin = Coin
	r.db[findIndex].Transmitter = Transmitter
	err := db.Writer(r.db)
	if err != nil {
		return err
	}
	return nil
}

func (r *repositoryTrxs) Patch(Amount float64, id int, CodeTrxs string) error {
	if len(r.db) == 0 {
		return errors.New("Data not found")
	}
	findIndex := FindIndexById(id, r.db)
	if findIndex == -1 {
		return errors.New("Data not found")
	}
	r.db[findIndex].Amount = Amount
	r.db[findIndex].CodeTrxs = CodeTrxs

	err := db.Writer(r.db)
	if err != nil {
		return err
	}
	return nil
}

func (r *repositoryTrxs) GetAll() []db.Transaction {
	return r.db
}

func (r *repositoryTrxs) Delete(id int) error {
	if len(r.db) == 0 {
		return errors.New("Data not found")
	}
	findIndex := FindIndexById(id, r.db)
	if findIndex == -1 {
		return errors.New("Data not found")
	}
	r.db = append(r.db[:findIndex], r.db[findIndex+1:]...)
	err := db.Writer(r.db)
	if err != nil {
		return err
	}
	return nil
}

func (r *repositoryTrxs) GetLastId() int {
	indexMax := 1
	if len(r.db) == 0 {
		return indexMax
	}
	for _, trx := range r.db {
		if trx.Id > indexMax {
			indexMax = trx.Id
		}
	}
	return indexMax
}

func FindIndexById(id int, slice []db.Transaction) int {
	for i, v := range slice {
		if v.Id == id {
			return i
		}
	}
	return -1
}
