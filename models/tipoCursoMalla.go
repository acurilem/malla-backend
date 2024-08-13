package models

type TipoCursoMalla struct {
	ID          uint   `gorm:"column:cod_tipo_curso_malla;primaryKey",json:"cod_tipo_curso_malla"`
	Descripcion string `json:"descripcion"`
}

type TiposCursosMallas []TipoCursoMalla

// TableName overrides the table name used by User to `profiles`
func (TipoCursoMalla) TableName() string {
	return "tipo_curso_malla"
}
