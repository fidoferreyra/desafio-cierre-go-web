package handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/bootcamp-go/desafio-go-web/internal/tickets"
	"github.com/bootcamp-go/desafio-go-web/pkg/web"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	service tickets.Service
}

func NewController(s tickets.Service) *Controller {
	return &Controller{
		service: s,
	}
}

func (s *Controller) Ping() gin.HandlerFunc {
	return func(c *gin.Context) {
		web.Success(c, http.StatusOK, "pong")
	}
}

func (s *Controller) GetTicketsByCountry() gin.HandlerFunc {
	return func(c *gin.Context) {

		destination := c.Param("dest")

		if destination == "" {
			web.Failure(c, http.StatusBadRequest, errors.New("destination is empty or not valid"))
			return
		}

		tickets, err := s.service.GetTotalTickets(c, destination)
		if err != nil {
			web.Failure(c, http.StatusInternalServerError, err)
			return
		}

		responseBody := fmt.Sprintf("Hay %d tickets con destino a %s", tickets, destination)
		web.Success(c, http.StatusOK, responseBody)
	}
}

func (s *Controller) GetDestinationPercentage() gin.HandlerFunc {
	return func(c *gin.Context) {

		destination := c.Param("dest")

		if destination == "" {
			web.Failure(c, http.StatusBadRequest, errors.New("destination is empty or not valid"))
			return
		}

		pct, err := s.service.GetDestinationPercentage(c, destination)
		if err != nil {
			web.Failure(c, http.StatusInternalServerError, err)
			return
		}
		ticketPercent := fmt.Sprintf("%.2f%%", pct)
		responseBody := fmt.Sprintf("El porcentaje de tickets con destino a %s es del %s", destination, ticketPercent)
		web.Success(c, http.StatusOK, responseBody)
	}
}
