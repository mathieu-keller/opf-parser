package epub

import (
	"archive/zip"
	"fmt"
	"strconv"

	"github.com/mathieu-keller/epub-parser/model"
	"github.com/mathieu-keller/epub-parser/epub_v2"
	"github.com/mathieu-keller/epub-parser/epub_v3"
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
		epub_v3.ParseOpf(book)
	case ebookVersion >= 2.0 && ebookVersion < 3.0:
		epub_v2.ParseOpf(book)
	default:
		return nil, fmt.Errorf("%f not supported yet!", ebookVersion)
	}
	return book, nil
}
