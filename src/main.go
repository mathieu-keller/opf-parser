package epub

import (
	"archive/zip"
	"fmt"
	"strconv"

	"github.com/mathieu-keller/epub-parser/model"
	v2 "github.com/mathieu-keller/epub-parser/v2"
	v3 "github.com/mathieu-keller/epub-parser/v3"
)

func OpenBook(reader *zip.Reader) (*model.Book, error) {
	book := &model.Book{ZipReader: reader}
	err := book.ReadXML("META-INF/container.xml", &book.Container)
	if err != nil {
		return nil, err
	}
	header := model.Package{}
	err = book.ReadXML(book.Container.Rootfile.Path, &header)
	if err != nil {
		return nil, err
	}
	ebookVersion, err := strconv.ParseFloat(header.Version, 64)
	if err != nil {
		return nil, err
	}
	switch {
	case ebookVersion >= 3.0 && ebookVersion < 4.0:
		v3.ParseOpf(book)
	case ebookVersion >= 2.0 && ebookVersion < 3.0:
		v2.ParseOpf(book)
	default:
		return nil, fmt.Errorf("%f not supported yet!", ebookVersion)
	}
	return book, nil
}
