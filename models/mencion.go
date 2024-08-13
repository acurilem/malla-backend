package models

type Mencion struct {
	ID         uint   `gorm:"column:cod_mencion;primaryKey",json:"cod_mencion"`
	NomMencion string `json:"nom_mencion"`

	CarrerasTotales []CarreraTotal `gorm:"foreignKey:cod_mencion"`
}

type Menciones []Mencion

// TableName overrides the table name used by User to `profiles`
func (Mencion) TableName() string {
	return "menciones"
}
