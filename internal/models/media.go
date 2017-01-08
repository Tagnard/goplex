package models

import "encoding/xml"

type Media struct {
	XMLName xml.Name `xml:"Media"`
	Part    []Part   `xml:"Part"`
}