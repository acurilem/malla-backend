package models

type MallaRequisito struct {
	ID         uint `gorm:"column:id_requisito;primaryKey",json:"id_requisito"`
	IdMalla    uint `json:"id_malla"`
	IdMallaReq uint `json:"id_malla_req"`
	Tipo       uint `json:"tipo"`

	MallaPionero MallaPionero `gorm:"foreignKey:IdMallaReq"`
	//TO DO: Relacion id_malla
}

type MallasRequisitos []MallaRequisito

// TableName overrides the table name used by User to `profiles`
func (MallaRequisito) TableName() string {
	return "malla_requisitos"
}
