package main

type Ticket struct {
	Code string ""
	ReadAt string ""
}

type TicketRepository struct {
	Tickets map[string]Ticket
}

func main() {
	ticket, repo := &Ticket{}, &TicketRepository{ make(map[string]Ticket) }
	Routes(ticket, repo)
}
