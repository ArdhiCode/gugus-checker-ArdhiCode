package routes

import (
	"github.com/ArdhiCode/gugus-checker-ArdhiCode/internal/api/controller"
	"github.com/gin-gonic/gin"
)

func Mahasiswa(app *gin.Engine, mahasiswaController controller.MahasiswaController) {

	routes := app.Group("/api/v1/mahasiswa")
	{
		routes.POST("/get-by-nrp", mahasiswaController.GetByNRP)
	}

}
