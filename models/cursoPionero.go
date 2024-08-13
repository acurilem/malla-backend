package models

type CursoPionero struct {
	ID                        uint    `gorm:"column:id_curso_pionero;primaryKey" json:"id_curso_pionero"`
	IdCompAsig                uint    `json:"id_comp_asig"`
	IdNivelComp               uint    `json:"id_nivel_comp"`
	CodTipoModulo             uint    `json:"cod_tipo_modulo"`
	SctTotal                  float64 `json:"sct_total"`
	HrsCronoPresenSemestral   float64 `json:"hrs_crono_presen_semestral"`
	HrsCronoNoPresenSemestral float64 `json:"hrs_crono_no_presen_semestral"`
	HrsTeoricas               float64 `json:"hrs_teoricas"`
	HrsEjercicios             float64 `json:"hrs_ejercicios"`
	HrsPracticas              float64 `json:"hrs_practicas"`
	Institucional             uint    `json:"institucional"`
	EsTEL                     uint    `json:"es_TEL"`

	CompetenciaAsignatura      CompetenciaAsignatura     `gorm:"foreignKey:IdCompAsig"`
	ModsubsAsigcoms            []ModsubAsigcom           `gorm:"foreignKey:id_curso_pionero"`
	CursosPionerosCompetencias []CursoPioneroCompetencia `gorm:"foreignKey:id_curso_pionero"`
	TelsScts                   []TelSct                  `gorm:"foreignKey:SCT_P"`
	ModulosSubcompetencias     []ModuloSubcompetencia    `gorm:"many2many:modsub_asigcom;foreignKey:ID;joinForeignKey:IdCursoPionero;References:ID;joinReferences:IdModsub"`
	//TO DO: relacion ID NIVEL COMP
}

type CursosPioneros []CursoPionero

// TableName overrides the table name used by User to `profiles`
func (CursoPionero) TableName() string {
	return "cursos_pionero"
}
