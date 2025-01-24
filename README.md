# Library Management System

A simple command-line application written in Go for managing a library of books and eBooks. This project demonstrates the use of interfaces, structs, and basic input/output operations in Go.

## Features
- Add a physical book to the library.
- Add an eBook to the library with file size details.
- Remove a book or eBook from the library using its ISBN.
- Search for books or eBooks by their title.
- List all books and eBooks currently in the library.

## How to Run
1. Make sure you have [Go](https://golang.org/) installed on your system.
2. Clone this repository or copy the source code.
3. Navigate to the project directory in your terminal.
4. Run the following command to execute the program:
   ```bash
   go run main.go
   ```

## How to Use
1. **Menu Options**:
   - `1`: Add a physical book by entering the title, author, and ISBN.
   - `2`: Add an eBook by entering the title, author, ISBN, and file size (in MB).
   - `3`: Remove a book by entering its ISBN.
   - `4`: Search for books by entering part or all of the title.
   - `5`: List all books and eBooks in the library.
   - `6`: Exit the program.

2. **User Input**:
   - Enter multi-word titles, authors, or ISBNs seamlessly.
   - Follow the prompts to provide required details.

## Example
### Adding a Book
```plaintext
Enter Title: The Great Gatsby
Enter Author: F. Scott Fitzgerald
Enter ISBN: 9780743273565
```

### Adding an eBook
```plaintext
Enter Title: Go Programming
Enter Author: John Doe
Enter ISBN: 1234567890
Enter File Size (MB): 5
```

### Removing a Book
```plaintext
Enter ISBN to remove: 9780743273565
```

### Listing All Books
```plaintext
Title: The Great Gatsby
Author: F. Scott Fitzgerald
ISBN: 9780743273565
Available: true
---
Title: Go Programming
Author: John Doe
ISBN: 1234567890
Available: true
File Size: 5 MB
---
```

## Screenshots
