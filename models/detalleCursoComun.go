package models

type DetalleCursoComun struct {
	IdCursoDictar uint `gorm:"column:id_curso_dictar;primaryKey",json:"id_curso_dictar"`
	IdCursoComun  uint `gorm:"column:id_curso_comun;primaryKey",json:"id_curso_comun"`
	CodTipoPlan   uint `json:"cod_tipo_plan"`

	CursoDictar CursoDictar `gorm:"foreignKey:IdCursoDictar"`
	CursoComun  CursoComun  `gorm:"foreignKey:IdCursoComun"`
	TipoPlan    TipoPlan    `gorm:"foreignKey:CodTipoPlan"`
}

type DetallesCursosComunes []DetalleCursoComun

// TableName overrides the table name used by User to `profiles`
func (DetalleCursoComun) TableName() string {
	return "detalle_curso_comun"
}
