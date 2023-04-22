package epub

import (
	"archive/zip"
	"strconv"
	"testing"
)

func Test_parse_epub_3_0_opf(t *testing.T) {
	zipReader, err := zip.OpenReader("./test_epub_v3_0.epub")
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	defer zipReader.Close()

	book, err := OpenBook(zipReader)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	opf := book.Opf

	assertEquals("Lang", t, opf.Lang, "us")
	assertEquals("UniqueIdentifier", t, opf.UniqueIdentifier, "p1234")
	assertEquals("Version", t, opf.Version, "3.0")
	assertEquals("ID", t, opf.ID, "test-id")
	assertEquals("Prefix", t, opf.Prefix, "pre")

	testMetaData(t, *opf.Metadata)
	testManifest(t, *opf.Manifest)

}

func testManifest(t *testing.T, manifest Manifest) {
	assertEquals("manifest.Id", t, manifest.Id, "manifest-test")
	testItem(t, *manifest.Item)
}

func testItem(t *testing.T, item []Item) {
	assertSize("item", t, len(item), 1)
	assertEquals("item[0].Id", t, item[0].Id, "cover")
	assertEquals("item[0].Href", t, item[0].Href, "cover.jpeg")
	assertEquals("item[0].MediaType", t, item[0].MediaType, "image/jpeg")
	assertEquals("item[0].MediaOverlay", t, item[0].MediaOverlay, "ch1_audio")
	assertEquals("item[0].Properties", t, item[0].Properties, "cover-image")
	assertEquals("item[0].Fallback", t, item[0].Fallback, "fall1")
}

func testMetaData(t *testing.T, metaData Metadata) {
	testMeta(t, *metaData.Meta)
	testIdentifier(t, *metaData.Identifier)
	testLanguage(t, *metaData.Language)
	testTitle(t, *metaData.Title)
	testDate(t, *metaData.Date)
	testType(t, *metaData.Type)
	testContributor(t, *metaData.Contributor)
	testCoverage(t, *metaData.Coverage)
	testCreator(t, *metaData.Creator)
	testDescription(t, *metaData.Description)
	testPublisher(t, *metaData.Publisher)
	testRelation(t, *metaData.Relation)
	testRights(t, *metaData.Rights)
	testSubject(t, *metaData.Subject)
	testLink(t, *metaData.Link)

}

func testMeta(t *testing.T, meta []Meta) {
	assertSize("meta", t, len(meta), 2)
	assertEquals("meta[0].Dir", t, meta[0].Dir, "ltr")
	assertEquals("meta[0].Id", t, meta[0].Id, "meta-test")
	assertEquals("meta[0].Scheme", t, meta[0].Scheme, "scheme")
	assertEquals("meta[0].Lang", t, meta[0].Lang, "us")
	assertEquals("meta[0].Refines", t, meta[0].Refines, "#t1")
	assertEquals("meta[0].Property", t, meta[0].Property, "title-type")
	assertEquals("meta[0].Text", t, meta[0].Text, "main")
	assertEquals("meta[0].Content", t, meta[0].Content, "")
	assertEquals("meta[0].Name", t, meta[0].Name, "")

	assertEquals("meta[1].Dir", t, meta[1].Dir, "")
	assertEquals("meta[1].Id", t, meta[1].Id, "")
	assertEquals("meta[1].Scheme", t, meta[1].Scheme, "")
	assertEquals("meta[1].Lang", t, meta[1].Lang, "")
	assertEquals("meta[1].Refines", t, meta[1].Refines, "")
	assertEquals("meta[1].Property", t, meta[1].Property, "")
	assertEquals("meta[1].Text", t, meta[1].Text, "")
	assertEquals("meta[1].Content", t, meta[1].Content, "vertical-rl")
	assertEquals("meta[1].Name", t, meta[1].Name, "primary-writing-mode")
}

func testLink(t *testing.T, link []Link) {
	assertSize("link", t, len(link), 1)
	assertEquals("link[0].Id", t, link[0].Id, "link-test")
	assertEquals("link[0].Href", t, link[0].Href, "front.xhtml#meta-json")
	assertEquals("link[0].Rel", t, link[0].Rel, "record")
	assertEquals("link[0].Rel", t, link[0].MediaType, "application/xhtml+xml")
	assertEquals("link[0].Properties", t, link[0].Properties, "nav")
	assertEquals("link[0].Refines", t, link[0].Refines, "#creator0")
}

func testSubject(t *testing.T, subject []DefaultAttributes) {
	assertSize("subject", t, len(subject), 2)
	assertEquals("subject[0].Id", t, subject[0].Id, "action-test")
	assertEquals("subject[0].Dir", t, subject[0].Dir, "ltr")
	assertEquals("subject[0].Lang", t, subject[0].Lang, "us")
	assertEquals("subject[0].Text", t, subject[0].Text, "Action")

	assertEquals("subject[1].Id", t, subject[1].Id, "fantasy-test")
	assertEquals("subject[1].Dir", t, subject[1].Dir, "ltr")
	assertEquals("subject[1].Lang", t, subject[1].Lang, "us")
	assertEquals("subject[1].Text", t, subject[1].Text, "Fantasy")
}

func testRights(t *testing.T, rights []DefaultAttributes) {
	assertSize("rights", t, len(rights), 1)
	assertEquals("rights[0].Id", t, rights[0].Id, "rights-test")
	assertEquals("rights[0].Dir", t, rights[0].Dir, "ltr")
	assertEquals("rights[0].Lang", t, rights[0].Lang, "us")
	assertEquals("rights[0].Text", t, rights[0].Text, "Test Rights")
}

func testRelation(t *testing.T, relation []DefaultAttributes) {
	assertSize("relation", t, len(relation), 1)
	assertEquals("relation[0].Id", t, relation[0].Id, "relation-test")
	assertEquals("relation[0].Dir", t, relation[0].Dir, "ltr")
	assertEquals("relation[0].Lang", t, relation[0].Lang, "us")
	assertEquals("relation[0].Text", t, relation[0].Text, "Test relation")
}

func testPublisher(t *testing.T, publisher []DefaultAttributes) {
	assertSize("publisher", t, len(publisher), 1)
	assertEquals("publisher[0].Id", t, publisher[0].Id, "publisher-test")
	assertEquals("publisher[0].Dir", t, publisher[0].Dir, "ltr")
	assertEquals("publisher[0].Lang", t, publisher[0].Lang, "us")
	assertEquals("publisher[0].Text", t, publisher[0].Text, "Test Publisher")
}

func testDescription(t *testing.T, description []DefaultAttributes) {
	assertSize("description", t, len(description), 1)
	assertEquals("description[0].Id", t, description[0].Id, "description-test")
	assertEquals("description[0].Dir", t, description[0].Dir, "ltr")
	assertEquals("description[0].Lang", t, description[0].Lang, "us")
	assertEquals("description[0].Text", t, description[0].Text, "Test description")
}

func testCreator(t *testing.T, creator []DefaultAttributes) {
	assertSize("creator", t, len(creator), 1)
	assertEquals("creator[0].Id", t, creator[0].Id, "creator0")
	assertEquals("creator[0].Dir", t, creator[0].Dir, "ltr")
	assertEquals("creator[0].Lang", t, creator[0].Lang, "us")
	assertEquals("creator[0].Text", t, creator[0].Text, "Last Name, First Name")
}

func testCoverage(t *testing.T, coverage []DefaultAttributes) {
	assertSize("coverage", t, len(coverage), 1)
	assertEquals("coverage[0].Id", t, coverage[0].Id, "coverage-test")
	assertEquals("coverage[0].Dir", t, coverage[0].Dir, "ltr")
	assertEquals("coverage[0].Lang", t, coverage[0].Lang, "us")
	assertEquals("coverage[0].Text", t, coverage[0].Text, "Test coverage")
}

func testContributor(t *testing.T, contributor []DefaultAttributes) {
	assertSize("contributor", t, len(contributor), 1)
	assertEquals("contributor[0].Id", t, contributor[0].Id, "contributor-test")
	assertEquals("contributor[0].Dir", t, contributor[0].Dir, "ltr")
	assertEquals("contributor[0].Lang", t, contributor[0].Lang, "us")
	assertEquals("contributor[0].Text", t, contributor[0].Text, "Test contributor")
}

func testType(t *testing.T, epubType []ID) {
	assertSize("type", t, len(epubType), 1)
	assertEquals("type[0].Id", t, epubType[0].Id, "type-test")
	assertEquals("type[0].Text", t, epubType[0].Text, "test type")
}

func testDate(t *testing.T, date []ID) {
	assertSize("date", t, len(date), 1)
	assertEquals("date[0].Id", t, date[0].Id, "date-test")
	assertEquals("date[0].Text", t, date[0].Text, "2020-03-04")
}

func testTitle(t *testing.T, title []DefaultAttributes) {
	assertSize("title", t, len(title), 3)
	assertEquals("title[0].Id", t, title[0].Id, "t1")
	assertEquals("title[0].Dir", t, title[0].Dir, "ltr")
	assertEquals("title[0].Lang", t, title[0].Lang, "us")
	assertEquals("title[0].Text", t, title[0].Text, "Test Book 01")

	assertEquals("title[1].Id", t, title[1].Id, "t2")
	assertEquals("title[1].Dir", t, title[1].Dir, "ltr")
	assertEquals("title[1].Lang", t, title[1].Lang, "us")
	assertEquals("title[1].Text", t, title[1].Text, "-")

	assertEquals("title[2].Id", t, title[2].Id, "t3")
	assertEquals("title[2].Dir", t, title[2].Dir, "ltr")
	assertEquals("title[2].Lang", t, title[2].Lang, "us")
	assertEquals("title[2].Text", t, title[2].Text, "Test Books")
}

func testLanguage(t *testing.T, language []ID) {
	assertSize("language", t, len(language), 1)
	assertEquals("language[0].ID", t, language[0].Id, "lang-test")
	assertEquals("language[0].Text", t, language[0].Text, "us")
}

func testIdentifier(t *testing.T, identifier []ID) {
	assertSize("identifier", t, len(identifier), 1)
	assertEquals("identifier[0].Id", t, identifier[0].Id, "identifier-test")
	assertEquals("identifier[0].Text", t, identifier[0].Text, "test identifier")
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
