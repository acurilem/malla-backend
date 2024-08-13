package models

type Plan struct {
	ID              uint   `gorm:"column:id_planes;primaryKey",json:"id_planes"`
	IdCarreraTotal  uint   `json:"id_carrera_total"`
	AnoPlan         uint   `gorm:"column:anoplan;", json:"anoplan"`
	Vigencia        uint   `json:"vigencia"`
	CodTipoPlan     uint   `json:"cod_tipo_plan"`
	Duracion        uint   `json:"duracion"`
	CodRegimen      uint   `json:"cod_regimen"`
	CodNivelGlobal  uint   `json:"cod_nivel_global"`
	CodNivelCarrera uint   `json:"cod_nivel_carrera"`
	Caracteristicas string `json:"caracteristicas"`
	CodJornada      uint   `json:"cod_jornada"`
	CodModalidad    uint   `gorm:"column:cod_modalidad;", json:"cod_modalidad"`
	IdCargaParcial  uint   `json:"id_carga_parcial"`

	CarreraTotal       CarreraTotal       `gorm:"foreignKey:IdCarreraTotal"`
	CargaParcial       CargaParcial       `gorm:"foreignKey:IdCargaParcial"`
	TipoPlan           TipoPlan           `gorm:"foreignKey:CodTipoPlan"`
	NivelCarrera       NivelCarrera       `gorm:"foreignKey:CodNivelCarrera"`
	NivelGlobalCarrera NivelGlobalCarrera `gorm:"foreignKey:CodNivelGlobal"`
	Regimen            Regimen            `gorm:"foreignKey:CodRegimen"`
	Matriculas         []Matricula        `gorm:"foreignKey:id_planes"`
	MallasPioneros     []MallaPionero     `gorm:"foreignKey:id_malla"`
}

type Planes []Plan

// TableName overrides the table name used by User to `profiles`
func (Plan) TableName() string {
	return "planes"
}
