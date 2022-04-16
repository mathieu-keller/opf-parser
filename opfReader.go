package epub

import (
	"encoding/xml"
	"io"
)

func ReadOpf(in io.ReadCloser) (*Package, error) {
	defer in.Close()
	dec := xml.NewDecoder(in)
	var opf Package
	err := dec.Decode(&opf)
	if err != nil {
		return nil, err
	}
	return &opf, nil
}
