package main

import (
	"fmt"
	"time"
)

func (repository *TicketRepository) Validate(ticket *Ticket) (string, Ticket) {
	if len(ticket.Code) < 12 {
		return "ERROR", *ticket
	} else if len(ticket.Code) > 18 {
		return "INVALID", *ticket
	} else if _, ok := repository.Tickets[ticket.Code] ; ok {
		return "VALIDATED", repository.Tickets[ticket.Code]
	}
	repository.Tickets[ticket.Code] = Ticket{
		ticket.Code,
		time.Now().Format(time.RFC3339),
	}
	fmt.Println(time.Now().Format(time.RFC3339))
	return "SUCCESS", repository.Tickets[ticket.Code]
}
