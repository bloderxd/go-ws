package main

import (
	"database/sql"
	"log"
	"os"
	_ "github.com/mattn/go-sqlite3"
	"time"
)

func (repository *TicketRepository) PersistentTickets() {
	db, err := openConnectionFrom("./ticket.db")
	if weHaveNotProblem(err, "PersistentTickets") {
		repository.fetchTickets(db)
	}
}

func (ticket *Ticket) PersistTicket() {
	db, err := openConnectionFrom("./ticket.db")
	if weHaveNotProblem(err) {
		prepare, err := db.Prepare("INSERT INTO ticket(code, readAt) VALUES (?, ?)")
		if weHaveNotProblem(err, "PersistTicket") {
			prepare.Exec(ticket.Code, time.Now().Format(time.RFC3339))
		}
	}
}

func (repository *TicketRepository) fetchTickets(db *sql.DB) {
	rows, err := db.Query("SELECT * FROM ticket")
	if weHaveNotProblem(err, "FetchTickets") {
		var uid int
		var code, readAt string
		for rows.Next() {
			err := rows.Scan(&uid, &code, & readAt)
			if weHaveNotProblem(err) {
				repository.Tickets[code] = Ticket{ code, readAt, }
			}
		}
	}
}

func openConnectionFrom(file string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", file)
	return db, err
}

func weHaveNotProblem(err error, from ...string) (bool) {
	if err != nil {
		if from != nil {
			log.New(os.Stderr, "[FROM]: ", 0).Println(from)
		}
		log.New(os.Stderr, "[LOCAL-SERVER] ERROR: ", 0).Println(err.Error())
		return false
	}
	return true
}
