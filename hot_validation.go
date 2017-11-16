package main

import (
	"time"
)

func (repository *TicketRepository) Validate(ticket *Ticket) (string, *Ticket) {
	if len(ticket.Code) < 12 {
		return "ERROR", ticket
	} else if len(ticket.Code) > 18 {
		return "INVALID", ticket
	} else if response, ok := ticket.isValidated(repository.Tickets); ok {
		return "VALIDATED", response
	}
	validTicket := Ticket{ ticket.Code, time.Now().Format(time.RFC3339) }
	validTicket.PersistTicket()
	repository.FetchAllTickets()
	return "SUCCESS", &validTicket
}

func (ticket *Ticket) isValidated(tickets []Ticket) (*Ticket, bool) {
	for i := range tickets {
		if ticket.Code == tickets[i].Code {
			return &tickets[i], true
		}
	}
	return nil, false
}
