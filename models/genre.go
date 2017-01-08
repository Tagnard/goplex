package models

import "encoding/xml"

type Genre struct {
	XMLName xml.Name `xml:"Genre"`
	Id      int      `xml:"id,attr"`
	Count   int      `xml:"count,attr"`
	Tag     string   `xml:"tag,attr"`
}