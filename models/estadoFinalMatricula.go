package models

type EstadoFinalMatricula struct {
	ID          uint        `gorm:"column:cod_estado_final;primaryKey",json:"cod_estado_final"`
	Descripcion string      `json:"descripcion"`
	Matriculas  []Matricula `gorm:"foreignKey:cod_estado_final"`
}

// TableName overrides the table name used by User to `profiles`
func (EstadoFinalMatricula) TableName() string {
	return "estado_final_matricula"
}
