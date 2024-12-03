package epub_v2

import (
	"encoding/xml"
	"github.com/mathieu-keller/epub-parser/model"
)

func getTitles(metaData []DefaultAttributes) *[]model.Title {
	titles := make([]model.Title, len(metaData))
	for i, title := range metaData {
		titles[i] = model.Title{
			Title:    title.Text,
			Language: title.Lang,
			Type:     "main",
			FileAs:   title.Text,
		}
	}
	return &titles
}

func getLanguages(metaData []ID) *[]string {
	languages := make([]string, len(metaData))
	for i, language := range metaData {
		languages[i] = language.Text
	}
	return &languages
}

func getCreators(metaData []Creator) *[]model.Creator {
	if metaData != nil {
		creators := make([]model.Creator, len(metaData))
		for i, creator := range metaData {
			role, ok := model.Relator[creator.Role]
			if !ok && creator.Role != "" {
				role = "unknown"
			}
			creators[i] = model.Creator{
				Name:     creator.Text,
				FileAs:   creator.FileAs,
				RawRole:  creator.Role,
				Language: creator.Lang,
				Role:     role,
			}
		}
		return &creators
	}
	return nil
}

func getDefaultAttributes(metaData []DefaultAttributes) *[]model.DefaultAttributes {
	if metaData != nil {
		defaultAttributes := make([]model.DefaultAttributes, len(metaData))
		for i, defaultAttribute := range metaData {
			defaultAttributes[i] = model.DefaultAttributes{
				Text:     defaultAttribute.Text,
				Language: defaultAttribute.Lang,
			}
		}
		return &defaultAttributes
	}
	return nil
}

func getDate(metaData []Date) *[]string {
	if metaData != nil {
		dates := make([]string, len(metaData))
		for i, date := range metaData {
			dates[i] = date.Text
		}
		return &dates
	}
	return nil
}

func ParseOpf(book *model.Book) error {
	opf := Package{}
	err := book.ReadXML(book.Container.Rootfile.Path, &opf)
	if err != nil {
		return err
	}

	identifiers := make([]model.Identifier, len(*opf.Metadata.Identifier))
	for i, identifier := range *opf.Metadata.Identifier {
		identifiers[i] = model.Identifier{
			Id:     identifier.Text,
			Scheme: identifier.Scheme,
		}
		if identifier.Id == opf.UniqueIdentifier {
			book.Metadata.MainId = model.Identifier{
				Id:     identifier.Text,
				Scheme: identifier.Scheme,
			}
		}
	}
	book.Metadata.Identifiers = &identifiers

	book.Metadata.Titles = getTitles(*opf.Metadata.Title)
	book.Metadata.Languages = getLanguages(*opf.Metadata.Language)
	book.Metadata.Creators = getCreators(*opf.Metadata.Creator)
	book.Metadata.Contributors = getCreators(*opf.Metadata.Contributor)
	book.Metadata.Publishers = getDefaultAttributes(*opf.Metadata.Publisher)
	book.Metadata.Subjects = getDefaultAttributes(*opf.Metadata.Subject)
	book.Metadata.Descriptions = getDefaultAttributes(*opf.Metadata.Description)
	book.Metadata.Dates = getDate(*opf.Metadata.Date)

	return err
}

type Package struct {
	XMLName          xml.Name  `xml:"package"`
	Metadata         *Metadata `xml:"metadata"`
	Manifest         *Manifest `xml:"manifest"`
	Version          string    `xml:"version,attr"`
	UniqueIdentifier string    `xml:"unique-identifier,attr"`
	ID               string    `xml:"id,attr,omitempty"`
	Prefix           string    `xml:"prefix,attr,omitempty"`
	Lang             string    `xml:"lang,attr,omitempty"`
	Dir              string    `xml:"dir,attr,omitempty"`
}

type Metadata struct {
	Title       *[]DefaultAttributes `xml:"title"`
	Identifier  *[]Identifier        `xml:"identifier"`
	Language    *[]ID                `xml:"language"`
	Creator     *[]Creator           `xml:"creator,omitempty"`
	Contributor *[]Creator           `xml:"contributor,omitempty"`
	Publisher   *[]DefaultAttributes `xml:"publisher,omitempty"`
	Subject     *[]DefaultAttributes `xml:"subject,omitempty"`
	Description *[]DefaultAttributes `xml:"description,omitempty"`
	Date        *[]Date              `xml:"date,omitempty"`
	Type        *[]ID                `xml:"type,omitempty"`
	Format      *[]ID                `xml:"format,omitempty"`
	Source      *[]DefaultAttributes `xml:"source,omitempty"`
	Relation    *[]DefaultAttributes `xml:"relation,omitempty"`
	Coverage    *[]DefaultAttributes `xml:"coverage,omitempty"`
	Rights      *[]DefaultAttributes `xml:"rights,omitempty"`
	Meta        *[]Meta              `xml:"meta,omitempty"`
}

type Creator struct {
	Text   string `xml:",chardata"`
	FileAs string `xml:"file-as,attr,omitempty"`
	Id     string `xml:"id,attr,omitempty"`
	Lang   string `xml:"lang,attr,omitempty"`
	Role   string `xml:"role,attr,omitempty"`
}

type DefaultAttributes struct {
	Text string `xml:",chardata"`
	Id   string `xml:"id,attr,omitempty"`
	Lang string `xml:"lang,attr,omitempty"`
}

type Date struct {
	Text  string `xml:",chardata"`
	Event string `xml:"event,attr,omitempty"`
	Id    string `xml:"id,attr,omitempty"`
}

type ID struct {
	Text string `xml:",chardata"`
	Id   string `xml:"id,attr,omitempty"`
}

type Identifier struct {
	Text   string `xml:",chardata"`
	Id     string `xml:"id,attr,omitempty"`
	Scheme string `xml:"scheme,attr,omitempty"`
}

type Meta struct {
	Text   string `xml:",chardata"`
	Lang   string `xml:"lang,attr,omitempty"`
	Name   string `xml:"name,attr"`
	Scheme string `xml:"scheme,attr,omitempty"`
}

type Manifest struct {
	Id   string  `xml:"id,attr,omitempty"`
	Item *[]Item `xml:"item"`
}

type Item struct {
	Id                string `xml:"id,attr"`
	Href              string `xml:"href,attr"`
	MediaType         string `xml:"media-type,attr"`
	Fallback          string `xml:"fallback,attr,omitempty"`
	FallbackStyle     string `xml:"fallback-style,attr,omitempty"`
	RequiredModules   string `xml:"required-modules,attr,omitempty"`
	RequiredNamespace string `xml:"required-namespace,attr,omitempty"`
}
