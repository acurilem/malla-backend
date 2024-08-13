package models

type CentroUniversitario struct {
	ID                    uint           `gorm:"column:cod_centro;primaryKey" json:"cod_centro"`
	NomCentro             string         `json:"nom_centro"`
	AnoInicio             uint           `json:"ano_inicio"`
	AnoCierre             uint           `json:"ano_cierre"`
	CodResolucionCreacion string         `json:"cod_resolucion_creacion"`
	CarrerasTotales       []CarreraTotal `gorm:"foreignKey:cod_centro"`
}

type CentrosUniversitarios []CentroUniversitario

// TableName overrides the table name used by User to `profiles`
func (CentroUniversitario) TableName() string {
	return "centro_universitario"
}
