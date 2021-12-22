package User

type User struct {
	Name     string `gorm:"type:varchar(30);not null"`
	Password string `gorm:"size:255;not null"`
}
