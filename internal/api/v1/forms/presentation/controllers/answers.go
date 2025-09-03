package controllers

import (
	"common/domain/customctx"
	"common/domain/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *FormsController) Answers(ctx *gin.Context) {

	entry := logger.FromContext(ctx)

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
	entry.Info("Retrieving answers of form: ", id)

	response := c.formsService.Answers(cc, id)

	ctx.JSON(response.StatusCode, response.ToMapWithCustomContext(cc))

}
