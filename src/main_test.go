package epub

import (
	"archive/zip"
	"bytes"
	"github.com/mathieu-keller/epub-parser/model"
	"os"
	"strconv"
	"testing"
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

func Test_parse_epub_3_0_opf(t *testing.T) {
	binaryFile, err := os.ReadFile("./test_epub_v3.0.epub")
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

	titles := *metaData.Titles
	assertSize("titles size", t, len(titles), 1)
	assertEquals("title.Title", t, titles[0].Title, "Test epub")
	assertEquals("title.Language", t, titles[0].Language, "en")
	assertEquals("title.Type", t, titles[0].Type, "main")
	assertEquals("title.FileAs", t, titles[0].FileAs, "Test epub")

	identifiers := *metaData.Identifiers
	assertSize("identifiers size", t, len(identifiers), 3)
	assertEquals("identifier[0].Id", t, identifiers[0].Id, "04f24751-f869-48a4-9100-7a2858f94b47")
	assertEquals("identifier[0].Scheme", t, identifiers[0].Scheme, "uuid")
	assertEquals("identifier[1].Id", t, identifiers[1].Id, "afaec86f-3684-4802-9f5b-df8df60e1e6a")
	assertEquals("identifier[1].Scheme", t, identifiers[1].Scheme, "calibre")
	assertEquals("identifier[2].Id", t, identifiers[2].Id, "1529044197")
	assertEquals("identifier[2].Scheme", t, identifiers[2].Scheme, "ISBN")

	languages := *metaData.Languages
	assertSize("languages size", t, len(languages), 1)
	assertEquals("language", t, languages[0], "en")

	creators := *metaData.Creators
	assertSize("creators size", t, len(creators), 1)
	assertEquals("creators.FileAs", t, creators[0].FileAs, "Doe, John")
	assertEquals("creators.Language", t, creators[0].Language, "")
	assertEquals("creators.Name", t, creators[0].Name, "John, Doe")
	assertEquals("creators.Role", t, creators[0].Role, "author")
	assertEquals("creators.RawRole", t, creators[0].RawRole, "aut")

	contributors := *metaData.Contributors
	assertSize("contributors size", t, len(contributors), 1)
	assertEquals("contributors.FileAs", t, contributors[0].FileAs, "")
	assertEquals("contributors.Language", t, contributors[0].Language, "")
	assertEquals("contributors.Name", t, contributors[0].Name, "GoLang")
	assertEquals("contributors.Role", t, contributors[0].Role, "book producer")
	assertEquals("contributors.RawRole", t, contributors[0].RawRole, "bkp")

	publishers := *metaData.Publishers
	assertSize("publishers size", t, len(publishers), 1)
	assertEquals("publishers.Text", t, publishers[0].Text, "Test Publisher")
	assertEquals("publishers.Language", t, publishers[0].Language, "")

	subjects := *metaData.Subjects
	assertSize("subjects size", t, len(subjects), 2)
	assertEquals("subjects[0].Text", t, subjects[0].Text, "Novel")
	assertEquals("subjects[0].Language", t, subjects[0].Language, "")
	assertEquals("subjects[1].Text", t, subjects[1].Text, "Comic science fiction")
	assertEquals("subjects[1].Language", t, subjects[1].Language, "")

	descriptions := *metaData.Descriptions
	assertSize("descriptions size", t, len(descriptions), 1)
	assertEquals("descriptions.Text", t, descriptions[0].Text, `an ordinary Earthling suddenly swept into a wild space adventure when his planet is slated for demolition. With nothing but a towel and a knack for making friends in strange places, Finn hitches rides across the galaxy alongside a quirky crew: a chronically pessimistic robot, a free-spirited alien guide, and a ship with an attitude.

From cosmic dive bars to worlds with surreal rules, Finn navigates the chaos in search of purpose—or, at least, survival. Along the way, he learns one universal truth: when hitchhiking through the stars, never lose track of your towel.`)
	assertEquals("descriptions.Language", t, descriptions[0].Language, "en")

	dates := *metaData.Dates
	assertSize("dates size", t, len(dates), 1)
	assertEquals("dates[0]", t, dates[0], "2024-11-10T22:00:00Z")
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
