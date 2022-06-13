package handlers

import (
	"net/http"

	"github.com/MalukiMuthusi/cryptocurrencies/internal/logger"
	"github.com/MalukiMuthusi/cryptocurrencies/internal/models"
	"github.com/MalukiMuthusi/cryptocurrencies/internal/store"
	"github.com/gin-gonic/gin"
)

// List cryptocurrencies saved in the databases
type List struct {
	Store store.Store
}

// Handle list cryptocurrencies saved in the databases
func (h List) Handle(c *gin.Context) {

	l, err := h.Store.List(c.Request.Context())

	if err != nil {

		logger.Log.Info(err)

		c.JSON(http.StatusInternalServerError, models.BasicError{
			Code:    "SERVER_ERROR",
			Message: "failed to get cryptocurrencies",
		})

		return
	}

	c.JSON(http.StatusOK, l)
}
