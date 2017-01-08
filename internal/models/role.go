package models

import "encoding/xml"

type Role struct {
	XMLName xml.Name `xml:"Role"`
	Id      int      `xml:"id,attr"`
	Role    string   `xml:"role,attr"`
	Tag     string   `xml:"tag,attr"`
	Thumb   string   `xml:"thumb,attr"`
}
