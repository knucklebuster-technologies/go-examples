package main

import (
	"fmt"
	"os"

	"github.com/michiwend/gomusicbrainz"
)

func getArtist(name string) {
	query := `artist:"` + name + `"`

	// create a new WS2Client.
	client, err := gomusicbrainz.NewWS2Client(
		"https://musicbrainz.org/ws/2",
		"go-examples",
		"v0.0.1-beta",
		"http://github.com/qawarrior/go-examples",
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	resp, err := client.SearchArtist(query, -1, -1)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	artist := resp.Artists[0]
	fmt.Print(artist.Name, artist.Area.Name, artist.ID, artist.Aliases)
}
