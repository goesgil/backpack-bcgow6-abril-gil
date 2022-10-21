package transactions

type Service interface {
	GetAll() ([]*Transaction, error)
	UpdateName(id int, codeTrxs string) (*Transaction, error)
	Update(id int, amount float64, codeTrxs, coin, transmitter string) (*Transaction, error)
	Delete(id int) error
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{repo: r}
}

func (s *service) GetAll() ([]*Transaction, error) {
	return s.repo.GetAll()
}

func (s *service) UpdateName(id int, codeTrxs string) (*Transaction, error) {
	return s.repo.UpdateName(id, codeTrxs)
}

func (s *service) Update(id int, amount float64, codeTrxs, coin, transmitter string) (*Transaction, error) {
	return s.repo.Update(id, amount, codeTrxs, coin, transmitter)
}

func (s *service) Delete(id int) error {
	return s.repo.Delete(id)
}
