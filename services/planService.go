package services

import (
	"fmt"

	"github.com/citiaps/SID-UMAG-ESTUDIANTE-Malla-backend/config"
	"github.com/citiaps/SID-UMAG-ESTUDIANTE-Malla-backend/models"
)

const (
	CollectionNamePlan = "Plan"
)

func GetPlanesMatriculaAlumnoService(codPersona int) ([]uint, error) {
	var planes_ids []uint
	matricula := models.Matricula{}
	err := config.Database.Model(&matricula).Where("cod_persona = ?", codPersona).Distinct().Pluck("id_planes", &planes_ids).Error
	return planes_ids, err
}

func GetPlanesAlumnoService(codPersona int) ([]models.Plan, error) {
	var planesAlumno []models.Plan

	planesIDs, err := GetPlanesMatriculaAlumnoService(codPersona)
	if err != nil {
		return nil, err
	}
	db := config.Database.Debug()
	query := db.
		Preload("CarreraTotal.AreaCarrera").
		Preload("CarreraTotal.PropietarioCarrera").
		Preload("CarreraTotal.Carrera").
		Preload("CarreraTotal.Centro").
		Preload("CarreraTotal.Especialidad").
		Preload("CarreraTotal.Mencion").
		Preload("CarreraTotal.NivelCarrera").
		Preload("NivelGlobalCarrera").
		Preload("TipoPlan").
		Preload("Regimen").
		Where("id_planes IN (?)", planesIDs).
		Where("vigencia IN (?)", 1).
		Where("cod_nivel_global != ?", 2).
		Find(&planesAlumno).
		Statement
	if query.Error != nil {
		fmt.Println("Error:", query.Error)
		return nil, query.Error
	}
	sql := query.SQL.String()
	fmt.Println(sql)
	return planesAlumno, nil
}

func GetPlanService(idPlan int) ([]models.Plan, error) {
	var plan []models.Plan
	db := config.Database.Debug()
	query := db.
		Preload("CarreraTotal.AreaCarrera").
		Preload("CarreraTotal.PropietarioCarrera").
		Preload("CarreraTotal.Carrera").
		Preload("CarreraTotal.Centro").
		Preload("CarreraTotal.Especialidad").
		Preload("CarreraTotal.Mencion").
		Preload("CarreraTotal.NivelCarrera").
		Preload("NivelGlobalCarrera").
		Preload("TipoPlan").
		Preload("Regimen").
		Where("id_planes IN (?)", idPlan).
		Where("vigencia IN (?)", 1).
		Where("cod_nivel_global != ?", 2).
		Find(&plan).
		Statement
	if query.Error != nil {
		fmt.Println("Error:", query.Error)
		return nil, query.Error
	}
	sql := query.SQL.String()
	fmt.Println(sql)
	return plan, nil
}

func GetAllPlanesService() ([]models.Plan, error) {
	var planes []models.Plan
	db := config.Database.Debug()
	query := db.
		Preload("CarreraTotal.AreaCarrera").
		Preload("CarreraTotal.PropietarioCarrera").
		Preload("CarreraTotal.PropietarioCarrera", "cod_propietario = ?", "1").
		Preload("CarreraTotal.Carrera").
		Preload("CarreraTotal.Centro").
		Preload("CarreraTotal.Especialidad").
		Preload("CarreraTotal.Mencion").
		Preload("CarreraTotal.NivelCarrera").
		Preload("NivelGlobalCarrera").
		Preload("TipoPlan").
		Preload("Regimen").
		Where("vigencia IN (?)", 1).

		//	Where("cod_nivel_global != ?", 2).
		Find(&planes).
		Statement
	if query.Error != nil {
		fmt.Println("Error:", query.Error)
		return nil, query.Error
	}
	sql := query.SQL.String()
	fmt.Println(sql)
	return planes, nil
}
