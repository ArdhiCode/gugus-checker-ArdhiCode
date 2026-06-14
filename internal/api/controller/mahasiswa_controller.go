package controller

import (
	"github.com/ArdhiCode/gugus-checker-ArdhiCode/internal/api/service"
	"github.com/ArdhiCode/gugus-checker-ArdhiCode/internal/dto"
	"github.com/ArdhiCode/gugus-checker-ArdhiCode/internal/pkg/response"
	"github.com/gin-gonic/gin"
)

type MahasiswaController interface {
	GetByNRP(ctx *gin.Context)
}

type mahasiswaController struct {
	mahasiswaService service.MahasiswaService
}

func NewMahasiswa(mahasiswaService service.MahasiswaService) MahasiswaController {
	return &mahasiswaController{mahasiswaService}
}

func (c *mahasiswaController) GetByNRP(ctx *gin.Context) {
	var req dto.GetByNRPRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.NewFailedWithCode(400, "invalid body request", err).Send(ctx)
		return
	}

	res, err := c.mahasiswaService.GetByNRP(ctx.Request.Context(), req)
	if err != nil {
		response.NewFailedWithCode(401, "failed to get data", err).Send(ctx)
		return
	}

	response.NewSuccess("success get data", res).Send(ctx)
}
