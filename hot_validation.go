package main

import "fmt"

func (repository *TicketRepository) Validate(ticket *Ticket) (string) {
	if len(ticket.Number) < 12 {
		fmt.Println("ERROR")
		return "ERROR"
	} else if len(ticket.Number) > 18 {
		fmt.Println("INVALID")
		return "INVALID"
	} else if _, ok := repository.Tickets[ticket.Number] ; ok {
		fmt.Println("VALIDATED")
		return "VALIDATED"
	}
	fmt.Println(ticket.Number)
	repository.Tickets[ticket.Number] = ticket.Number
	return "SUCCESS"
}
