package handler

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"stepp-backend/src/constant"
	basepayload "stepp-backend/src/model/base"
	"strconv"
)

func (h *handler) GetPostHandler(c *gin.Context) {

	postId := c.Param("post-id")

	id, err := strconv.Atoi(postId)
	if err != nil {
		c.JSON(http.StatusBadRequest, basepayload.GetErrorResponse(constant.REQUEST_ERROR, "Post ID Invalid"))
		return
	}

	resp, err := h.service.GetPost(c, id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusOK, basepayload.GetErrorResponse(constant.GENERIC_ERROR, "Post ID Not Found"))
			return
		}
		c.JSON(http.StatusBadRequest, basepayload.GetErrorResponse(constant.GENERIC_ERROR, err))
		return
	}

	c.JSON(http.StatusOK, resp)
}
