package entities

type TodoWithAuthor struct {
	Todo
	AuthorName string `gorm:"name:author_name"`
}
type TodosWithAuthor []TodoWithAuthor
