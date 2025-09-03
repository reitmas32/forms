package controllers

import (
	"common/domain/customctx"
	"common/domain/logger"
	"common/interface/cdtos"
	"fomrs/internal/api/v1/answers/presentation/dtos"

	"github.com/gin-gonic/gin"
)

func (c *AnswerController) Create(ctx *gin.Context) {

	entry := logger.FromContext(ctx)

	entry.Info("Creating answer")

	cc := customctx.NewCustomContext(ctx.Request.Context())

	dto := cdtos.GetDTOWithResponse[dtos.CreateAnswerDTO](ctx, cc)

	if dto.Error != nil {
		entry.Error("Error getting dto", dto.Error)
		ctx.JSON(dto.StatusCode, dto.ToMapWithCustomContext(cc))
		return
	}

	command := dto.Data.ToCommand()

	response := c.service.Create(cc, &command)

	ctx.JSON(response.StatusCode, response.ToMapWithCustomContext(cc))

}
