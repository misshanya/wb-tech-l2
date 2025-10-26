package event

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/misshanya/wb-tech-l2/18/internal/models"
	"github.com/misshanya/wb-tech-l2/18/internal/transport/http/dto"
)

func (h *handler) Create(c echo.Context) error {
	var req dto.EventCreateRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.HTTPStatus{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	date, err := time.Parse("2006-01-02", string(req.Date))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.HTTPStatus{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	event := &models.Event{
		UserID: req.UserID,
		Title:  req.Title,
		Date:   date,
	}
	newEvent := h.service.Create(event)

	resp := &dto.EventCreateResponse{
		ID:     newEvent.ID,
		UserID: newEvent.UserID,
		Title:  newEvent.Title,
		Date:   newEvent.Date,
	}
	return c.JSON(http.StatusCreated, resp)
}
