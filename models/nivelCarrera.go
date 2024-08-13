package models

type NivelCarrera struct {
	ID          uint   `gorm:"column:cod_nivel_carrera;primaryKey",json:"cod_nivel_carrera"`
	Descripcion string `json:"descripcion"`

	CarrerasTotales []CarreraTotal `gorm:"foreignKey:cod_nivel_carrera"`
	Planes          []Plan         `gorm:"foreignKey:cod_nivel_carrera"`
}

type NivelesCarreras []NivelCarrera

// TableName overrides the table name used by User to `profiles`
func (NivelCarrera) TableName() string {
	return "nivel_carrera"
}
