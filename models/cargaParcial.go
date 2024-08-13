package models

type CargaParcial struct {
	ID                      uint         `gorm:"column:id_carga_parcial;primaryKey" json:"id_carga_parcial"`
	TipoCurso               uint         `json:"tipo_curso"`
	CodCursoDictar          uint         `json:"cod_curso_dictar"`
	Sem                     uint         `json:"sem"`
	Ano                     uint         `json:"ano"`
	HrsCronoPresenSemestral float64      `json:"hrs_crono_presen_semestral"`
	HrsTeoria               float64      `json:"hrs_teoria"`
	HrsEjercicio            float64      `json:"hrs_ejercicio"`
	HrsPractica             float64      `json:"hrs_practica"`
	CantGruposTeoria        uint         `json:"cant_grupos_teoria"`
	CantGruposPractica      uint         `json:"cant_grupos_practica"`
	TotalHrsDocencia        float64      `json:"total_hrs_docencia"`
	CargasTotales           []CargaTotal `gorm:"foreignKey:id_carga_parcial"`
	Planes                  []Plan       `gorm:"foreignKey:id_carga_parcial"`
}

type CargasParciales []CargaParcial

type Tabler interface {
	TableName() string
}

// TableName overrides the table name used by User to `profiles`
func (CargaParcial) TableName() string {
	return "carga_parcial"
}
