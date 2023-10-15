package handler

import (
	"github.com/atadzan/more-tech/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) getOffices(c *gin.Context) {
	offices, err := h.services.Office.GetAll(c.Request.Context())
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, newResponse("internal server error"))
		return
	}
	c.JSON(http.StatusOK, offices)
}

func (h *Handler) getOfficeById(c *gin.Context) {
	var input models.InputId
	if err := c.ShouldBindUri(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, newResponse("invalid input param"))
		return
	}

	office, err := h.services.Office.GetById(c.Request.Context(), input.Id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, newResponse("internal server error"))
		return
	}
	c.JSON(http.StatusOK, office)
}
