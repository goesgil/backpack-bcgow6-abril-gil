package internal

type Service interface {
	Create(Amount float64, CodeTrxs, Coin, Transmitter string) (Transaction, error)
	Update(Amount float64, Id int, CodeTrxs, Coin, Transmitter string) error
	GetAll() []Transaction
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

func (s *service) Create(Amount float64, CodeTrxs, Coin, Transmitter string) (Transaction, error) {
	id := s.repository.GetLastId()

	newTrxs, err := s.repository.Create(Amount, id, CodeTrxs, Coin, Transmitter)
	if err != nil {
		return Transaction{}, err
	}
	return newTrxs, nil
}

func (s *service) Update(Amount float64, Id int, CodeTrxs, Coin, Transmitter string) error {
	err := s.repository.Update(Amount, Id, CodeTrxs, Coin, Transmitter)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) GetAll() []Transaction {
	return s.repository.GetAll()
}

func (s *service) Delete(Id int) error {
	err := s.repository.Delete(Id)
	if err != nil {
		return err
	}
	return nil
}
