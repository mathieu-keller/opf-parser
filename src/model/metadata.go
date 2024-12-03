package model

type Metadata struct {
	MainId       Identifier
	Titles       *[]Title
	Identifiers  *[]Identifier
	Languages    *[]string
	Creators     *[]Creator
	Contributors *[]Creator
	Publishers   *[]DefaultAttributes
	Subjects     *[]DefaultAttributes
	Descriptions *[]DefaultAttributes
	Dates        *[]string
}

type Creator struct {
	Name     string
	Language string
	FileAs   string
	Role     string
	RawRole  string
}

type Title struct {
	Title    string
	Language string
	Type     string
	FileAs   string
}

type DefaultAttributes struct {
	Text     string
	Language string
}

type Identifier struct {
	Id     string
	Scheme string
}
