package data

type Book struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
}

var Books = []*Book{
	{Id: 1, Name: "Dune", Author: "Frank Herbert"},
	{Id: 2, Name: "Fire & Blood", Author: "George RR Martin"},
}

func FindBook(id int) *Book {
	id = id - 1
	for i, b := range Books {
		if id == i {
			return b
		}
	}

	return nil
}

func AddBook(book *Book) {
	book.Id = getNextID()
	Books = append(Books, book)
}

func getNextID() int {
	lp := Books[len(Books)-1]
	return lp.Id + 1
}
