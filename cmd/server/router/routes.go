package router

import (
	"github.com/bootcamp-go/desafio-go-web/cmd/server/handler"
	"github.com/bootcamp-go/desafio-go-web/internal/domain"
	"github.com/bootcamp-go/desafio-go-web/internal/tickets"
	"github.com/gin-gonic/gin"
)

type router struct {
	engine     *gin.Engine
	controller *handler.Controller
}

// NewRouter devuelve un gin.engine con todas las dependencias inyectadas y el handler inicializado
func NewRouter(routerBase *gin.Engine, list []domain.Ticket) router {
	router := router{engine: routerBase}
	repository := tickets.NewRepository(list)
	service := tickets.NewService(repository)
	router.controller = handler.NewController(service)
	return router
}

func (r *router) MapRoutes() {
	r.engine.GET("/ping", r.controller.Ping())

	ticket := r.engine.Group("/ticket")
	{
		ticket.GET("getByCountry/:dest", r.controller.GetTicketsByCountry())
		ticket.GET("getPercentage/:dest", r.controller.GetDestinationPercentage())
	}
}
