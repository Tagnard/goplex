package models

import "encoding/xml"

type Player struct {
	XMLName           xml.Name `xml:"Player"`
	Address           string `xml:"address,attr"`
	Device            string `xml:"device,attr"`
	MachineIdentifier string `xml:"machineIdentifier,attr"`
	Model             string `xml:"model,attr"`
	Platform          string `xml:"platform,attr"`
	PlatformVersion   string `xml:"platformVersion,attr"`
	Product           string `xml:"product,attr"`
	Profile           string `xml:"profile,attr"`
	State             string `xml:"state,attr"`
	Title             string `xml:"title,attr"`
	Vendor            string `xml:"vendor,attr"`
	Version           string `xml:"versio,attr"`
}
