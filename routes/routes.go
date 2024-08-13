package routes

import (
	"github.com/citiaps/SID-UMAG-ESTUDIANTE-Malla-backend/controller"
	"github.com/citiaps/SID-UMAG-ESTUDIANTE-Malla-backend/middleware"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		authenticatedRoutes := v1.Group("/mallas")
		authenticatedRoutes.Use(middleware.SetRoles("ALUMNOS"), middleware.AuthMiddleware())
		{
			authenticatedRoutes.GET("/sin_cursos_complementarios/:id", controller.GetMallaSinCursosComplementarios)
			authenticatedRoutes.GET("/cursos_complementarios/:id", controller.GetMallaCursosComplementarios)

			planGroup := authenticatedRoutes.Group("/plan")
			{
				planGroup.GET("/carreras/", controller.GetPlanesAlumno)
				planGroup.GET("/plan/:id", controller.GetPlan)
				planGroup.GET("/", controller.GetAllPlanes)
			}

			pdfGroup := authenticatedRoutes.Group("/pdf")
			{
				pdfGroup.GET("/malla/:id", controller.GenerarMalla)
			}
		}
	}
}
