package models

type Carrera struct {
	ID         string `gorm:"column:cod_carrera;primaryKey" json:"cod_carrera"`
	NomCarrera string `json:"nom_carrera"`

	CarrerasTotales []CarreraTotal `gorm:"foreignKey:cod_carrera"`
}

type Carreras []Carrera

// TableName overrides the table name used by User to `profiles`
func (Carrera) TableName() string {
	return "carreras"
}
