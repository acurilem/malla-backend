package models

type Especialidad struct {
	ID              uint           `gorm:"column:cod_especialidad;primaryKey",json:"cod_especialidad"`
	NomEspecialidad string         `json:"nom_especialidad"`
	CarrerasTotales []CarreraTotal `gorm:"foreignKey:cod_especialidad"`
}

type Especialidades []Especialidad

// TableName overrides the table name used by User to `profiles`
func (Especialidad) TableName() string {
	return "especialidades"
}
