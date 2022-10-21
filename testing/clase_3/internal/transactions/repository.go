package transactions

import (
	"fmt"

	"github.com/goesgil/backpack-bcgow6-abril-gil/testing/clase_3/pkg/store"
)

type Transaction struct {
	Id          int     `json:"id"`
	CodeTrxs    string  `json:"codeTrxs"`
	Coin        string  `json:"coin"`
	Amount      float64 `json:"amount"`
	Transmitter string  `json:"transmitter"`
	Created_at  string  `json:"created_at"`
}

type Repository interface {
	GetAll() ([]*Transaction, error)
	UpdateName(id int, codeTrxs string) (*Transaction, error)
	Update(id int, amount float64, codeTrxs, coin, transmitter string) (*Transaction, error)
	Delete(id int) error
}

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll() ([]*Transaction, error) {
	var trxs []*Transaction
	if err := r.db.Read(&trxs); err != nil {
		return nil, err
	}

	return trxs, nil
}

func (r *repository) UpdateName(id int, codeTrxs string) (*Transaction, error) {
	var updated bool = false
	var trxs []*Transaction
	if err := r.db.Read(&trxs); err != nil {
		return nil, err
	}

	var trx *Transaction
	for _, value := range trxs {
		if value.Id == id {
			value.CodeTrxs = codeTrxs
			trx = value
			updated = true
		}
	}

	if !updated {
		return nil, fmt.Errorf("transacción id %d no encontrado", id)
	}

	if err := r.db.Write(&trxs); err != nil {
		return nil, err
	}

	return trx, nil
}

func (r *repository) Update(id int, amount float64, codeTrxs, coin, transmitter string) (*Transaction, error) {
	update := false
	newTrxs := &Transaction{
		Id:          id,
		CodeTrxs:    codeTrxs,
		Amount:      amount,
		Coin:        coin,
		Transmitter: transmitter,
	}

	var trxs []*Transaction
	if err := r.db.Read(&trxs); err != nil {
		return nil, err
	}

	for value := range trxs {
		if trxs[value].Id == id {
			trxs[value] = newTrxs
			update = true
		}
	}

	if !update {
		return nil, fmt.Errorf("transacción id %d no encontrada", id)
	}

	if err := r.db.Write(&trxs); err != nil {
		return nil, err
	}

	return newTrxs, nil
}

func (r *repository) Delete(id int) error {
	deleted := false
	var indice int

	var trxs []*Transaction
	if err := r.db.Read(&trxs); err != nil {
		return err
	}
	for value := range trxs {
		if trxs[value].Id == id {
			indice = value
			deleted = true
		}
	}

	if !deleted {
		return fmt.Errorf("la transacción id %d no existe", id)
	}

	trxs = append(trxs[:indice], trxs[indice+1:]...)
	if err := r.db.Write(&trxs); err != nil {
		return err
	}

	return nil
}
