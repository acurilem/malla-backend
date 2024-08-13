package models

type CompetenciaRcSubcomp struct {
	ID            uint `gorm:"column:id_modsub;primaryKey",json:"id_modsub"`
	IdCompetencia uint `json:"id_competencia"`
}

// TableName overrides the table name used by User to `profiles`
func (CompetenciaRcSubcomp) TableName() string {
	return "competencia_rc_subcomp"
}
