package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler)GetHelloHandler(c *gin.Context) {

	resp, err := h.service.GetHello(c)

	if err != nil {
		return
	}

	c.JSON(http.StatusOK, resp)
}