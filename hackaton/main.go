package main

import (
	"fmt"
	db "hackaton/internal/dbFile"
	"hackaton/internal/service"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	path := "./tickets.csv"
	newTicket := service.Ticket{
		Id: 1001,
		Price: 500,
		Names: "Abril Gil",
		Date: "23:00",
		Destination: "Alaska",
		Email: "abril@example.com",
	}
	updateTicket := service.Ticket{
		Price: 2500,
		Names: "Abril Gil",
		Date: "00:00",
		Destination: "Alaska",
		Email: "abril@example.com",
	}

	// Funcion para obtener tickets del archivo csv
	fileWithTickets := &db.DB{}
	tickets, err := fileWithTickets.Read(&path)
	if err != nil {
		panic(err)
	}
	bookings := service.NewBookings(tickets)

	ticketCreated, err := bookings.Create(newTicket)
	if err != nil {
		panic(err)
	}
	fmt.Println(ticketCreated,"ticketCreated")

	ticketUpdated, err := bookings.Update(1001, &updateTicket)
	if err != nil {
		panic(err)
	}
	fmt.Println(ticketUpdated,"ticketUpdated")

	ticketGet , err := bookings.Read(1001)
	if err != nil {
		panic(err)
	}
	fmt.Println(ticketGet,"ticketGet")

	ticketDeleted , err := bookings.Delete(6)
	if err != nil {
		panic(err)
	}
	fmt.Println(ticketDeleted,"ticketDeleted") 


	ticketGetted , err := bookings.Read(1001)
	if err != nil {
		panic(err)
	}
	fmt.Println(ticketGetted,"ticketGetted")

}
