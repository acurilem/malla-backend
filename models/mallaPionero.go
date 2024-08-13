package models

type MallaPionero struct {
	ID                uint   `gorm:"column:id_malla;primaryKey",json:"id_malla"`
	IdPlanes          uint   `json:"id_planes"`
	IdCursoPionero    uint   `json:"id_curso_pionero"`
	CodTipoCursoMalla uint   `json:"cod_tipo_curso_malla"`
	IncluyeGrado      uint   `json:"incluye_grado"`
	NroSem            uint   `json:"nro_sem"`
	NroAno            uint   `json:"nro_ano"`
	FrecuenciaADictar uint   `json:"frecuencia_a_dictar"`
	CursoBasico       uint   `json:"curso_basico"`
	Duracion          string `json:"duracion"`
	IdModsub          uint   `json:"id_modsub"`
	EsProyecto        uint   `json:"es_proyecto"`

	CursoPionero     CursoPionero     `gorm:"foreignKey:IdCursoPionero"`
	Plan             Plan             `gorm:"foreignKey:IdPlanes"`
	MallasRequisitos []MallaRequisito `gorm:"foreignKey:id_malla"`
	TipoCursoMalla   TipoCursoMalla   `gorm:"foreignKey:CodTipoCursoMalla"`
	CursosDictar     []CursoDictar    `gorm:"foreignKey:id_malla"`
}

type MallasPioneros []MallaPionero

// TableName overrides the table name used by User to `profiles`
func (MallaPionero) TableName() string {
	return "malla_pionero"
}
