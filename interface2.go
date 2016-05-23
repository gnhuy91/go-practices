package main

import "fmt"

func main() {
	fmt.Printf("%+v\n", new(Item)) // &{id:0 deleted:false by:}
	fmt.Printf("%+v\n", &Item{2, true, "me"})

	item := &Item{1, false, "Huy"}
	intf := Interface(item)

	intfs := []Interface{item}

	fmt.Printf("%+v\n", intf)     // &{id:0 deleted:false by:}
	fmt.Printf("%+v\n", intfs[0]) // &{id:0 deleted:false by:}

	fmt.Println(DoItem(item))
	book := &Book{2, "John"}
	fmt.Println(DoItem(book))
}

type Interface interface {
	Type() string
	Maker() string
}

type Item struct {
	id      int
	deleted bool
	by      string
}

func (it *Item) Type() string {
	return "item"
}

func (it *Item) Maker() string {
	return it.by
}

func DoItem(data Interface) string {
	return fmt.Sprintf("Do things with %v by %v", data.Type(), data.Maker())
}

type Book struct {
	id int
	by string
}

func (book *Book) Type() string {
	return "book"
}

func (book *Book) Maker() string {
	return book.by
}
