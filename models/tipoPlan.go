package models

type TipoPlan struct {
	ID                    uint                `gorm:"column:cod_tipo_plan;primaryKey",json:"cod_tipo_plan"`
	NomTipoPlan           string              `json:"nom_tipo_plan"`
	Planes                []Plan              `gorm:"foreignKey:cod_tipo_plan"`
	DetallesCursosComunes []DetalleCursoComun `gorm:"foreignKey:cod_tipo_plan"`
}

type TiposPlanes []TipoPlan

// TableName overrides the table name used by User to `profiles`
func (TipoPlan) TableName() string {
	return "tipo_plan"
}
