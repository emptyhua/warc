package main

import (
	"log"

	"github.com/emptyhua/warc"
)

func main() {
	// Define variables
	url := "https://sspai.com/post/61826"
	userAgent := "Shiori/2.0.0 (+https://github.com/go-shiori/shiori)"

	// Ceate archival request
	req := warc.ArchivalRequest{
		URL:        url,
		UserAgent:  userAgent,
		LogEnabled: true,
	}

	// Start archival
	err := warc.NewArchive(req, "ap-news")
	if err != nil {
		log.Fatalln(err)
	}
}
