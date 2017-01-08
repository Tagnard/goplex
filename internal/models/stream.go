package models

import "encoding/xml"

type Stream struct {
	XMLName            xml.Name `xml:"Stream"`
	BitDepth           int      `xml:"bitDepth,attr"`
	Bitrate            int      `xml:"bitrate,attr"`
	Cabac              int      `xml:"cabac,attr"`
	ChromaSubsampling  string   `xml:"chromaSubsampling,attr"`
	Codec              string   `xml:"codec,attr"`
	CodecID            string   `xml:"codedID,attr"`
	Default            bool     `xml:"default,attr"`
	Duration           int      `xml:"duration,attr"`
	FrameRate          string   `xml:"frameRate,attr"`
	FrameRateMode      string   `xml:"frameRateMode,attr"`
	HasScalingMatrix   bool     `xml:"hasScalingMatrix,attr"`
	Width              int      `xml:"width,attr"`
	Height             int      `xml:"height,attr"`
	Id                 int      `xml:"id,attr"`
	Index              int      `xml:"index,attr"`
	Level              int      `xml:"level,attr"`
	PixelFormat        string   `xml:"pixelFormat,attr"`
	Profile            string   `xml:"profile,attr"`
	RefFrames          int      `xml:"refFrames,attr"`
	RequiredBandwidths string   `xml:"requiredBandwidths,attr"`
	ScanType           string   `xml:"scanType,attr"`
	StreamIdentifier   int      `xml:"streamIdentifier,attr"`
	StreamType         int      `xml:"streamType,attr"`
}