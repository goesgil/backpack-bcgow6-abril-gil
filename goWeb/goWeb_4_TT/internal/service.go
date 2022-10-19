package internal

import "github.com/goesgil/backpack-bcgow6-abril-gil/goWeb/goWeb_4_TT/pkg/db"

type Service interface {
	Create(Amount float64, CodeTrxs, Coin, Transmitter string) (db.Transaction, error)
	Put(Amount float64, Id int, CodeTrxs, Coin, Transmitter string) error
	Patch(Amount float64, Id int, CodeTrxs string) error
	GetAll() []db.Transaction
	Delete(Id int) error
}

type service struct {
	repository Trxs
}

func NewService(r Trxs) Service {
	return &service{
		repository: r,
	}
}

func (s *service) Create(Amount float64, CodeTrxs, Coin, Transmitter string) (db.Transaction, error) {
	Id := s.repository.GetLastId() + 1
	trxs, err := s.repository.Create(Amount, Id, CodeTrxs, Coin, Transmitter)
	if err != nil {
		return db.Transaction{}, err
	}
	return trxs, nil
}

func (s *service) Put(Amount float64, Id int, CodeTrxs, Coin, Transmitter string) error {
	err := s.repository.Put(Amount, Id, CodeTrxs, Coin, Transmitter)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) Patch(Amount float64, Id int, CodeTrxs string) error {
	err := s.repository.Patch(Amount, Id, CodeTrxs)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) GetAll() []db.Transaction {
	return s.repository.GetAll()
}

func (s *service) Delete(Id int) error {
	err := s.repository.Delete(Id)
	if err != nil {
		return err
	}
	return nil
}
