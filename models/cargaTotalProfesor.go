package models

type CargaTotalProfesor struct {
	ID             uint       `gorm:"column:id_carga_total_profesor;primaryKey" json:"id_carga_total_profesor"`
	CodPersona     string     `json:"cod_persona"`
	CodEstadoCarga uint       `json:"cod_estado_carga"`
	TipoCarga      uint       `json:"tipo_carga"`
	IdCargaTotal   uint       `json:"id_carga_total"`
	CargaTotal     CargaTotal `gorm:"foreignKey:IdCargaTotal"`
}

type CargasTotalesProfesores []CargaTotalProfesor

// TableName overrides the table name used by User to `profiles`
func (CargaTotalProfesor) TableName() string {
	return "carga_total_profesores"
}
