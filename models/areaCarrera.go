package models

type AreaCarrera struct {
	ID              uint           `gorm:"column:cod_area_carrera;primaryKey" json:" cod_carea_carrera"`
	NomAreaCarrera  string         `json:"nom_area_carrera"`
	CarrerasTotales []CarreraTotal `gorm:"foreignKey:cod_area_carrera"`
}

type AreasCarreras []AreaCarrera

// TableName overrides the table name used by User to `profiles`
func (AreaCarrera) TableName() string {
	return "area_carrera"
}
