package event

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/misshanya/wb-tech-l2/18/internal/errorz"
	"github.com/misshanya/wb-tech-l2/18/internal/transport/http/dto"
)

func (h *handler) GetForMonth(c echo.Context) error {
	dateStr := c.Param("date")
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.HTTPStatus{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	events := h.service.GetForMonth(date)
	if len(events) == 0 {
		return c.JSON(http.StatusNotFound, dto.HTTPStatus{
			Code:    http.StatusNotFound,
			Message: errorz.EventsNotFound.Error(),
		})
	}

	resp := make(dto.EventsGetForMonthResponse, len(events))
	for i, event := range events {
		resp[i] = dto.Event{
			ID:     event.ID,
			UserID: event.UserID,
			Title:  event.Title,
			Date:   event.Date,
		}
	}
	return c.JSON(http.StatusOK, resp)
}
