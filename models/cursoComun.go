package models

type CursoComun struct {
	ID                    uint                `gorm:"column:id_curso_comun;primaryKey" json:"id_curso_comun"`
	NomCursoComun         string              `json:"nom_curso_comun"`
	HrsCronoPresen        float64             `json:"hrs_crono_presen"`
	CodTipoCursoDictar    uint                `json:"cod_tipo_curso_dictar"`
	CodUnidadResponsable  uint                `json:"cod_unidad_responsable"`
	DetallesCursosComunes []DetalleCursoComun `gorm:"foreignKey:id_curso_comun"`
	CargaParcial          CargaParcial        `gorm:"foreignKey:cod_curso_dictar"`
}
type CursosComunes []CursoComun

// TableName overrides the table name used by User to `profiles`
func (CursoComun) TableName() string {
	return "cursos_comun"
}
