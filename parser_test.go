package epub

import (
	"io/ioutil"
	"os"
	"strconv"
	"testing"
)

func Test_parse_epub_3_0_opf(t *testing.T) {
	b, err := os.Open("./test_opf.xml")
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	defer b.Close()
	reader := ioutil.NopCloser(b)
	defer reader.Close()
	opf, err := ReadOpf(reader)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	assertEquals(t, opf.Lang, "us")
	assertEquals(t, opf.UniqueIdentifier, "p1234")
	assertEquals(t, opf.Version, "3.0")
	assertEquals(t, opf.ID, "test-id")
	assertEquals(t, opf.Prefix, "pre")

	testMeta(t, *opf.Metadata.Meta)
	testIdentifier(t, *opf.Metadata.Identifier)
	testLanguage(t, *opf.Metadata.Language)
	testTitle(t, *opf.Metadata.Title)
	testDate(t, *opf.Metadata.Date)
	testType(t, *opf.Metadata.Type)
	testContributor(t, *opf.Metadata.Contributor)
	testCoverage(t, *opf.Metadata.Coverage)
	testCreator(t, *opf.Metadata.Creator)
	testDescription(t, *opf.Metadata.Description)
	testPublisher(t, *opf.Metadata.Publisher)
	testRelation(t, *opf.Metadata.Relation)
	testRights(t, *opf.Metadata.Rights)
	testSubject(t, *opf.Metadata.Subject)
	testLink(t, *opf.Metadata.Link)

}

func testMeta(t *testing.T, meta []Meta) {
	assertSize(t, len(meta), 2)
	assertEquals(t, meta[0].Dir, "ltr")
	assertEquals(t, meta[0].Id, "meta-test")
	assertEquals(t, meta[0].Scheme, "scheme")
	assertEquals(t, meta[0].Lang, "us")
	assertEquals(t, meta[0].Refines, "#t1")
	assertEquals(t, meta[0].Property, "title-type")
	assertEquals(t, meta[0].Text, "main")
	assertEquals(t, meta[0].Content, "")
	assertEquals(t, meta[0].Name, "")

	assertEquals(t, meta[1].Dir, "")
	assertEquals(t, meta[1].Id, "")
	assertEquals(t, meta[1].Scheme, "")
	assertEquals(t, meta[1].Lang, "")
	assertEquals(t, meta[1].Refines, "")
	assertEquals(t, meta[1].Property, "")
	assertEquals(t, meta[1].Text, "")
	assertEquals(t, meta[1].Content, "vertical-rl")
	assertEquals(t, meta[1].Name, "primary-writing-mode")
}

func testLink(t *testing.T, link []Link) {
	assertSize(t, len(link), 1)
	assertEquals(t, link[0].Id, "link-test")
	assertEquals(t, link[0].Href, "front.xhtml#meta-json")
	assertEquals(t, link[0].Rel, "record")
	assertEquals(t, link[0].MediaType, "application/xhtml+xml")
	assertEquals(t, link[0].Properties, "nav")
	assertEquals(t, link[0].Refines, "#creator0")
}

func testSubject(t *testing.T, subject []DefaultAttributes) {
	assertSize(t, len(subject), 2)
	assertEquals(t, subject[0].Id, "action-test")
	assertEquals(t, subject[0].Dir, "ltr")
	assertEquals(t, subject[0].Lang, "us")
	assertEquals(t, subject[0].Text, "Action")

	assertEquals(t, subject[1].Id, "fantasy-test")
	assertEquals(t, subject[1].Dir, "ltr")
	assertEquals(t, subject[1].Lang, "us")
	assertEquals(t, subject[1].Text, "Fantasy")
}

func testRights(t *testing.T, rights []DefaultAttributes) {
	assertSize(t, len(rights), 1)
	assertEquals(t, rights[0].Id, "rights-test")
	assertEquals(t, rights[0].Dir, "ltr")
	assertEquals(t, rights[0].Lang, "us")
	assertEquals(t, rights[0].Text, "Test Rights")
}

func testRelation(t *testing.T, relation []DefaultAttributes) {
	assertSize(t, len(relation), 1)
	assertEquals(t, relation[0].Id, "relation-test")
	assertEquals(t, relation[0].Dir, "ltr")
	assertEquals(t, relation[0].Lang, "us")
	assertEquals(t, relation[0].Text, "Test relation")
}

func testPublisher(t *testing.T, publisher []DefaultAttributes) {
	assertSize(t, len(publisher), 1)
	assertEquals(t, publisher[0].Id, "publisher-test")
	assertEquals(t, publisher[0].Dir, "ltr")
	assertEquals(t, publisher[0].Lang, "us")
	assertEquals(t, publisher[0].Text, "Test Publisher")
}

func testDescription(t *testing.T, description []DefaultAttributes) {
	assertSize(t, len(description), 1)
	assertEquals(t, description[0].Id, "description-test")
	assertEquals(t, description[0].Dir, "ltr")
	assertEquals(t, description[0].Lang, "us")
	assertEquals(t, description[0].Text, "Test description")
}

func testCreator(t *testing.T, creator []DefaultAttributes) {
	assertSize(t, len(creator), 1)
	assertEquals(t, creator[0].Id, "creator0")
	assertEquals(t, creator[0].Dir, "ltr")
	assertEquals(t, creator[0].Lang, "us")
	assertEquals(t, creator[0].Text, "Last Name, First Name")
}

func testCoverage(t *testing.T, coverage []DefaultAttributes) {
	assertSize(t, len(coverage), 1)
	assertEquals(t, coverage[0].Id, "coverage-test")
	assertEquals(t, coverage[0].Dir, "ltr")
	assertEquals(t, coverage[0].Lang, "us")
	assertEquals(t, coverage[0].Text, "Test coverage")
}

func testContributor(t *testing.T, contributor []DefaultAttributes) {
	assertSize(t, len(contributor), 1)
	assertEquals(t, contributor[0].Id, "contributor-test")
	assertEquals(t, contributor[0].Dir, "ltr")
	assertEquals(t, contributor[0].Lang, "us")
	assertEquals(t, contributor[0].Text, "Test contributor")
}

func testType(t *testing.T, epubType []ID) {
	assertSize(t, len(epubType), 1)
	assertEquals(t, epubType[0].Id, "type-test")
	assertEquals(t, epubType[0].Text, "test type")
}

func testDate(t *testing.T, date []ID) {
	assertSize(t, len(date), 1)
	assertEquals(t, date[0].Id, "date-test")
	assertEquals(t, date[0].Text, "2020-03-04")
}

func testTitle(t *testing.T, title []DefaultAttributes) {
	assertSize(t, len(title), 3)
	assertEquals(t, title[0].Id, "t1")
	assertEquals(t, title[0].Dir, "ltr")
	assertEquals(t, title[0].Lang, "us")
	assertEquals(t, title[0].Text, "Test Book 01")

	assertEquals(t, title[1].Id, "t2")
	assertEquals(t, title[1].Dir, "ltr")
	assertEquals(t, title[1].Lang, "us")
	assertEquals(t, title[1].Text, "-")

	assertEquals(t, title[2].Id, "t3")
	assertEquals(t, title[2].Dir, "ltr")
	assertEquals(t, title[2].Lang, "us")
	assertEquals(t, title[2].Text, "Test Books")
}

func testLanguage(t *testing.T, language []ID) {
	assertSize(t, len(language), 1)
	assertEquals(t, language[0].Id, "lang-test")
	assertEquals(t, language[0].Text, "us")
}

func testIdentifier(t *testing.T, identifier []ID) {
	assertSize(t, len(identifier), 1)
	assertEquals(t, identifier[0].Id, "identifier-test")
	assertEquals(t, identifier[0].Text, "test identifier")
}

func assertSize(t *testing.T, actuallySize int, expectedSize int) {
	if actuallySize != expectedSize {
		t.Log(strconv.Itoa(actuallySize) + " != " + strconv.Itoa(expectedSize))
		t.Fail()
	}
}

func assertEquals(t *testing.T, a string, b string) {
	if a != b {
		t.Log(a + " != " + b)
		t.Fail()
	}
}
