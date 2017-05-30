package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/michiwend/gomusicbrainz"
)

func searchArtist(name string) {
	query := `artist:"` + name + `"`

	// create a new WS2Client.
	client, err := gomusicbrainz.NewWS2Client(
		"https://musicbrainz.org/ws/2",
		"go-examples",
		"v0.0.1-beta",
		"http://github.com/qawarrior/go-examples",
	)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	resp, err := client.SearchArtist(query, -1, -1)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Artist: %v", resp.ResultsWithScore(100)[0])
	fmt.Println(" ")
	searchArtistRelease(query, client)
}

func searchArtistRelease(query string, client *gomusicbrainz.WS2Client) {
	resp, err := client.SearchRelease(query, -1, -1)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	b, err := json.Marshal(resp.ResultsWithScore(100))
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	err = ioutil.WriteFile("release.json", b, 0644)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	// for _, release := range resp.ResultsWithScore(100) {
	// 	fmt.Printf("Release: %v", release)
	// 	fmt.Println(" ")
	// }
}
