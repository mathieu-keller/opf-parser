package model

import (
	"archive/zip"
	"encoding/xml"
	"fmt"
	"io"
	"path"
)

type Book struct {
	Metadata  Metadata
	Container Container
	ZipReader *zip.Reader
}

func (book *Book) Open(fileName string) (io.ReadCloser, error) {
	return book.open(book.getFileFromRootPath(fileName))
}

func (book *Book) getFileFromRootPath(fileName string) string {
	return path.Join(path.Dir(book.Container.Rootfile.Path), fileName)
}

func (book *Book) ReadXML(fileName string, targetStruct interface{}) error {
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
