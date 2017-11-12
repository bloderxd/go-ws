package main

import (
	"net"
	"log"
	"os"
)

type Ticket struct {
	Code string ""
	ReadAt string ""
}

type TicketRepository struct {
	Tickets map[string]Ticket
}


func ip() (string, bool) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return err.Error(), false
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	if localAddr != nil {
		return localAddr.IP.String(), true
	}
	return "ERROR", false
}

func infoKey() {
	if key, success := ip(); success {
		log.New(os.Stdout, "[LOCAL-SERVER] KEY: ", 0).Println(key)
	} else {
		log.New(os.Stderr, "[LOCAL-SERVER] ERROR: ", log.Ldate|log.Ltime|log.Lshortfile).Println(key)
	}
}

func main() {
	infoKey()
	ticket, repo := &Ticket{}, &TicketRepository{ make(map[string]Ticket) }
	repo.PersistentTickets()
	Routes(ticket, repo)
}
