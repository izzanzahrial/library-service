package service

import "github.com/izzanzahrial/library-service/entity"

type User interface {
	RegisterUser(name, role string) error
	GetUser(name string) (*entity.User, error)
}

func NewUser(userPool map[string]*entity.User) User {
	return &user{
		users: userPool,
	}
}

type Library interface {
	RegisterBook(title, author string, qty int) error
	ChangeBookQty(title string, qty int) error
	GetBook(title string) (*entity.Book, error)
	ReturnBook(title string) error
}

func NewLibrary(bookPool map[string]*entity.Book) Library {
	return &library{
		books: bookPool,
	}
}

type Borrowing interface {
	BorrowBooks(title []string, user *entity.User) error
	ReturnBooks(books []*entity.Book, user *entity.User) error
}

func NewBorrowing(user User, library Library, borrowPool map[int]*entity.Borrow) Borrowing {
	return &borrowing{
		user:      user,
		library:   library,
		borrowing: borrowPool,
	}
}
