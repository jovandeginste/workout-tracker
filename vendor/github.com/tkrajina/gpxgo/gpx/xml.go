// Copyright 2013, 2014 Peter Vasil, Tomo Krajina. All
// rights reserved. Use of this source code is governed
// by a BSD-style license that can be found in the
// LICENSE file.

package gpx

import (
	"bytes"
	"encoding/xml"
	"errors"
	"io"
	"os"
	"strings"
	"time"

	"golang.org/x/net/html/charset"
)

const formattingTimelayout = "2006-01-02T15:04:05Z"

// parsingTimelayouts defines a list of possible time formats
var parsingTimelayouts = []string{
	"2006-01-02T15:04:05.000Z",
	formattingTimelayout,
	"2006-01-02T15:04:05+00:00",
	"2006-01-02T15:04:05",
	"2006-01-02 15:04:05Z",
	"2006-01-02 15:04:05",
}

func init() {
	/*
		fmt.Println("----------------------------------------------------------------------------------------------------")
		fmt.Println("This API is experimental, it *will* change")
		fmt.Println("----------------------------------------------------------------------------------------------------")
	*/
}

// ToXmlParams contains settings for xml transformation
type ToXmlParams struct {
	Version string
	Indent  bool
}

// ToXml returns the xml representation of the GPX object.
// Params are optional, you can set null to use GPXs Version and no indentation.
func ToXml(g *GPX, params ToXmlParams) ([]byte, error) {
	version := g.Version
	if len(params.Version) > 0 {
		version = params.Version
	}
	indentation := params.Indent

	var replacemends map[string]string
	var gpxDoc interface{}
	if version == "1.0" {
		gpxDoc = convertToGpx10Models(g)
	} else if version == "1.1" {
		gpxDoc, replacemends = convertToGpx11Models(g)
	} else {
		g.Version = "1.1"
		gpxDoc, replacemends = convertToGpx11Models(g)
	}

	var buffer bytes.Buffer
	buffer.WriteString(xml.Header)
	if indentation {
		b, err := xml.MarshalIndent(gpxDoc, "", "	")
		if err != nil {
			return nil, err
		}
		buffer.Write(b)
	} else {
		b, err := xml.Marshal(gpxDoc)
		if err != nil {
			return nil, err
		}
		buffer.Write(b)
	}

	byts := buffer.Bytes()

	for replKey, replVal := range replacemends {
		byts = bytes.Replace(byts, []byte(replKey), []byte(replVal), -1)
	}

	return byts, nil
}

func guessGPXVersion(bytes []byte) (string, error) {
	bytesCount := 1000
	if len(bytes) < 1000 {
		bytesCount = len(bytes)
	}

	startOfDocument := string(bytes[:bytesCount])

	parts := strings.Split(startOfDocument, "<gpx")
	if len(parts) <= 1 {
		return "", errors.New("invalid GPX file, cannot find start of <gpx>")
	}
	parts = strings.Split(parts[1], "version=")

	if len(parts) <= 1 {
		return "", errors.New("invalid GPX file, cannot find version in <gpx ...>")
	}

	version := strings.TrimLeft(parts[1], `'" `)
	if strings.HasPrefix(version, "1.0") {
		return "1.0", nil
	} else if strings.HasPrefix(version, "1.1") {
		return "1.1", nil
	}

	return "", errors.New("invalid GPX file, cannot find version")
}

func parseGPXTime(timestr string) (*time.Time, error) {
	if strings.Contains(timestr, ".") {
		// Probably seconds with milliseconds
		timestr = strings.Split(timestr, ".")[0]
	}
	timestr = strings.Trim(timestr, " \t\n\r")
	for _, timeLayout := range parsingTimelayouts {
		t, err := time.Parse(timeLayout, timestr)

		if err == nil {
			return &t, nil
		}
	}

	return nil, errors.New("Cannot parse " + timestr)
}

func formatGPXTime(time *time.Time) string {
	if time == nil {
		return ""
	}
	if time.Year() <= 1 {
		// Invalid date:
		return ""
	}
	return time.Format(formattingTimelayout)
}

// ParseFile parses a gpx file and returns a GPX object
func ParseFile(fileName string) (*GPX, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	return Parse(f)
}

// ParseBytes parses GPX from bytes
func ParseBytes(buf []byte) (*GPX, error) {
	return Parse(bytes.NewReader(buf))
}

func ParseDecoder(decoder *xml.Decoder, initialBytes []byte) (*GPX, error) {
	version, err := guessGPXVersion(initialBytes)
	if err != nil {
		// Unknown version, try with 1.1
		version = "1.1"
	}

	switch version {
	case "1.0":
		g := &gpx10Gpx{}
		err = decoder.Decode(&g)
		if err != nil {
			return nil, err
		}
		return convertFromGpx10Models(g), nil
	case "1.1":
		g := &gpx11Gpx{}
		err = decoder.Decode(&g)
		if err != nil {
			return nil, err
		}
		return convertFromGpx11Models(g), nil
	default:
		return nil, errors.New("Invalid version:" + version)
	}
}

// Parse parses GPX from io.Reader
func Parse(inReader io.Reader) (*GPX, error) {
	// at most 1000 bytes will make guessGPXVersion happy
	buf := make([]byte, 1000)

	n, err := inReader.Read(buf)
	if err != nil {
		return nil, err
	}
	buf = buf[:n]

	reader := io.MultiReader(bytes.NewReader(buf), inReader)
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReaderLabel

	return ParseDecoder(decoder, buf)
}

// ParseString parses GPX from string
func ParseString(str string) (*GPX, error) {
	return Parse(strings.NewReader(str))
}
