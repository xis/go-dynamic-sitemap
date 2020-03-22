package sitemap

import (
	"encoding/xml"
)

// Sitemap @ sitemap contains raw xml bytes
type Sitemap struct {
	Data []byte
}

// URLSet @ url set struct
type URLSet struct {
	XMLName xml.Name `xml:"urlset"`
	Xmlns   string   `xml:"xmlns,attr"`
	URL     []URL    `xml:"url"`
}

// URL @ url struct
type URL struct {
	Loc string `xml:"loc"`
}

// Create @ creates xml
func Create(prefix string, data []string) *URLSet {
	url := URL{}
	set := URLSet{
		Xmlns: "http://www.sitemaps.org/schemas/sitemap/0.9",
	}
	for i := range data {
		url.Loc = prefix + data[i]
		set.URL = append(set.URL, url)
	}

	return &set
}

// After @ resumes to create xml
func (u *URLSet) After(prefix string, data []string) *URLSet {
	url := URL{}
	for i := range data {
		url.Loc = prefix + data[i]
		u.URL = append(u.URL, url)
	}
	return u
}

// Result @ get sitemap struct adress and writes xml raw bytes on it
func (u *URLSet) Result(s *Sitemap) error {
	out, err := xml.MarshalIndent(u, " ", "  ")
	if err != nil {
		return err
	}
	header := []byte(xml.Header)
	out = append(header, out...)
	s.Data = out
	return nil
}
