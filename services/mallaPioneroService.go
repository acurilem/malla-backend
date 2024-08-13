package services

import (
	"github.com/citiaps/SID-UMAG-ESTUDIANTE-Malla-backend/config"
	"github.com/citiaps/SID-UMAG-ESTUDIANTE-Malla-backend/models"
)

const (
	CollectionNameMallaPionero = "MallaPionero"
)

func GetMallaSinCursosComplementariosService(planId int) ([]models.MallaPionero, error) {
	mallas_pioneros := models.MallasPioneros{}
	err := config.Database.
		Preload("CursoPionero.ModulosSubcompetencias.CompetenciasRc").
		Preload("CursoPionero.CursosPionerosCompetencias.CompetenciaRc").
		Preload("CursoPionero.CompetenciaAsignatura").
		Preload("MallasRequisitos.MallaPionero.CursoPionero.CompetenciaAsignatura").
		Preload("TipoCursoMalla").Where("id_planes = ?", planId).
		Where("cod_tipo_curso_malla <= ?", 2).
		Order("nro_ano, nro_sem").
		Find(&mallas_pioneros).Error
	return mallas_pioneros, err
}

func GetMallaCursosComplementariosService(planId int) ([]models.MallaPionero, error) {
	mallas_pioneros := models.MallasPioneros{}
	err := config.Database.Preload("CursoPionero.ModulosSubcompetencias.CompetenciasRc").
		Preload("CursoPionero.CursosPionerosCompetencias.CompetenciaRc").
		Preload("CursoPionero.CompetenciaAsignatura").
		Preload("MallasRequisitos.MallaPionero.CursoPionero.CompetenciaAsignatura").
		Preload("TipoCursoMalla").
		Where("id_planes = ?", planId).
		Where("cod_tipo_curso_malla >= ?", 3).
		Order("nro_ano, nro_sem").
		Find(&mallas_pioneros).Error
	return mallas_pioneros, err
}
