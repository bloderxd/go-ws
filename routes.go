package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
)

func Routes(ticket * Ticket, ticketRepository *TicketRepository) {
	ticketHotValidationRoute(ticket, ticketRepository)
}

func ticketHotValidationRoute(ticket *Ticket, repository *TicketRepository) {
	ticket.postRouter("/validate", func(c *gin.Context) {
		if err := c.ShouldBindJSON(ticket); err == nil {
			ticket.validateResponse(c, repository)
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error" : err.Error()})
		}
	})
}

func (ticket *Ticket) validateResponse(context *gin.Context, repository *TicketRepository) {
	if str := repository.Validate(ticket); str == "SUCCESS" {
		fmt.Println(str)
		context.JSON(http.StatusOK, gin.H{})
	} else {
		context.JSON(http.StatusBadRequest, gin.H{"error" : str})
	}
}

func (ticket *Ticket) postRouter(endPoint string, listener func(c *gin.Context)) {
	router := gin.Default()
	router.POST(endPoint, func(context *gin.Context) { listener(context) })
	router.Run(":8080")
}