package models

import "encoding/xml"

type Country struct {
	XMLName xml.Name `xml:"Country"`
	Id      int      `xml:"id,attr"`
	Count   int      `xml:"count,attr"`
	Tag     string   `xml:"tag,attr"`
}
