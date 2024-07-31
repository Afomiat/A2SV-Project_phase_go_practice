package services

import(
	"errors"
	"library_management/models"
) 


type LibraryManager interface {
    AddBook(book models.Book)
    RemoveBook(bookID int) error
    BorrowBook(bookID int, memberID int) error
    ReturnBook(bookID int, memberID int) error
    ListAvailableBooks() []models.Book
    ListBorrowedBooks(memberID int) []models.Book
}

type Library struct {
    books   map[int]models.Book
    members map[int]models.Member
}

func NewLibrary() *Library {
    return &Library{
        books:   make(map[int]models.Book),
        members: make(map[int]models.Member),
    }
}

func (lib *Library) AddBook(book models.Book) {
    lib.books[book.ID] = book
}

func (lib *Library) RemoveBook(bookID int) error {
    if _, exists := lib.books[bookID]; !exists {
        return errors.New("book not found")
    }
    delete(lib.books, bookID)
    return nil
}

func (lib *Library) BorrowBook(bookID int, memberID int) error {
    book, bookExists := lib.books[bookID]
    if !bookExists {
        return errors.New("book not found")
    }
    if book.Status == "Borrowed" {
        return errors.New("book already borrowed")
    }
    member, memberExists := lib.members[memberID]
    if !memberExists {
        return errors.New("member not found")
    }
    book.Status = "Borrowed"
    lib.books[bookID] = book
    member.BorrowedBooks = append(member.BorrowedBooks, book)
    lib.members[memberID] = member
    return nil
}

func (lib *Library) ReturnBook(bookID int, memberID int) error {
    book, bookExists := lib.books[bookID]
    if !bookExists {
        return errors.New("book not found")
    }
    member, memberExists := lib.members[memberID]
    if !memberExists {
        return errors.New("member not found")
    }
    book.Status = "Available"
    lib.books[bookID] = book

    for i, b := range member.BorrowedBooks {
        if b.ID == bookID {
            member.BorrowedBooks = append(member.BorrowedBooks[:i], member.BorrowedBooks[i+1:]...)
            break
        }
    }
    lib.members[memberID] = member
    return nil
}

func (lib *Library) ListAvailableBooks() []models.Book {
    availableBooks := []models.Book{}
    for _, book := range lib.books {
        if book.Status == "Available" {
            availableBooks = append(availableBooks, book)
        }
    }
    return availableBooks
}

func (lib *Library) ListBorrowedBooks(memberID int) []models.Book {
    member, exists := lib.members[memberID]
    if !exists {
        return nil
    }
    return member.BorrowedBooks
}
