package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
	"os"
	"fmt"
)

func ConfigAutoMigration() {
	db, err := canOpenConnection()
	if weHaveNotProblemWith(err) && !db.HasTable(&Ticket{}) {
		defer db.Close()
		db.AutoMigrate(&Ticket{})
	}
}

func(repository *TicketRepository) FetchAllTickets()  {
	db, err := canOpenConnection()
	if weHaveNotProblemWith(err) {
		defer db.Close()
		db.Find(&repository.Tickets)
		fmt.Println(repository.Tickets)
	}
}

func(ticket *Ticket) PersistTicket() {
	db, err := canOpenConnection()
	if weHaveNotProblemWith(err) {
		defer db.Close()
		db.NewRecord(&ticket)
		db.Create(&ticket)
	}
}

func canOpenConnection() (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", "ticket.db")
	if weHaveNotProblemWith(err) {
		return db, err
	}
	return nil, err
}

func weHaveNotProblemWith(err error, from ...string) (bool) {
	if err != nil {
		if from != nil {
			log.New(os.Stderr, "[FROM]: ", 0).Println(from)
		}
		log.New(os.Stderr, "[LOCAL-SERVER] ERROR: ", 0).Println(err.Error())
		return false
	}
	return true
}