package models

import "encoding/xml"

type Writer struct {
	XMLName xml.Name `xml:"Writer"`
	Id      int      `xml:"id,attr"`
	Tag     string   `xml:"tag,attr"`
}
