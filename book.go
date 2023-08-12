package epub

import (
	"archive/zip"
	"encoding/xml"
	"fmt"
	"io"
	"path"
)

type Book struct {
	Opf       Package
	Container Container
	ZipReader *zip.Reader
}

func OpenBook(reader *zip.Reader) (*Book, error) {
	book := &Book{ZipReader: reader}
	err := book.readXML("META-INF/container.xml", &book.Container)
	if err != nil {
		return nil, err
	}
	err = book.readXML(book.Container.Rootfile.Path, &book.Opf)
	if err != nil {
		return nil, err
	}
	if book.Opf.Version != "3.0" {
		return nil, fmt.Errorf("%s not supported yet!", book.Opf.Version)
	}
	return book, nil
}

func (book *Book) Open(fileName string) (io.ReadCloser, error) {
	return book.open(book.getFileFromRootPath(fileName))
}

func (book *Book) getFileFromRootPath(fileName string) string {
	return path.Join(path.Dir(book.Container.Rootfile.Path), fileName)
}

func (book *Book) readXML(fileName string, targetStruct interface{}) error {
	reader, err := book.open(fileName)
	if err != nil {
		return err
	}
	defer reader.Close()
	dec := xml.NewDecoder(reader)
	return dec.Decode(targetStruct)
}

func (book *Book) open(fileName string) (io.ReadCloser, error) {
	for _, file := range book.ZipReader.File {
		if file.Name == fileName {
			return file.Open()
		}
	}
	return nil, fmt.Errorf("file %s not exist", fileName)
}
