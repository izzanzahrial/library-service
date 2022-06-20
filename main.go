package main

import (
	"github.com/izzanzahrial/library-service/entity"
	"github.com/izzanzahrial/library-service/service"
)

var userPool = make(map[string]*entity.User)
var booksPool = make(map[string]*entity.Book)
var borrowPool = make(map[int]*entity.Borrow)

func main() {
	userService := service.NewUser(userPool)
	libraryService := service.NewLibrary(booksPool)
	borrowService := service.NewBorrowing(userService, libraryService, borrowPool)
}
