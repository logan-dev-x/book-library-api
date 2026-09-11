package book

type Repository interface {
	Save(book CreateBookInput) (Book, error)
	GetAll() []Book
	GetByID(id int) (Book, error)
	Update(id int, book UpdateBookInput) error
	Delete(id int) error
}
