package models

type TelSct struct {
	TEL_P uint `gorm:"column:TEL_P;primaryKey",json:"TEL_P"`
	SCT_P uint `gorm:"column:SCT_P",json:"SCT_P"`
}

type TelsScts []TelSct

// TableName overrides the table name used by User to `profiles`
func (TelSct) TableName() string {
	return "TEL_SCT"
}
