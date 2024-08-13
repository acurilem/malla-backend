package models

type EstadoInscripcion struct {
	ID                   uint   `gorm:"column:cod_estado_inscripcion;primaryKey",json:"cod_estado_inscripcion"`
	NomEstadoInscripcion string `json:"nom_estado_inscripcion"`
}

// TableName overrides the table name used by User to `profiles`
func (EstadoInscripcion) TableName() string {
	return "estado_inscripcion"
}
