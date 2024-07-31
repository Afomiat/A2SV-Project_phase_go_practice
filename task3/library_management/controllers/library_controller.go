package controllers

import (
    "fmt"
    "library_management/models"
    "library_management/services"
)

type LibraryController struct {
    service services.LibraryManager
}

func NewLibraryController(service services.LibraryManager) *LibraryController {
    return &LibraryController{service: service}
}

func (lc *LibraryController) AddBook(id int, title string, author string) {
    book := models.Book{ID: id, Title: title, Author: author, Status: "Available"}
    lc.service.AddBook(book)
    fmt.Println("Book added:", book)
}

func (lc *LibraryController) RemoveBook(id int) {
    err := lc.service.RemoveBook(id)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("Book removed with ID:", id)
}

func (lc *LibraryController) BorrowBook(bookID int, memberID int) {
    err := lc.service.BorrowBook(bookID, memberID)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("Book borrowed with ID:", bookID, "by member ID:", memberID)
}

func (lc *LibraryController) ReturnBook(bookID int, memberID int) {
    err := lc.service.ReturnBook(bookID, memberID)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("Book returned with ID:", bookID, "by member ID:", memberID)
}

func (lc *LibraryController) ListAvailableBooks() {
    books := lc.service.ListAvailableBooks()
    fmt.Println("Available books:")
    for _, book := range books {
        fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
    }
}

func (lc *LibraryController) ListBorrowedBooks(memberID int) {
    books := lc.service.ListBorrowedBooks(memberID)
    fmt.Println("Borrowed books for member ID:", memberID)
    for _, book := range books {
        fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
    }
}
