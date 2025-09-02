package controllers

import (
	"common/domain/customctx"
	"common/domain/logger"
	"common/interface/cdtos"
	"fomrs/internal/api/v1/forms/presentation/dtos"

	"github.com/gin-gonic/gin"
)

func (c *FormsController) Create(ctx *gin.Context) {

	entry := logger.FromContext(ctx)

	entry.Info("Creating form")

	cc := customctx.NewCustomContext(ctx)

	dto := cdtos.GetDTOWithResponse[dtos.CreateFormDTO](ctx, cc)

	if dto.Error != nil {
		ctx.JSON(dto.StatusCode, dto.ToMapWithCustomContext(cc))
		return
	}

	response := c.formsService.CreateForm(cc, dto.Data.ToCommand())

	ctx.JSON(response.StatusCode, response.ToMapWithCustomContext(cc))
}
