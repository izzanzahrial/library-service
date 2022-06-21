package entity

type Book struct {
	ID     int
	Title  string
	Author string
}

type BookQty struct {
	Book
	Qty int
}
