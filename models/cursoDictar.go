package models

type CursoDictar struct {
	ID                    uint                `gorm:"column:id_curso_dictar;primaryKey",json:"id_curso_dictar"`
	IdMalla               uint                `json:"id_malla"`
	Sem                   uint                `json:"sem"`
	Ano                   uint                `json:"ano"`
	CodEstadoCurso        uint                `json:"cod_estado_curso"`
	CodUnidadResponsable  uint                `json:"cod_unidad_responsable"`
	HrsCronoPresen        float64             `json:"hrs_crono_presen"`
	SeDicta               uint                `json:"se_dicta"`
	Modulo                uint                `json:"modulo"`
	DetallesCursosComunes []DetalleCursoComun `gorm:"foreignKey:id_curso_dictar"`
	CargasParciales       []CargaParcial      `gorm:"foreignKey:cod_curso_dictar"`
}
type CursosDictar []CursoDictar

// TableName overrides the table name used by User to `profiles`
func (CursoDictar) TableName() string {
	return "cursos_dictar"
}
