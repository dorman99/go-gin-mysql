package entity

type Book struct {
	ID          uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Title       string `gorm:"type:varchar(255)" json:"title"`
	Description string `gorm:"type:text" json:"description"`
	Deleted     bool   `gorm:"type:boolean; default=false" json:"-"`
	UserID      uint64 `gorm:"not null;" json:"userId"`
	User        User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user"`
}
