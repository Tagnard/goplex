package models

import "encoding/xml"

type User struct {
	XMLName  xml.Name `xml:"User"`
	Id       int      `xml:"id,attr"`
	Thumb    string   `xml:"thumb,attr"`
	Title    string   `xml:"title,attr"`
}
