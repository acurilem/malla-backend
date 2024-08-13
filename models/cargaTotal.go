package models

type CargaTotal struct {
	ID                      uint                 `gorm:"column:id_carga_total;primaryKey" json:"id_carga_total"`
	NomGrupo                string               `json:"nom_grupo"`
	CantidadEstimadaAlum    uint                 `json:"cantidad_estimada_alum"`
	NumGrupos               uint                 `json:"num_grupos"`
	Hcp                     uint                 `json:"hcp"`
	IdCargaParcial          uint                 `json:"id_carga_parcial"`
	CodTipoSeccion          uint                 `json:"cod_tipo_seccion"`
	CargaParcial            CargaParcial         `gorm:"foreignKey:IdCargaParcial"`
	CargasTotalesProfesores []CargaTotalProfesor `gorm:"foreignKey:id_carga_total"`
	UsuariosAulas           []UsuarioAula        `gorm:"foreignKey:id_carga_total"`
	TipoSeccion             TipoSeccion          `gorm:"foreignKey:CodTipoSeccion"`
}

type CargasTotales []CargaTotal

// TableName overrides the table name used by User to `profiles`
func (CargaTotal) TableName() string {
	return "carga_total"
}
