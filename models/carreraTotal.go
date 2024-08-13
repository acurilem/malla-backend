package models

type CarreraTotal struct {
	ID                    uint   `gorm:"column:id_carrera_total;primaryKey" json:"id_carrera_total"`
	CodCarrera            string `json:"cod_carrera"`
	CodEspecialidad       uint   `json:"cod_especialidad"`
	CodMencion            uint   `json:"cod_mencion"`
	CodCentro             uint   `json:"cod_centro"`
	CodAreaCarrera        uint   `json:"cod_area_carrera"`
	CodPropietario        uint   `json:"cod_propietario"`
	CodNivelCarrera       uint   `json:"cod_curso_dictar"`
	IdFacDep              uint   `json:"id_fac_dep"`
	AreaConocimientoCnaId uint   `json:"area_conocimiento_cna_id"`
	NroDecretoCreacion    string `json:"nro_decreto_creacion"`
	CodigoSies            string `json:"codigo_sies"`
	AnoCreacion           uint   `json:"ano_creacion"`
	AnoCierre             uint   `json:"ano_cierre"`
	Regimen               string `json:"regimen"`
	Vigencia              uint   `json:"vigencia"`

	Carrera            Carrera             `gorm:"foreignKey:CodCarrera"`
	Centro             CentroUniversitario `gorm:"foreignKey:CodCentro"`
	Especialidad       Especialidad        `gorm:"foreignKey:CodEspecialidad"`
	FacDep             FacDep              `gorm:"foreignKey:IdFacDep"`
	AreaCarrera        AreaCarrera         `gorm:"foreignKey:CodAreaCarrera"`
	PropietarioCarrera PropietarioCarrera  `gorm:"foreignKey:CodPropietario"`
	NivelCarrera       NivelCarrera        `gorm:"foreignKey:CodNivelCarrera"`
	Mencion            Mencion             `gorm:"foreignKey:CodMencion"`

	Planes []Plan `gorm:"foreignKey:id_carrera_total"`
}

type CarrerasTotales []CarreraTotal

// TableName overrides the table name used by User to `profiles`
func (CarreraTotal) TableName() string {
	return "carrera_total"
}
