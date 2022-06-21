package main

import (
	"fmt"

	"github.com/izzanzahrial/library-service/entity"
	"github.com/izzanzahrial/library-service/service"
)

var users = make(map[int]*entity.User)
var books = make(map[string]*entity.BookQty)
var borrowCache = make(map[int]*entity.Borrow)

func main() {
	userService := service.NewUser(users)
	libraryService := service.NewLibrary(books)
	borrowService := service.NewBorrowing(userService, libraryService, borrowCache)

	for {
		var command string
		fmt.Println("state your command: ")
		scan(&command)

		switch command {
		case "register_user":
			var name string
			fmt.Println("your name: ")
			scan(&name)
			if err := userService.RegisterUser(name); err != nil {
				panic(err)
			}
			fmt.Println("register user success")
		// case "get_user":
		// 	var userID int
		// 	fmt.Println("user ID: ")
		// 	scan(&userID)
		// 	user, err := userService.GetUser(userID)
		// 	if err != nil {
		// 		panic(err)
		// 	}
		// 	fmt.Println(user)
		case "register_book":
			var title, author string
			var qty int
			fmt.Println("book title: ")
			scan(&title)
			fmt.Println("book author: ")
			scan(&author)
			fmt.Println("book quantity: ")
			scan(&qty)
			if err := libraryService.RegisterBook(title, author, qty); err != nil {
				panic(err)
			}
			fmt.Println("register book success")
		case "change_book_qty":
			var title string
			var qty int
			fmt.Println("book title: ")
			scan(&title)
			fmt.Println("book quantity: ")
			scan(&qty)
			if err := libraryService.ChangeBookQty(title, qty); err != nil {
				panic(err)
			}
			fmt.Println("change book quantity success")
		// case "get_book":
		// 	var title string
		// 	fmt.Println("book title: ")
		// 	scan(&title)
		// 	book, err := libraryService.GetBook(title)
		// 	if err != nil {
		// 		panic(err)
		// 	}
		// 	fmt.Println(book)
		// case "return_book":
		// 	var title string
		// 	fmt.Println("book title: ")
		// 	scan(&title)
		// 	if err := libraryService.ReturnBook(title); err != nil {
		// 		panic(err)
		// 	}
		// 	fmt.Println("return book success")
		case "borrow_books":
			var titles []string
			var userID int

			fmt.Println("What is your ID: ")
			scan(&userID)

			for {
				var title string
				var next string

				fmt.Println("book title: ")
				scan(&title)
				titles = append(titles, title)
				fmt.Println("do you want to borrow another book? (yes/no) ")
				scan(&next)

				if next == "no" {
					break
				}
			}

			if err := borrowService.BorrowBooks(titles, userID); err != nil {
				panic(err)
			}
		case "return_books":
			var titles []string
			var userID int

			fmt.Println("What is your ID: ")
			scan(&userID)

			for {
				var title string
				var next string

				fmt.Println("book title: ")
				scan(&title)
				titles = append(titles, title)
				fmt.Println("do you want to return another book? (yes/no) ")
				scan(&next)

				if next == "no" {
					break
				}
			}

			if err := borrowService.ReturnBooks(titles, userID); err != nil {
				panic(err)
			}
		}
	}
}

func scan(arg interface{}) {
	_, err := fmt.Scanln(&arg)
	if err != nil && err.Error() != "unexpected newline" {
		panic(err)
	}
}
