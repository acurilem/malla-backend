package models

type NivelGlobalCarrera struct {
	ID          uint   `gorm:"column:cod_nivel_global;primaryKey",json:"cod_nivel_global"`
	Descripcion string `json:"descripcion"`
	Planes      []Plan `gorm:"foreignKey:cod_nivel_global"`
}

type NivelesGlobalesCarreras []NivelGlobalCarrera

// TableName overrides the table name used by User to `profiles`
func (NivelGlobalCarrera) TableName() string {
	return "nivel_global_carrera"
}
