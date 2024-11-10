package model

type Package struct {
	Version          string `xml:"version,attr"`
	UniqueIdentifier string `xml:"unique-identifier,attr"`
	ID               string `xml:"id,attr,omitempty"`
	Prefix           string `xml:"prefix,attr,omitempty"`
	Lang             string `xml:"lang,attr,omitempty"`
	Dir              string `xml:"dir,attr,omitempty"`
}
