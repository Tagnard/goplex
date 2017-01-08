package models

import "encoding/xml"

type TranscodeSession struct {
	XMLName          xml.Name `xml:"TranscodeSession"`
	Key              string   `xml:"key,attr"`
	Throttled        bool     `xml:"throttled,attr"`
	Complete         bool     `xml:"complete,attr"`
	Progress         float32  `xml:"progress,attr"`
	Speed            int      `xml:"speed,attr"`
	Duration         int      `xml:"duration,attr"`
	Remaining        int      `xml:"remaining,attr"`
	Context          string   `xml:"context,attr"`
	VideoDecision    string   `xml:"videoDecision,attr"`
	AudioDecision    string   `xml:"audioDecision,attr"`
	SubtitleDecision string   `xml:"subtitleDecision,attr"`
	Protocol         string   `xml:"protocol,attr"`
	Container        string   `xml:"container,attr"`
	VideoCodec       string   `xml:"videoCodec,attr"`
	AudioCodec       string   `xml:"audioCodec,attr"`
	AudioChannels    int      `xml:"audioChannels,attr"`
	Width            int      `xml:"width,attr"`
	Height           int      `xml:"height,attr"`
}