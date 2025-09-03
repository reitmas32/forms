package controllers

import (
	"common/domain/customctx"
	"common/domain/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *AnswerController) Retrieve(ctx *gin.Context) {

	entry := logger.FromContext(ctx)

	cc := customctx.NewCustomContext(ctx)

	id := ctx.Param("id")

	entry.Info("Retrieving answer: ", id)

	if id == "" {
		entry.Error("id is required")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":      "id is required",
			"success":    false,
			"statusCode": http.StatusBadRequest,
		})
		return
	}

	response := c.service.Retrieve(cc, id)

	ctx.JSON(response.StatusCode, response.ToMapWithCustomContext(cc))

}
