package epub

import (
	"archive/zip"
	"bytes"
	"os"
	"testing"
	"strconv"
)

func Test_parse_epub_3_0_opf(t *testing.T) {
	binaryFile, err := os.ReadFile("./test_epub_v3_0.epub")
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	zipReader, err := zip.NewReader(bytes.NewReader(binaryFile), int64(len(binaryFile)))
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}

	book, err := OpenBook(zipReader)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	metadata := book.Metadata

	assertSize("mainId", t, len(*metadata.Creators), 1)
	assertEquals("mainId", t, (*metadata.Creators)[0].Role, "author")
}

func assertSize(fieldName string, t *testing.T, actuallySize int, expectedSize int) {
	if actuallySize != expectedSize {
		t.Logf("'%s' length expected '%s' but is '%s'", fieldName, strconv.Itoa(expectedSize), strconv.Itoa(actuallySize))
		t.Fail()
	}
}

func assertEquals(fieldName string, t *testing.T, actuallyValue string, expectedValue string) {
	if actuallyValue != expectedValue {
		t.Logf("'%s' expected '%s' but is '%s'", fieldName, expectedValue, actuallyValue)
		t.Fail()
	}
}