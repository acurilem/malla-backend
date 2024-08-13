package models

type CompetenciaRc struct {
	ID             uint   `gorm:"column:id_competencia;primaryKey",json:" id_competencia"`
	NomCompetencia string `json:"nom_competencia"`
}

type CompetenciasRc []CompetenciaRc

// TableName overrides the table name used by User to `profiles`
func (CompetenciaRc) TableName() string {
	return "competencias_rc"
}
