package handler

import (
	"github.com/atadzan/more-tech/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) getAtms(c *gin.Context) {
	atms, err := h.services.Atm.GetAll(c.Request.Context())
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, newResponse("internal server error"))
		return
	}
	c.JSON(http.StatusOK, atms)
}

func (h *Handler) getAtmById(c *gin.Context) {
	var input models.InputId
	if err := c.BindUri(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, newResponse("invalid input param"))
		return
	}
	atm, err := h.services.Atm.GetById(c.Request.Context(), input.Id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, newResponse("internal server error"))
		return
	}
	c.JSON(http.StatusOK, atm)
}
