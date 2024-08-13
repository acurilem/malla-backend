package models

type Matricula struct {
	ID                   uint                 `gorm:"column:id_matricula;primaryKey" json:"id_matricula"`
	CodPersona           uint                 `json:"cod_persona"`
	RutAlum              string               `json:"rut_alum"`
	AnoProceso           uint                 `json:"ano_proceso"`
	AnoMatricula         uint                 `json:"ano_matricula"`
	CodTipoAlumno        uint                 `json:"cod_tipo_alumno"`
	CodModoIngreso       uint                 `json:"cod_modo_ingreso"`
	AnoIngreso           uint                 `json:"ano_ingreso"`
	IdPlanes             uint                 `json:"id_planes"`
	CodTipoMatricula     uint                 `json:"cod_tipo_matricula"`
	CodEstadoFinal       uint                 `json:"cod_estado_final"`
	AnoPrimerIngreso     uint                 `json:"ano_primer_ingreso"`
	SemDesvinculacion    uint                 `json:"sem_desvinculacion"`
	Plan                 Plan                 `gorm:"foreignKey:IdPlanes"`
	EstadoFinalMatricula EstadoFinalMatricula `gorm:"foreignKey:CodEstadoFinal"`
}

type Matriculas []Matricula

// TableName overrides the table name used by User to `profiles`
func (Matricula) TableName() string {
	return "matricula"
}
