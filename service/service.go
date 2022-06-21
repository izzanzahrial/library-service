package service

import "github.com/izzanzahrial/library-service/entity"

type User interface {
	RegisterUser(name string) error
	GetUser(userID int) (*entity.User, error)
}

func NewUser(users map[int]*entity.User) User {
	return &user{
		users: users,
	}
}

type Library interface {
	RegisterBook(title, author string, qty int) error
	ChangeBookQty(title string, qty int) error
	GetBook(title string) (*entity.BookQty, error)
	ReturnBook(title string) error
}

func NewLibrary(books map[string]*entity.BookQty) Library {
	return &library{
		books: books,
	}
}

type Borrowing interface {
	BorrowBooks(titles []string, userID int) error
	ReturnBooks(titles []string, userID int) error
}

func NewBorrowing(user User, library Library, borrowCache map[int]*entity.Borrow) Borrowing {
	return &borrowing{
		user:      user,
		library:   library,
		borrowing: borrowCache,
	}
}
