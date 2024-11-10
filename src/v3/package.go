package v3

import (
	"encoding/xml"
	"strings"
	"github.com/mathieu-keller/epub-parser/model"
)

func getMetadata(metaData map[string]map[string]Meta, id string, metaDataKey string) string {
	if id != "" {
		idData, idOk := metaData[id]
		if idOk {
			keyData, keyOk := idData[metaDataKey]
			if keyOk {
				return keyData.Text
			}
		}
	}
	return ""
}


func ParseOpf(book *model.Book) (error) {
	opf := Package{}
	err := book.ReadXML(book.Container.Rootfile.Path, &opf)
	if err != nil {
		return err
	}
	metaMap := make(map[string]map[string]Meta)
	for _, meta := range *opf.Metadata.Meta {
		if meta.Refines != "" && meta.Property != "" {
			id := strings.Replace(meta.Refines, "#", "", 1)
			innerMap, ok := metaMap[id]
			if !ok {
				innerMap = make(map[string]Meta)
				metaMap[id] = innerMap
			}
			innerMap[meta.Property] = meta
		}
	}

	identifiers := make([]model.Identifier, len(*opf.Metadata.Identifier))
	for i, identifier := range *opf.Metadata.Identifier {
		identifiers[i] = model.Identifier{
			Id: identifier.Text,
			Scheme: "",
		}
		if identifier.Id == opf.UniqueIdentifier {
			book.Metadata.MainId = model.Identifier{
				Id: identifier.Text,
				Scheme: "",
			}
		}
	}
	book.Metadata.Identifiers = &identifiers

	titles := make([]model.Title, len(*opf.Metadata.Title))
	for i, title := range *opf.Metadata.Title {
		fileAs := getMetadata(metaMap, title.Id, "file-as")
		titleType := getMetadata(metaMap, title.Id, "title-type")
		titles[i] = model.Title{
			Title: title.Text,
			Language: title.Lang,
			Type: titleType,
			FileAs: fileAs,
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
			fileAs := getMetadata(metaMap, creator.Id, "file-as")
			rawRole := getMetadata(metaMap, creator.Id, "role")
			role, ok := model.Relator[rawRole]
			if !ok {
				role = "unknown"
			}
			creators[i] = model.Creator{
				Name:     creator.Text,
				FileAs:   fileAs,
				RawRole:  rawRole,
				Language: creator.Lang,
				Role: role,
			}
		}
		book.Metadata.Creators = &creators
	}
	if opf.Metadata.Contributor != nil {
		contributors := make([]model.Creator, len(*opf.Metadata.Contributor))
		for i, contributor := range *opf.Metadata.Contributor {
			fileAs := getMetadata(metaMap, contributor.Id, "file-as")
			rawRole := getMetadata(metaMap, contributor.Id, "role")
			role, ok := model.Relator[rawRole]
			if !ok {
				role = "unknown"
			}
			contributors[i] = model.Creator{
				Name:     contributor.Text,
				FileAs:   fileAs,
				RawRole:  rawRole,
				Role: role,
				Language: contributor.Lang,
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
				Type:   "publication",
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

type ID struct {
	Id   string `xml:"id,attr,omitempty"`
	Text string `xml:",chardata"`
}

type DefaultAttributes struct {
	Id   string `xml:"id,attr,omitempty"`
	Dir  string `xml:"dir,attr,omitempty"`
	Lang string `xml:"lang,attr,omitempty"`
	Text string `xml:",chardata"`
}

type Manifest struct {
	Id   string  `xml:"id,attr,omitempty"`
	Item *[]Item `xml:"item"`
}

type Item struct {
	Id           string `xml:"id,attr"`
	Href         string `xml:"href,attr"`
	MediaType    string `xml:"media-type,attr"`
	Fallback     string `xml:"fallback,attr,omitempty"`
	MediaOverlay string `xml:"media-overlay,attr,omitempty"`
	Properties   string `xml:"properties,attr,omitempty"`
}

type Metadata struct {
	Identifier  *[]ID                `xml:"identifier"`
	Language    *[]ID                `xml:"language"`
	Title       *[]DefaultAttributes `xml:"title"`
	Meta        *[]Meta              `xml:"meta"`
	Date        *[]ID                `xml:"date,omitempty"`
	Type        *[]ID                `xml:"type,omitempty"`
	Contributor *[]DefaultAttributes `xml:"contributor,omitempty"`
	Coverage    *[]DefaultAttributes `xml:"coverage,omitempty"`
	Creator     *[]DefaultAttributes `xml:"creator,omitempty"`
	Description *[]DefaultAttributes `xml:"description,omitempty"`
	Publisher   *[]DefaultAttributes `xml:"publisher,omitempty"`
	Relation    *[]DefaultAttributes `xml:"relation,omitempty"`
	Rights      *[]DefaultAttributes `xml:"rights,omitempty"`
	Subject     *[]DefaultAttributes `xml:"subject,omitempty"`
	Link        *[]Link              `xml:"link,omitempty"`
}

type Link struct {
	Href       string `xml:"href,attr"`
	Rel        string `xml:"rel,attr"`
	Id         string `xml:"id,attr,omitempty"`
	MediaType  string `xml:"media-type,attr,omitempty"`
	Properties string `xml:"properties,attr,omitempty"`
	Refines    string `xml:"refines,attr,omitempty"`
}

type Meta struct {
	Id       string `xml:"id,attr,omitempty"`
	Dir      string `xml:"dir,attr,omitempty"`
	Lang     string `xml:"lang,attr,omitempty"`
	Property string `xml:"property,attr,omitempty"` //omitempty because the deprecated meta has no property
	Refines  string `xml:"refines,attr,omitempty"`
	Scheme   string `xml:"scheme,attr,omitempty"`
	Name     string `xml:"name,attr,omitempty"`    //deprecated
	Content  string `xml:"content,attr,omitempty"` //deprecated
	Text     string `xml:",chardata"`
}
