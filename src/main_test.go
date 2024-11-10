package epub

import (
	"archive/zip"
	"bytes"
	"os"
	"testing"
	"strconv"
	"github.com/mathieu-keller/epub-parser/model"
)

func Test_parse_epub_2_0_opf(t *testing.T) {
	binaryFile, err := os.ReadFile("./test_epub_v2.0.epub")
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
	assertMetadata(t, book.Metadata)
}

func assertMetadata(t *testing.T, metaData model.Metadata) {
	assertEquals("mainId.Id", t, metaData.MainId.Id, "04f24751-f869-48a4-9100-7a2858f94b47")
	assertEquals("mainId.Scheme", t, metaData.MainId.Scheme, "uuid")

	titles :=  *metaData.Titles;
	assertSize("title size", t, len(titles), 1)
	assertEquals("title.Title", t, titles[0].Title, "Test epub")
	assertEquals("title.Language", t, titles[0].Language, "en")
	assertEquals("title.Type", t, titles[0].Type, "main")
	assertEquals("title.FileAs", t, titles[0].FileAs, "Test epub")
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