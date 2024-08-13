package controller

import (
	"net/http"
	"strconv"

	"github.com/citiaps/SID-UMAG-ESTUDIANTE-Malla-backend/services"

	"github.com/gin-gonic/gin"
)

const (
	CollectionNamePlan = "Plan"
)

// GetPlanesAlumno godoc
// @Summary        Obtiene los planes de estudio del usuario autenticado
// @Description    Esta función obtiene los planes de estudio del usuario autenticado en base al Bearer Token que viene en el encabezado del request.
// @Tags           Plan
// @Accept         json
// @Product        json
// @Method         GET
// @Security       Bearer
// @Success		   200 {array} models.Plan "Consulta realizada con éxito"
// @Failure        404 {object} ErrorResponse "No se encontró datos."
// @Router         /mallas/plan/carreras/ [get]
func GetPlanesAlumno(ctx *gin.Context) {
	user, err := services.GetUser(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	codPersona := int(user.ID)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resultsPlanes, err := services.GetPlanesAlumnoService(codPersona)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, resultsPlanes)
}

// GetPlan godoc
// @Summary        Obtiene un plan de estudio
// @Description    Esta función obtiene un plan de estudio en base a su id.
// @Tags           Plan
// @Accept         json
// @Product        json
// @Method         GET
// @Security       Bearer
// @Param          id path string true "id del plan a consultar"
// @Success		   200 {array} models.Plan "Consulta realizada con éxito"
// @Failure        404 {object} ErrorResponse "No se encontró datos."
// @Router         /mallas/plan/plan/{id} [get]
func GetPlan(ctx *gin.Context) {
	idPlan, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//	codPersona := 48484
	resultsPlan, err := services.GetPlanService(idPlan)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, resultsPlan)
}

// GetAllPlanes godoc
// @Summary        Obtiene todos los planes de estudio
// @Description    Esta función obtiene todos los planes de estudio.
// @Tags           Plan
// @Accept         json
// @Product        json
// @Method         GET
// @Security       Bearer
// @Success		   200 {array} models.Plan "Consulta realizada con éxito"
// @Failure        404 {object} ErrorResponse "No se encontró datos."
// @Router         /mallas/plan/ [get]
func GetAllPlanes(ctx *gin.Context) {

	resultsPlanes, err := services.GetAllPlanesService()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, resultsPlanes)
}
