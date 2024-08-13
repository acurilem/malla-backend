package controller

import (
	"net/http"
	"strconv"

	"github.com/citiaps/SID-UMAG-ESTUDIANTE-Malla-backend/services"
	"github.com/gin-gonic/gin"
)

const (
	CollectionNameMallaPionero = "MallaPionero"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

// GetMallaSinCursosComplementarios godoc
// @Summary        Obtiene la malla sin cursos complementarios del usuario autenticado
// @Description    Esta función obtiene la malla sin cursos complementarios del usuario autenticado en base al Bearer Token que viene en el encabezado del request.
// @Tags           Malla
// @Accept         json
// @Product        json
// @Method         GET
// @Param          id path string true "id de la malla a consultar"
// @Security       Bearer
// @Success		   200 {array} models.MallaPionero "Consulta realizada con éxito"
// @Failure        404 {object} ErrorResponse "No se encontró datos."
// @Router         /mallas/sin_cursos_complementarios/{id} [get]
func GetMallaSinCursosComplementarios(ctx *gin.Context) {
	planId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(404, gin.H{"error": err.Error()})
		return
	}

	resultMallaPionero, err := services.GetMallaSinCursosComplementariosService(planId)
	if err != nil {
		ctx.JSON(404, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, resultMallaPionero)
}

// GetMallaCursosComplementarios godoc
// @Summary        Obtiene la malla con cursos complementarios del usuario autenticado
// @Description    Esta función obtiene la malla con cursos complementarios del usuario
// @Tags           Malla
// @Accept         json
// @Product        json
// @Method         GET
// @Param          id path string true "id de la malla a consultar"
// @Security       Bearer
// @Success		   200 {array} models.MallaPionero "Consulta realizada con éxito"
// @Failure        404 {object} ErrorResponse "No se encontró datos."
// @Router         /mallas/cursos_complementarios/{id} [get]
func GetMallaCursosComplementarios(ctx *gin.Context) {
	planId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(404, gin.H{"error": err.Error()})
		return
	}

	resultMallaPionero, err := services.GetMallaCursosComplementariosService(planId)
	if err != nil {
		ctx.JSON(404, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, resultMallaPionero)
}
