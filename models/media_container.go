package models

import "encoding/xml"

type MediaContainer struct {
	XMLName             xml.Name    `xml:"MediaContainer"`
	Size                int         `xml:"size,attr"`
	AllowSync           bool        `xml:"allowSync,attr"`
	Art                 string      `xml:"art,attr"`
	identifier          string      `xml:"identifier,attr"`
	librarySectionID    int         `xml:"librarySectionID,attr"`
	librarySectionTitle string      `xml:"librarySectionTitle,attr"`
	librarySectionUUID  string      `xml:"librarySectionUUID,attr"`
	mediaTagPrefix      string      `xml:"mediaTagPrefix,attr"`
	mediaTagVersion     string      `xml:"mediaTagVersion,attr"`
	Thumb               string      `xml:"thumb,attr"`
	Title1              string      `xml:"title1,attr"`
	Title2              string      `xml:"title2,attr"`
	ViewGroup           string      `xml:"viewGroup,attr"`
	ViewMode            int         `xml:"viewMode,attr"`

	Video               []Video     `xml:"Video"`
	Directory           []Directory `xml:"Directory"`
}


/*

<MediaContainer
	size="697"
	allowSync="1"
	art="/:/resources/movie-fanart.jpg"
	identifier="com.plexapp.plugins.library"
	librarySectionID="2"
	librarySectionTitle="Movies"
	librarySectionUUID="b83a5078-d0e1-406d-8a7b-824d22f9e95d"
	mediaTagPrefix="/system/bundle/media/flags/"
	mediaTagVersion="1473718680"
	thumb="/:/resources/movie.png"
	title1="Movies"
	title2="All Movies"
	viewGroup="movie"
	viewMode="65592">

*/