package domain

type Book struct {
	Base
	Title      string     `gorm:"type:varchar(255);not null"`
	Synopsis   string     `gorm:"type:text;not null"`
	Categories []Category `gorm:"many2many:book_categories;"`
	Authors    []Author   `gorm:"many2many:book_authors;"`
}
