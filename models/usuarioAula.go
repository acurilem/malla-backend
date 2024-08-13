package models

type UsuarioAula struct {
	ID                   uint              `gorm:"column:id_usuario_aula;primaryKey" json:"id_usuario_aula"`
	CodPersona           uint              `json:"cod_persona"`
	IdCargaTotal         uint              `json:"id_carga_total"`
	Sem                  uint              `json:"sem"`
	Ano                  uint              `json:"ano"`
	IdTipoUsuario        uint              `json:"id_tipo_usuario"`
	CodEstadoInscripcion uint              `json:"cod_estado_inscripcion"`
	Opcion               uint              `json:"opcion"`
	IdMalla              uint              `json:"id_malla"`
	IdMallaCd            uint              `json:"id_malla_cd"`
	EstadoInscripcion    EstadoInscripcion `gorm:"foreignKey:CodEstadoInscripcion"`
	MallaPionero         MallaPionero      `gorm:"foreignKey:IdMalla"`
	TipoUsuario          TipoUsuario       `gorm:"foreignKey:IdTipoUsuario"`
}

type UsuariosAulas []UsuarioAula

// TableName overrides the table name used by User to `profiles`
func (UsuarioAula) TableName() string {
	return "usuarios_aula"
}
