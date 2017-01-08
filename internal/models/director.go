package models

import "encoding/xml"

type Director struct {
	XMLName xml.Name `xml:"Director"`
	Id      int      `xml:"id,attr"`
	Tag     string   `xml:"tag,attr"`
}
