package handler

import (
	"net/http"

	"github.com/ekkinox/yokai-es/internal/elastic"
	"github.com/labstack/echo/v4"
)

// ExampleHandler is an example HTTP handler.
type ExampleHandler struct {
	indexer *elastic.Indexer
}

// NewExampleHandler returns a new [ExampleHandler].
func NewExampleHandler(indexer *elastic.Indexer) *ExampleHandler {
	return &ExampleHandler{
		indexer: indexer,
	}
}

// Handle handles HTTP requests.
func (h *ExampleHandler) Handle() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := h.indexer.ListIndex(c.Request().Context())
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, res)
	}
}
