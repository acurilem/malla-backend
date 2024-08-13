package models

type ModuloSubcompetencia struct {
	ID                uint            `gorm:"column:id_modsub;primaryKey",json:"id_modsub"`
	NomModsub         string          `json:"nom_modsub"`
	CodTipoModulo     uint            `json:"cod_tipo_modulo"`
	Sct               uint            `json:"sct"`
	HrsPresenciales   float64         `json:"hrs_presenciales"`
	HrsNoPresenciales float64         `json:"hrs_no_presenciales"`
	HrsAGuiado        float64         `json:"hrs_a_guiado"`
	Institucional     uint            `json:"institucional"`
	CompetenciasRc    []CompetenciaRc `gorm:"many2many:competencia_rc_subcomp;foreignKey:ID;joinForeignKey:IdModsub;References:ID;joinReferences:IdCompetencia"`
}

type ModulosSubcompetencias []ModuloSubcompetencia

// TableName overrides the table name used by User to `profiles`
func (ModuloSubcompetencia) TableName() string {
	return "modulo_subcompetencia"
}
