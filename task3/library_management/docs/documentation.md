
## Models

### Book

Represents a book in the library.

Fields:
- ID (int)
- Title (string)
- Author (string)
- Status (string)

### Member

Represents a library member.

Fields:
- ID (int)
- Name (string)
- BorrowedBooks ([]Book)

## LibraryManager Interface

Defines the methods for managing the library.

Methods:
- AddBook(book Book)
- RemoveBook(bookID int)
- BorrowBook(bookID int, memberID int) error
- ReturnBook(bookID int, memberID int) error
- ListAvailableBooks() []Book
- ListBorrowedBooks(memberID int) []Book

## Library Service

Implements the LibraryManager interface. Manages books and members using maps.

## Console Interaction

Provides a simple console interface for interacting with the library management system. The following actions are supported:
- Add a new book.
- Remove an existing book.
- Borrow a book.
- Return a book
