package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type BookInterface interface {
	DisplayDetails() string
}

type Book struct {
	Title     string
	Author    string
	ISBN      string
	Available bool
}

func NewBook(title, author, isbn string) Book {
	return Book{
		Title:     title,
		Author:    author,
		ISBN:      isbn,
		Available: true,
	}
}

func (b Book) DisplayDetails() string {
	return fmt.Sprintf("Title: %s\nAuthor: %s\nISBN: %s\nAvailable: %t",
		b.Title, b.Author, b.ISBN, b.Available)
}

type EBook struct {
	Book
	FileSize int
}

func NewEBook(title, author, isbn string, fileSize int) EBook {
	return EBook{
		Book:     NewBook(title, author, isbn),
		FileSize: fileSize,
	}
}

func (e EBook) DisplayDetails() string {
	return fmt.Sprintf("%s\nFile Size: %d MB",
		e.Book.DisplayDetails(), e.FileSize)
}

type Library struct {
	Books []BookInterface
}

func (l *Library) AddBook(book BookInterface) {
	l.Books = append(l.Books, book)
}

func (l *Library) RemoveBook(isbn string) bool {
	for i, book := range l.Books {
		var b Book
		switch v := book.(type) {
		case Book:
			b = v
		case EBook:
			b = v.Book
		}

		if b.ISBN == isbn {
			l.Books = append(l.Books[:i], l.Books[i+1:]...)
			return true
		}
	}
	return false
}

func (l *Library) SearchBookByTitle(title string) []BookInterface {
	var results []BookInterface
	for _, book := range l.Books {
		var b Book
		switch v := book.(type) {
		case Book:
			b = v
		case EBook:
			b = v.Book
		}

		if strings.Contains(strings.ToLower(b.Title), strings.ToLower(title)) {
			results = append(results, book)
		}
	}
	return results
}

func (l *Library) ListBooks() {
	if len(l.Books) == 0 {
		fmt.Println("No books in the library.")
		return
	}

	for _, book := range l.Books {
		fmt.Println(book.DisplayDetails())
		fmt.Println("---")
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	library := &Library{}

	for {
		fmt.Println("\nLibrary Management System")
		fmt.Println("1. Add Book")
		fmt.Println("2. Add EBook")
		fmt.Println("3. Remove Book")
		fmt.Println("4. Search Books")
		fmt.Println("5. List All Books")
		fmt.Println("6. Exit")
		fmt.Print("Enter your choice: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Print("Enter Title: ")
			title, _ := reader.ReadString('\n')
			title = strings.TrimSpace(title)

			fmt.Print("Enter Author: ")
			author, _ := reader.ReadString('\n')
			author = strings.TrimSpace(author)

			fmt.Print("Enter ISBN: ")
			isbn, _ := reader.ReadString('\n')
			isbn = strings.TrimSpace(isbn)

			library.AddBook(NewBook(title, author, isbn))

		case 2:
			fmt.Print("Enter Title: ")
			title, _ := reader.ReadString('\n')
			title = strings.TrimSpace(title)

			fmt.Print("Enter Author: ")
			author, _ := reader.ReadString('\n')
			author = strings.TrimSpace(author)

			fmt.Print("Enter ISBN: ")
			isbn, _ := reader.ReadString('\n')
			isbn = strings.TrimSpace(isbn)

			fmt.Print("Enter File Size (MB): ")
			var fileSize int
			fmt.Scanln(&fileSize)

			library.AddBook(NewEBook(title, author, isbn, fileSize))

		case 3:
			fmt.Print("Enter ISBN to remove: ")
			isbn, _ := reader.ReadString('\n')
			isbn = strings.TrimSpace(isbn)
			if library.RemoveBook(isbn) {
				fmt.Println("Book removed successfully.")
			} else {
				fmt.Println("Book not found.")
			}

		case 4:
			fmt.Print("Enter title to search: ")
			title, _ := reader.ReadString('\n')
			title = strings.TrimSpace(title)
			results := library.SearchBookByTitle(title)
			if len(results) > 0 {
				for _, book := range results {
					fmt.Println(book.DisplayDetails())
					fmt.Println("---")
				}
			} else {
				fmt.Println("No books found.")
			}

		case 5:
			library.ListBooks()

		case 6:
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid choice. Try again.")
		}
	}
}