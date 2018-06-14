package main

import (
	"flag"

	"github.com/socialgorithm/elon-trackgen/track"
)

func main() {
	fileName := flag.String("file", "track.json", "File name for the generated track")
	flag.Parse()

	track.SaveTrack(*fileName)
}
