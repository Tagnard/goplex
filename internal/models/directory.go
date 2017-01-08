package models

import "encoding/xml"

type Directory struct {
	XMLName    xml.Name `xml:"Directory"`
	AllowSync  bool     `xml:"allowSync,attr"`
	Art        string   `xml:"art,attr"`
	Composite  string   `xml:"composite,attr"`
	Tag        string   `xml:"tag,attr"`
	Filters    int      `xml:"filters,attr"`
	Refreshing string   `xml:"refreshing,attr"`
	Thumb      string   `xml:"thumb,attr"`
	Key        int      `xml:"key,attr"`
	Type       string   `xml:"type,attr"`
	Title      string   `xml:"title,attr"`
	Agent      string   `xml:"agent,attr"`
	Scanner    string   `xml:"scanner,attr"`
	Language   string   `xml:"language,attr"`
	Uuid       string   `xml:"uuid,attr"`
	UpdatedAt  int      `xml:"updatedAt,attr"`
	CreatedAt  string   `xml:"createdAt,attr"`
	Location   Location `xml:"Location"`
}