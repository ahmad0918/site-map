package main

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"time"
)

type URL struct {
	Loc        string    `xml:"loc"`
	LastMod    time.Time `xml:"lastmod"`
	ChangeFreq string    `xml:"changefreq"`
	Priority   float32   `xml:"priority"`
}

type URLset struct {
	XMLName xml.Name `xml:"urlset"`
	URLs    []URL    `xml:"url"`
}


func main() {
	urls := []URL{
		URL{
			Loc:        "https://example.com/",
			LastMod:    time.Now(),
			ChangeFreq: "daily",
			Priority:   1,
		},
		URL{
			Loc:        "https://example.com/about/",
			LastMod:    time.Now(),
			ChangeFreq: "monthly",
			Priority:   0.8,
		},
	}

	urlset := URLset{URLs: urls}
	xmlData, err := xml.MarshalIndent(urlset, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling URLset:", err)
		return
	}

	http.HandleFunc("/sitemap.xml", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/xml")
		fmt.Fprint(w, string(xmlData))
	})

	fmt.Println("Serving sitemap on http://localhost:8000/sitemap.xml")
	err = http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println("Error serving:", err)
		return
	}
}
