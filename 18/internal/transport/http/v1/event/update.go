package event

import (
	"errors"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/misshanya/wb-tech-l2/18/internal/errorz"
	"github.com/misshanya/wb-tech-l2/18/internal/models"
	"github.com/misshanya/wb-tech-l2/18/internal/transport/http/dto"
)

func (h *handler) Update(c echo.Context) error {
	var req dto.EventUpdateRequest
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
		ID:     req.ID,
		UserID: req.UserID,
		Title:  req.Title,
		Date:   date,
	}
	err = h.service.Update(event)
	switch {
	case errors.Is(err, errorz.EventNotFound):
		return c.JSON(http.StatusNotFound, dto.HTTPStatus{
			Code:    http.StatusNotFound,
			Message: err.Error(),
		})
	case err != nil:
		return c.JSON(http.StatusInternalServerError, dto.HTTPStatus{
			Code:    http.StatusInternalServerError,
			Message: errorz.InternalServerError.Error(),
		})
	}

	return c.NoContent(http.StatusOK)
}
