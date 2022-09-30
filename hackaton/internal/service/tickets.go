package service

import (
	"fmt"
	"hackaton/internal/customError"
)

type Bookings interface {
	// Create create a new Ticket
	Create(t Ticket) (Ticket, error)
	// Read read a Ticket by id
	Read(id int) (Ticket, error)
	// Update update values of a Ticket
	Update(id int, t *Ticket) (Ticket, error)
	// Delete delete a Ticket by id
	Delete(id int) (int, error)
}

type bookings struct {
	Tickets []Ticket
}

type Ticket struct {
	Id                              int
	Names, Email, Destination, Date string
	Price                           int
}

// NewBookings creates a new bookings service
func NewBookings(Tickets []Ticket) Bookings {
	return &bookings{Tickets: Tickets}
}

func (b *bookings) Create(t Ticket) (Ticket, error) {
	newError := &customError.CustomError{}
	for _, dd := range b.Tickets {
		if dd.Id == t.Id {
			return t, newError.BadRequest(fmt.Errorf("id already exist"))
		}
	}
	b.Tickets = append(b.Tickets, t)
	return t, nil
}

func (b *bookings) Read(id int) (Ticket, error) {
	newError := &customError.CustomError{}

	for _, t := range b.Tickets {
		if t.Id == id {
			return t, nil
		}
	}
	return Ticket{}, newError.BadRequest(fmt.Errorf("ticket not exist"))
}

func (b *bookings) Update(id int, t *Ticket) (Ticket, error) {
	newError := &customError.CustomError{}
	index := 0
	for indexOFTicket, ticket := range b.Tickets {
		if id == ticket.Id {
			index = indexOFTicket
		}
	}
	if index == 0 {
		return Ticket{}, newError.BadRequest(fmt.Errorf("ticket not exist"))
	}
	b.Tickets[index] = *t
	b.Tickets[index].Id = id
	return b.Tickets[index], nil
}

func (b *bookings) Delete(id int) (int, error) {
	newError := &customError.CustomError{}
	var idExist bool
	var Array1 []Ticket
	for index, t := range b.Tickets {
		if t.Id == id {
			idExist = true
			Array1 = b.Tickets[:index]
			Array1 = append(Array1, b.Tickets[index+1:]...)
		}
	}
	if idExist == false {
		return 0, newError.BadRequest(fmt.Errorf("Id not exist"))
	}
	b.Tickets = Array1
	fmt.Println(b.Tickets)
	return 1, nil
}
