package models

type PropietarioCarrera struct {
	ID             uint   `gorm:"column:cod_propietario;primaryKey",json:"cod_propietario"`
	NomPropietario string `json:"nom_propietario"`

	CarrerasTotales []CarreraTotal `gorm:"foreignKey:cod_propietario"`
}

type PropietariosCarreras []PropietarioCarrera

// TableName overrides the table name used by User to `profiles`
func (PropietarioCarrera) TableName() string {
	return "propietario_carrera"
}
