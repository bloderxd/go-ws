package main


type Ticket struct {
	Number string ""
}

type TicketRepository struct {
	Tickets map[string]string
}

func main() {
	ticket, repo := &Ticket{}, &TicketRepository{ make(map[string]string) }
	Routes(ticket, repo)
}
