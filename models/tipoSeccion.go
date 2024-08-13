package models

type TipoSeccion struct {
	ID             uint   `gorm:"column:cod_tipo_seccion;primaryKey",json:"cod_tipo_seccion"`
	NomTipoSeccion string `json:"nom_tipo_seccion"`
}

// TableName overrides the table name used by User to `profiles`
func (TipoSeccion) TableName() string {
	return "tipo_seccion"
}
