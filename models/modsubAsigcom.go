package models

type ModsubAsigcom struct {
	IdCursoPionero uint `gorm:"column:id_curso_pionero;primaryKey",json:"id_curso_pionero"`
	IdModsub       uint `gorm:"column:id_modsub;primaryKey",json:"id_modsub"`

	CursoPionero CursoPionero `gorm:"foreignKey:IdCursoPionero"`
	//TO DO: relacion Modsub
}

type ModsubsAsigcoms []ModsubAsigcom

// TableName overrides the table name used by User to `profiles`
func (ModsubAsigcom) TableName() string {
	return "modsub_asigcom"
}
