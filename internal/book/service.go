package book

import "errors"

type Service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return Service{repo: repo}
}

func (s Service) Create(book CreateBookInput) (Book, error) {
	created, err := s.repo.Save(book)
	if err != nil {
		return Book{}, err
	}
	return created, nil
}

func (s Service) GetAll() []Book {
	return s.repo.GetAll()
}

func (s Service) GetOne(id int) (Book, error) {
	book, err := s.repo.GetByID(id)
	if err != nil || book == (Book{}) {
		return Book{}, errors.New("book not found")
	}
	return book, nil
}

func (s Service) Update(id int, book UpdateBookInput) error {
	b, err := s.repo.GetByID(id)
	if err != nil && b == (Book{}) {
		return errors.New("book not found")
	}

	return s.repo.Update(b.ID, book)
}

func (s Service) Delete(id int) error {
	b, err := s.GetOne(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(b.ID)
}
