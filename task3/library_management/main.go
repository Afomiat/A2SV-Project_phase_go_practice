package main

import (
    "bufio"
    "fmt"
    "library_management/controllers"
    "library_management/services"
    "os"
    "strconv"
)

func main() {
    library := services.NewLibrary()
    controller := controllers.NewLibraryController(library)

    scanner := bufio.NewScanner(os.Stdin)

    for {
        fmt.Println("\nLibrary Management System")
        fmt.Println("1. Add Book")
        fmt.Println("2. Remove Book")
        fmt.Println("3. Borrow Book")
        fmt.Println("4. Return Book")
        fmt.Println("5. List Available Books")
        fmt.Println("6. List Borrowed Books by Member")
        fmt.Println("7. Exit")
        fmt.Print("Enter your choice: ")

        scanner.Scan()
        choice, _ := strconv.Atoi(scanner.Text())

        switch choice {
        case 1:
            fmt.Print("Enter Book ID: ")
            scanner.Scan()
            bookID, _ := strconv.Atoi(scanner.Text())
            fmt.Print("Enter Book Title: ")
            scanner.Scan()
            title := scanner.Text()
            fmt.Print("Enter Book Author: ")
            scanner.Scan()
            author := scanner.Text()
            controller.AddBook(bookID, title, author)
        case 2:
            fmt.Print("Enter Book ID to remove: ")
            scanner.Scan()
            bookID, _ := strconv.Atoi(scanner.Text())
            controller.RemoveBook(bookID)
        case 3:
            fmt.Print("Enter Book ID to borrow: ")
            scanner.Scan()
            bookID, _ := strconv.Atoi(scanner.Text())
            fmt.Print("Enter Member ID: ")
            scanner.Scan()
            memberID, _ := strconv.Atoi(scanner.Text())
            controller.BorrowBook(bookID, memberID)
        case 4:
            fmt.Print("Enter Book ID to return: ")
            scanner.Scan()
            bookID, _ := strconv.Atoi(scanner.Text())
            fmt.Print("Enter Member ID: ")
            scanner.Scan()
            memberID, _ := strconv.Atoi(scanner.Text())
            controller.ReturnBook(bookID, memberID)
        case 5:
            controller.ListAvailableBooks()
        case 6:
            fmt.Print("Enter Member ID to list borrowed books: ")
            scanner.Scan()
            memberID, _ := strconv.Atoi(scanner.Text())
            controller.ListBorrowedBooks(memberID)
        case 7:
            fmt.Println("Exiting...")
            return
        default:
            fmt.Println("Invalid choice, please try again.")
        }
    }
}
