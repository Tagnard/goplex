package models

import "encoding/xml"

type Part struct {
	XMLName xml.Name `xml:"Part"`
	Stream  []Stream `xml:"Stream"`
}