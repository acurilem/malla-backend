package models

type FacDep struct {
	ID                uint `gorm:"column:id_fac_dep;primaryKey",json:"id_fac_dep"`
	CodUnidadFacultad uint `json:"cod_unidad_facultad"`
	CodUnidadDepto    uint `json:"cod_unidad_depto"`

	CarrerasTotales []CarreraTotal `gorm:"foreignKey:id_fac_dep"`
	//TO DO: RELACION UNODAD FACULTAD  y UNIDAD DEPTO
}

type FacsDeps []FacDep

// TableName overrides the table name used by User to `profiles`
func (FacDep) TableName() string {
	return "fac_dep"
}
