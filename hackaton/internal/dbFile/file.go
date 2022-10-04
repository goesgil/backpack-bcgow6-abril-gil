package db

import (
	"fmt"
	"hackaton/internal/customError"
	"hackaton/internal/service"
	"os"
	"strconv"
	"strings"
)

type DB struct {
	path string
}

func (f *DB) Read(path *string) ([]service.Ticket, error) {
	f.path = *path
	newError := &customError.CustomError{}
	contentByte, err := os.ReadFile(f.path)
	if err != nil {
		return nil, newError.Detail("ERROR_READING_FILE")
	}
	arrayTickets := strings.Split(string(contentByte), "\n")

	dataParsed := []service.Ticket{}
	for i, t := range arrayTickets {
		if i != len(arrayTickets) {
			destructTicket := strings.Split(t, ",")

			id, err := strconv.Atoi(destructTicket[0])
			if err != nil {
				return []service.Ticket{}, newError.Detail("ERROR_PARSED_BY_ID_TICKET")
			}
			price, err := strconv.Atoi(destructTicket[5])
			if err != nil {
				return []service.Ticket{}, newError.Detail("ERROR_PARSED_BY_DATE_TICKET")
			}

			ticketParsed := &service.Ticket{}
			ticketParsed.Id = id
			ticketParsed.Price = price
			ticketParsed.Names = destructTicket[1]
			ticketParsed.Email = destructTicket[2]
			ticketParsed.Destination = destructTicket[3]
			ticketParsed.Date = destructTicket[4]

			dataParsed = append(dataParsed, *ticketParsed)
		}
	}
	return dataParsed, nil
}

func (f *DB) Write(tickets *[]service.Ticket) error {
	newError := &customError.CustomError{}
	file, err := os.OpenFile(f.path, os.O_RDWR|os.O_CREATE, 777)
	if err != nil {
		return newError.Detail("CANT_READING_FILE")
	}
	dataToCSV := ""
	for index, t := range *tickets {
		if index == 0 {
			dataToCSV += fmt.Sprintf("%v,%s,%s,%s,%s,%v", t.Id, t.Names, t.Email, t.Destination, t.Date, t.Price)
		} else {
			dataToCSV += fmt.Sprintf("\n%v,%s,%s,%s,%s,%v", t.Id, t.Names, t.Email, t.Destination, t.Date, t.Price)
		}
	}
	integer, err := file.WriteString(dataToCSV)
	if err != nil {
		return newError.Detail("CANT_WRITE_FILE")
	}
	fmt.Println(integer)
	return nil
}

func (f *DB) Update(ticket service.Ticket) error {
	newError := &customError.CustomError{}
	contentByte, err := os.ReadFile(f.path)
	if err != nil {
		return newError.Detail("ERROR_READING_FILE")
	}
	arrayTickets := strings.Split(string(contentByte), "\n")

	dataParsed := []service.Ticket{}
	for i, t := range arrayTickets {
		if i != len(arrayTickets) {
			destructTicket := strings.Split(t, ",")
			id, err := strconv.Atoi(destructTicket[0])
			if err != nil {
				return newError.Detail("ERROR_PARSED_BY_ID_TICKET")
			}
			fmt.Println(id, "+", ticket.Id)
			if id == ticket.Id {
				fmt.Println(id, "+", ticket.Id)
				ticketParsed := &service.Ticket{}
				price, err := strconv.Atoi(destructTicket[5])
				if err != nil {
					return newError.Detail("ERROR_PARSED_BY_DATE_TICKET")
				}

				ticketParsed.Id = id
				ticketParsed.Price = price
				ticketParsed.Names = destructTicket[1]
				ticketParsed.Email = destructTicket[2]
				ticketParsed.Destination = destructTicket[3]
				ticketParsed.Date = destructTicket[4]
				dataParsed = append(dataParsed, *ticketParsed)
			} else {
				dataParsed = append(dataParsed, ticket)
			}
		}
	}
	f.Write(&dataParsed)
	return nil
}

func (f *DB) Create(ticket service.Ticket) error {
	newError := &customError.CustomError{}
	arrayTickets := f.Read()
	arrayTickets = append(arrayTickets, ticket)
	f.Write(arrayTickets)
	return nil
}

func (f *DB) Delete(ticket service.Ticket) error {
	newError := &customError.CustomError{}
	contentByte, err := os.ReadFile(f.path)
	if err != nil {
		return newError.Detail("ERROR_READING_FILE")
	}
	arrayTickets := strings.Split(string(contentByte), "\n")

	dataParsed := []service.Ticket{}
	for i, t := range arrayTickets {
		if i != len(arrayTickets) {
			destructTicket := strings.Split(t, ",")
			id, err := strconv.Atoi(destructTicket[0])
			if err != nil {
				return newError.Detail("ERROR_PARSED_BY_ID_TICKET")
			}
			if id == ticket.Id {
				ticketParsed := &service.Ticket{}
				price, err := strconv.Atoi(destructTicket[5])
				if err != nil {
					return newError.Detail("ERROR_PARSED_BY_DATE_TICKET")
				}

				ticketParsed.Id = id
				ticketParsed.Price = price
				ticketParsed.Names = destructTicket[1]
				ticketParsed.Email = destructTicket[2]
				ticketParsed.Destination = destructTicket[3]
				ticketParsed.Date = destructTicket[4]
				dataParsed = append(dataParsed, *ticketParsed)
			}
		}
	}
	f.Write(&dataParsed)
	return nil
}
