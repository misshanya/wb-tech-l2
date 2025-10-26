package event

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/misshanya/wb-tech-l2/18/internal/transport/http/dto"
)

func (h *handler) Delete(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.HTTPStatus{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	h.service.Delete(id)

	return c.NoContent(http.StatusOK)
}
