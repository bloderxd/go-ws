package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const port = ":8080"

func Routes(ticket * Ticket, ticketRepository *TicketRepository) {
	router := gin.Default()
	ticketHotValidationRoute(router, ticket, ticketRepository)
	router.Run(port)
}

func ticketHotValidationRoute(router *gin.Engine, ticket *Ticket, repository *TicketRepository) {
	postRouter(router, "/validate", func(c *gin.Context) {
		if err := c.ShouldBindJSON(ticket); err == nil {
			ticket.validateResponse(c, repository)
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error" : err.Error()})
		}
	})
}

func (ticket *Ticket) validateResponse(context *gin.Context, repository *TicketRepository) {
	if str, response := repository.Validate(ticket); str == "SUCCESS" {
		context.JSON(http.StatusOK, gin.H{})
	} else if str == "VALIDATED" {
		context.JSON(http.StatusBadRequest, gin.H{
			"error" : str,
			"read_at" : response.ReadAt,
		})
	} else {
		context.JSON(http.StatusBadRequest, gin.H{"error" : str})
	}
}

func postRouter(router *gin.Engine, endPoint string, listener func(c *gin.Context)) {
	router.POST(endPoint, func(context *gin.Context) { listener(context) })
}

func getRouter(router *gin.Engine, endPoint string, listener func(c *gin.Context)) {
	router.GET(endPoint, func(context *gin.Context) { listener(context) })
}