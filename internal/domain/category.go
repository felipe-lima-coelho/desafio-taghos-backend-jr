package domain

type Category struct {
	Base
	Name string `gorm:"type:varchar(100);unique;not null"`
}
