package controllers

import (
	"common/domain/customctx"
	"common/domain/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *FormsController) Retrieve(ctx *gin.Context) {

	entry := logger.FromContext(ctx)

	entry.Info("Retrieving form")

	cc := customctx.NewCustomContext(ctx)

	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":      "id is required",
			"success":    false,
			"statusCode": http.StatusBadRequest,
		})
		return
	}

	response := c.formsService.Retrieve(cc, id)

	ctx.JSON(response.StatusCode, response.ToMapWithCustomContext(cc))

}
