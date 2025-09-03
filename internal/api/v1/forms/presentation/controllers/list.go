package controllers

import (
	"common/domain/customctx"
	"common/domain/logger"

	"github.com/gin-gonic/gin"
)

func (c *FormsController) List(ctx *gin.Context) {

	entry := logger.FromContext(ctx)

	entry.Info("Listing forms")

	cc := customctx.NewCustomContext(ctx)

	response := c.formsService.List(cc)

	ctx.JSON(response.StatusCode, response.ToMapWithCustomContext(cc))
}
