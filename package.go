package epub

import "encoding/xml"

type Package struct {
	XMLName          xml.Name  `xml:"package"`
	Metadata         *Metadata `xml:"metadata"`
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
	Text     string `xml:",chardata"`              //omitempty because the deprecated meta has no text
}
