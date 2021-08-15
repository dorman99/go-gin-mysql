package entity

type User struct {
	ID       uint64 `gorm:"primarykey_key:auto_increment" json:"id"`
	Name     string `gorm:"type:varchar(255)" json:"name"`
	Username string `gorm:"uniqueIndex;type:varchar(255);not null" json:"username"`
	Password string `gorm:"->;<-;not null" json:"-"`
}
