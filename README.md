# epub-parser

`epub-parser` is a Go library for parsing EPUB files, specifically versions 2.0 and 3.0. It extracts and interprets metadata from the EPUB's `OPF` (Open Packaging Format) files, enabling developers to programmatically access information such as titles, authors, publishers, and more.

This library is particularly useful for applications requiring detailed metadata extraction from EPUB files, such as e-book management tools, cataloging systems, or digital libraries.

---

## Features

- **Supports EPUB 2.0 and 3.0**: Parses both versions seamlessly.
- **Metadata Extraction**:
    - Titles
    - Identifiers (e.g., ISBN, UUID)
    - Languages
    - Creators (Authors)
    - Contributors
    - Publishers
    - Subjects
    - Descriptions
    - Dates
- **ZIP-based EPUB Parsing**: Reads EPUB files directly from ZIP archives.

---

## Installation

Add the library to your project using `go get`:

```sh
go get github.com/mathieu-keller/epub-parser
```

---

## Usage

Here’s an example test demonstrating how to use the library:

```go
package epub

import (
	"archive/zip"
	"bytes"
	"fmt"
	"os"
)

func ParseEPUB() {
	// Load the EPUB file
	binaryFile, err := os.ReadFile("./test_epub_v3.0.epub")
	if err != nil {
		fmt.Printf("Failed to read EPUB file: %v", err)
		os.Exit(1)
	}

	// Create a ZIP reader for the EPUB
	zipReader, err := zip.NewReader(bytes.NewReader(binaryFile), int64(len(binaryFile)))
	if err != nil {
		fmt.Printf("Failed to create ZIP reader: %v", err)
		os.Exit(1)
	}

	// Parse the book
	book, err := OpenBook(zipReader)
	if err != nil {
		fmt.Printf("Failed to parse EPUB book: %v", err)
		os.Exit(1)
	}
	fmt.Println(book.Metadata.MainId.Id)
}
```

---

## Model: Metadata

The core of the library is the `Metadata` struct, which encapsulates the detailed metadata of an EPUB file. Here’s the structure and its key components:


```go
type Metadata struct {
MainId       Identifier           // Main identifier of the EPUB (e.g., UUID)
Titles       *[]Title             // List of titles
Identifiers  *[]Identifier        // List of identifiers (e.g., UUID, ISBN, etc.)
Languages    *[]string            // List of languages
Creators     *[]Creator           // List of creators (e.g., authors)
Contributors *[]Creator           // List of contributors (e.g., editors, producers)
Publishers   *[]DefaultAttributes // List of publishers
Subjects     *[]DefaultAttributes // List of subjects (categories, genres)
Descriptions *[]DefaultAttributes // List of descriptions
Dates        *[]string            // List of publication dates
}
```

### Supporting Types

- **`Title`**: Represents a title in the EPUB.
  ```go
  type Title struct {
      Title    string
      Language string
      Type     string
      FileAs   string
  }
  ```

- **`Identifier`**: Represents an identifier like UUID or ISBN.
  ```go
  type Identifier struct {
      Id     string
      Scheme string
  }
  ```

- **`Creator`**: Represents an author or contributor.
  ```go
  type Creator struct {
      Name     string
      Language string
      FileAs   string
      Role     string
      RawRole  string
  }
  ```

- **`DefaultAttributes`**: Generic type for attributes like publishers, subjects, and descriptions.
  ```go
  type DefaultAttributes struct {
      Text     string
      Language string
  }
  ```

### Example Metadata Output:

- **Title**: Test epub
- **Language**: en
- **Creators**:
    - Name: John Doe
    - Role: Author
- **Publisher**: Test Publisher
- **Subjects**: Novel, Comic Science Fiction
- **Description**: A captivating space adventure...

---

## Contributing

1. Fork the repository.
2. Create a feature branch (`git checkout -b feature-name`).
3. Commit your changes (`git commit -m "Add feature"`).
4. Push the branch (`git push origin feature-name`).
5. Open a pull request.

---

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

Happy parsing! 🚀
