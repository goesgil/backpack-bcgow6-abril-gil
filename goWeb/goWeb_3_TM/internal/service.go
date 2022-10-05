package internal

type Service interface {
	Put(Amount float64, Id int, CodeTrxs, Coin, Transmitter string) error
	Patch(Amount float64, Id int, CodeTrxs string) error
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

func (s *service) Delete(Id int) error {
	err := s.repository.Delete(Id)
	if err != nil {
		return err
	}
	return nil
}
