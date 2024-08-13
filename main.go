// main.go

package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/citiaps/SID-UMAG-ESTUDIANTE-Malla-backend/docs"
	"github.com/citiaps/SID-UMAG-ESTUDIANTE-Malla-backend/middleware"
	"github.com/citiaps/SID-UMAG-ESTUDIANTE-Malla-backend/routes"
	"github.com/citiaps/SID-UMAG-ESTUDIANTE-Malla-backend/utils"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

// @title          Servicio de Mallas
// @version        1.0
// @description    Este servicio entrega las mallas y planes de estudio de un estudiante de la Universidad de Magallanes.
// @termsOfService http://swagger.io/terms/

// @contact.name  CITIAPS
// @contact.url   https://citiaps.cl
// @contact.email citiaps@usach.cl

// lincense.name  Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host     backend.localhost
// @BasePath /api/v1

// @securityDefinitions.apikey Bearer
// @in                         header
// @name                       Authorization
// @description                BearerJWTToken in Authorization Header

// @accept  json
// @produce json

// @schemes http https
func main() {
	utils.LoadEnv()
	gin.SetMode(os.Getenv("GIN_MODE"))
	app := gin.Default()
	gin.DefaultWriter = io.MultiWriter(os.Stdout, log.Writer())
	app.Use(middleware.CorsMiddleware())
	// Docs
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Version = "v1"
	docs.SwaggerInfo.Host = "backend.localhost"
	// Route docs
	app.GET("/api/v1/mallas/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	app.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "Servicio no encontrado."})
	})
	routes.InitRoutes(app)
	http.ListenAndServe(os.Getenv("ADDR"), app)
}
