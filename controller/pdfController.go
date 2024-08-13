package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/citiaps/SID-UMAG-ESTUDIANTE-Malla-backend/services"

	"github.com/gin-gonic/gin"
)

const (
	CollectionNamePdf = "Pdf"
)

// GenerarMalla godoc
// @Summary        Genera un archivo PDF con la malla del usuario autenticado
// @Description    Esta función genera un archivo PDF con la malla del usuario autenticado en base al Bearer Token que viene en el encabezado del request.
// @Tags           Pdf
// @Accept         json
// @Consumes	   application/json
// @Produces       application/pdf
// @Method         GET
// @Param          id path string true "id de la malla a consultar"
// @Security       Bearer
// @Success        200 {file} string "Un archivo PDF con la malla del usuario autenticado."
// @Failure        404 {object} ErrorResponse "No se encontró datos."
// @Router         /mallas/pdf/malla/{id} [get]
func GenerarMalla(ctx *gin.Context) {
	idMalla, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	m := services.GetPdfService(idMalla)

	document, err := m.Generate()
	if err != nil {
		log.Fatal(err.Error())
	}
	buf := document.GetBytes()
	ctx.Header("Content-Type", "application/pdf")
	ctx.Header("Content-Disposition", "attachment; filename=Malla.pdf")
	ctx.Data(http.StatusOK, "application/pdf", buf)
}
