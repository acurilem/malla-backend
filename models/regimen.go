package models

type Regimen struct {
	ID         uint   `gorm:"column:cod_regimen;primaryKey",json:"cod_regimen"`
	NomRegimen string `json:"nom_regimen"`
	Vigente    int    `json:"vigente"`
}

type Regimenes []Regimen

// TableName overrides the table name used by User to `profiles`
func (Regimen) TableName() string {
	return "regimen"
}
