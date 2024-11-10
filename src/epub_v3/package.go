package epub_v3

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

func getMetaMap(metaData []Meta) *map[string]map[string]Meta {
	metaMap := make(map[string]map[string]Meta)
	for _, meta := range metaData {
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
	return &metaMap;
}

func getTitles(metaData []DefaultAttributes, metaMap map[string]map[string]Meta) *[]model.Title {
	titles := make([]model.Title, len(metaData))
	for i, title := range metaData {
		fileAs := getMetadata(metaMap, title.Id, "file-as")
		titleType := getMetadata(metaMap, title.Id, "title-type")
		titles[i] = model.Title{
			Title: title.Text,
			Language: title.Lang,
			Type: titleType,
			FileAs: fileAs,
		}
	}
	return &titles;
}

func getLanguages(metaData []ID) *[]string {
	languages := make([]string, len(metaData))
	for i, language := range metaData {
		languages[i] = language.Text
	}
	return &languages;
}

func getCreators(metaData []DefaultAttributes, metaMap map[string]map[string]Meta) *[]model.Creator {
	if metaData != nil {
		creators := make([]model.Creator, len(metaData))
		for i, creator := range metaData {
			fileAs := getMetadata(metaMap, creator.Id, "file-as")
			rawRole := getMetadata(metaMap, creator.Id, "role")
			role, ok := model.Relator[rawRole]
			if !ok && rawRole != "" {
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
		return &creators;
	}
	return nil;
}

func getDefaultAttributes(metaData []DefaultAttributes) *[]model.DefaultAttributes {
	if metaData != nil {
		defaultAttributes := make([]model.DefaultAttributes, len(metaData))
		for i, defaultAttribute := range metaData {
			defaultAttributes[i] = model.DefaultAttributes{
				Text:     defaultAttribute.Text,
				Language:   defaultAttribute.Lang,
			}
		}
		return &defaultAttributes
	}
	return nil;
}

func getDates(metaData []ID) *[]model.Date {
	if metaData != nil {
		dates := make([]model.Date, len(metaData))
		for i, date := range metaData {
			dates[i] = model.Date{
				Date:     date.Text,
				Type:   "publication",
			}
		}
		return &dates;
	}
	return nil;
}


func ParseOpf(book *model.Book) (error) {
	opf := Package{}
	err := book.ReadXML(book.Container.Rootfile.Path, &opf)
	if err != nil {
		return err
	}
	metaMap:= getMetaMap(*opf.Metadata.Meta)

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


	book.Metadata.Titles = getTitles(*opf.Metadata.Title, *metaMap)
	book.Metadata.Languages = getLanguages(*opf.Metadata.Language)
	book.Metadata.Creators = getCreators(*opf.Metadata.Creator, *metaMap)
	book.Metadata.Contributors = getCreators(*opf.Metadata.Contributor, *metaMap)
	book.Metadata.Publishers = getDefaultAttributes(*opf.Metadata.Publisher)
	book.Metadata.Subjects = getDefaultAttributes(*opf.Metadata.Subject)
	book.Metadata.Descriptions = getDefaultAttributes(*opf.Metadata.Description)
	book.Metadata.Dates = getDates(*opf.Metadata.Date)
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
