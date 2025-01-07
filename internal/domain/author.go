package domain

type Author struct {
	Base
	Name string `gorm:"type:varchar(100);not null"`
}
