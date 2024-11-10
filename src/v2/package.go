package v2

import (
	"encoding/xml"

	"github.com/mathieu-keller/epub-parser/model"
)

func ParseOpf(book *model.Book) (error) {
	opf := Package{}
	err := book.ReadXML(book.Container.Rootfile.Path, &opf)
	if err != nil {
		return err
	}

	identifiers := make([]model.Identifier, len(*opf.Metadata.Identifier))
	for i, identifier := range *opf.Metadata.Identifier {
		identifiers[i] = model.Identifier{
			Id: identifier.Text,
			Scheme: identifier.Scheme,
		}
		if identifier.Id == opf.UniqueIdentifier {
			book.Metadata.MainId = model.Identifier{
				Id: identifier.Text,
				Scheme: identifier.Scheme,
			}
		}
	}
	book.Metadata.Identifiers = &identifiers

	titles := make([]model.Title, len(*opf.Metadata.Title))
	for i, title := range *opf.Metadata.Title {
		titles[i] = model.Title{
			Title: title.Text,
			Language: title.Lang,
			Type: "main",
			FileAs: title.Text,
		}
	}
	book.Metadata.Titles = &titles

	languages := make([]string, len(*opf.Metadata.Language))
	for i, language := range *opf.Metadata.Language {
		languages[i] = language.Text
	}
	book.Metadata.Languages = &languages

	if opf.Metadata.Creator != nil {
		creators := make([]model.Creator, len(*opf.Metadata.Creator))
		for i, creator := range *opf.Metadata.Creator {
			role, ok := model.Relator[creator.Role]
			if !ok {
				role = "unknown"
			}
			creators[i] = model.Creator{
				Name:     creator.Text,
				FileAs:   creator.FileAs,
				RawRole:  creator.Role,
				Language: creator.Lang,
				Role: role,
			}
		}
		book.Metadata.Creators = &creators
	}
	if opf.Metadata.Contributor != nil {
		contributors := make([]model.Creator, len(*opf.Metadata.Contributor))
		for i, contributor := range *opf.Metadata.Contributor {
			role, ok := model.Relator[contributor.Role]
			if !ok {
				role = "unknown"
			}
			contributors[i] = model.Creator{
				Name:     contributor.Text,
				FileAs:   contributor.FileAs,
				RawRole:  contributor.Role,
				Language: contributor.Lang,
				Role: role,
			}
		}
		book.Metadata.Contributors = &contributors
	}

	if opf.Metadata.Publisher != nil {
		publishers := make([]model.DefaultAttributes, len(*opf.Metadata.Publisher))
		for i, publisher := range *opf.Metadata.Publisher {
			publishers[i] = model.DefaultAttributes{
				Text:     publisher.Text,
				Language:   publisher.Lang,
			}
		}
		book.Metadata.Publishers = &publishers
	}

	if opf.Metadata.Subject != nil {
		subjects := make([]model.DefaultAttributes, len(*opf.Metadata.Subject))
		for i, subject := range *opf.Metadata.Subject {
			subjects[i] = model.DefaultAttributes{
				Text:     subject.Text,
				Language:   subject.Lang,
			}
		}
		book.Metadata.Subjects = &subjects
	}

	if opf.Metadata.Description != nil {
		descriptions := make([]model.DefaultAttributes, len(*opf.Metadata.Description))
		for i, description := range *opf.Metadata.Description {
			descriptions[i] = model.DefaultAttributes{
				Text:     description.Text,
				Language:   description.Lang,
			}
		}
		book.Metadata.Descriptions = &descriptions
	}

	if opf.Metadata.Date != nil {
		dates := make([]model.Date, len(*opf.Metadata.Date))
		for i, date := range *opf.Metadata.Date {
			dates[i] = model.Date{
				Date:     date.Text,
				Type:   date.Event,
			}
		}
		book.Metadata.Dates = &dates
	}

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
