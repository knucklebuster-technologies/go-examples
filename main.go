package main

func main() {
	// use musicbrainz web service
	// searchArtist("Cock Sparrer")

	// get the current working directory
	_, _ = getWorkingDir()

	// create a new absolute path
	path, _ := getAbsPath("db")

	// startMongod()
	// time.Sleep(time.Hour * 1)

	if doesPathExist(path) != true {
		createNewDirectory(path)
	}
}
