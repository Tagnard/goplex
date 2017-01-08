package models

import "encoding/xml"

type Video struct {
	XMLName               xml.Name         `xml:"Video"`
	Title                 string           `xml:"title,attr"`
	Art                   string           `xml:"art,attr"`
	ChapterSource         string           `xml:"chapterSource,attr"`
	ContentRating         string           `xml:"contentRating,attr"`
	Duration              int              `xml:"duration,attr"`
	ParentKey             string           `xml:"parentKey,attr"`
	ParentRatingKey       int              `xml:"parentRatingKey,attr"`
	ParentThumb           string           `xml:"parentThumb,attr"`
	GrandparentArt        string           `xml:"grandparentArt,attr"`
	GrandparentKey        string           `xml:"grandparentKey,attr"`
	GrandparentRatingKey  int              `xml:"grandparentRatingKey,attr"`
	GrandparentThumb      string           `xml:"grandparentThumb,attr"`
	GrandparentTitle      string           `xml:"grandparentTitle,attr"`
	Guid                  string           `xml:"guid,attr"`
	Index                 int              `xml:"index,attr"`
	Key                   string           `xml:"key,attr"`
	LibrarySectionId      int              `xml:"librarySectionID,attr"`
	ParentIndex           int              `xml:"parentIndex,attr"`
	Rating                string           `xml:"rating,attr"`
	RatingKey             int              `xml:"ratingKey,attr"`
	SessionKey            int              `xml:"sessionKey,attr"`
	Summary               string           `xml:"summary,attr"`
	Thumb                 string           `xml:"thumb,attr"`
	ViewOffset            int              `xml:"viewOffset,attr"`
	Year                  int              `xml:"year,attr"`
	AddedAt               int              `xml:"addedAt,attr"`
	UpdatedAt             int              `xml:"updatedAt,attr"`
	OriginallyAvailableAt string           `xml:"originallyAvailableAt,attr"`
	LastViewedAt          int              `xml:"lastViewedAt,attr"`
	Type                  string           `xml:"type,attr"`

	Media                 Media            `xml:"Media"`
	User                  User             `xml:"User"`
	Player                Player           `xml:"Player"`
	Genre                 []Genre 	       `xml:"Genre"`
	Director              []Director       `xml:"Director"`
	Writer                []Writer         `xml:"Writer"`
	Country               []Country        `xml:"Country"`
	Role                  []Role           `xml:"Role"`
	TranscodeSession      TranscodeSession `xml:"TranscodeSession"`
}