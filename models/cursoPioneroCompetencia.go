package models

type CursoPioneroCompetencia struct {
	ID             uint          `gorm:"column:id_curso_pionero_comp;primaryKey" json:" id_curso_pionero_comp"`
	IdCursoPionero uint          `json:"id_curso_pionero"`
	IdCompetencia  uint          `json:"id_competencia"`
	IdNivelComp    uint          `json:"id_nivel_comp"`
	CompetenciaRc  CompetenciaRc `gorm:"foreignKey:IdCompetencia"`
}

type CursosPionerosCompetencias []CursoPioneroCompetencia

// TableName overrides the table name used by User to `profiles`
func (CursoPioneroCompetencia) TableName() string {
	return "cursos_pionero_competencia"
}
