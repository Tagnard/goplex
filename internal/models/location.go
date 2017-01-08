package models

import "encoding/xml"

type Location struct {
	XMLName xml.Name `xml:"Location"`
	Id      int      `xml:"id,attr"`
	path    string   `xml:"path,attr"`
}