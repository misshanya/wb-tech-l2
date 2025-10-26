package event

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/misshanya/wb-tech-l2/18/internal/models"
)

type service interface {
	Create(e *models.Event) *models.Event
	GetForDay(date time.Time) []*models.Event
	Update(e *models.Event) error
	Delete(id int)
}

type handler struct {
	service service
}

func New(service service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) Setup(router *echo.Group) {
	router.POST("/create_event", h.Create)
	router.POST("/update_event", h.Update)
	router.POST("/delete_event/:id", h.Delete)
	router.GET("/events_for_day/:date", h.GetForDay)
}
