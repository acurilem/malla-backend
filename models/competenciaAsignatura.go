package models

type CompetenciaAsignatura struct {
	ID          uint   `gorm:"column:id_comp_asig;primaryKey",json:"id_comp_asig"`
	NomCompAsig string `json:"nom_comp_asig"`
}

type CompetenciasAsignaturas []CompetenciaAsignatura

// TableName overrides the table name used by User to `profiles`
func (CompetenciaAsignatura) TableName() string {
	return "competencia_asignatura"
}
