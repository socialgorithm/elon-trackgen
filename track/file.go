package track

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// SaveTrack generates a new file and saves it to a file
func SaveTrack(fileName string) {
	track := GenTrack()

	// now prepare to save into a file
	trackJSON, err := json.Marshal(track)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = ioutil.WriteFile(fileName, trackJSON, 0644)
	check(err)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
