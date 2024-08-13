package models

type TipoUsuario struct {
	ID             uint   `gorm:"column:id_tipo_usuario;primaryKey",json:"id_tipo_usuario"`
	NomTipoUsuario string `json:"nom_tipo_usuario"`
}

// TableName overrides the table name used by User to `profiles`
func (TipoUsuario) TableName() string {
	return "tipo_usuario"
}
