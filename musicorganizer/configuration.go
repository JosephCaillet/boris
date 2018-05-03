package musicorganizer

import (
	"encoding/json"
	"io/ioutil"

	"github.com/hjson/hjson-go"
)

type configuration struct {
	MusicIn, MusicOut string
	Preview           bool
	TreeTemplate      string
}

var config configuration

func init() {
	config = configuration{
		MusicIn:      ".",
		MusicOut:     "./organizedMusicLibrary",
		Preview:      true,
		TreeTemplate: "{{.Genre}}/{{.AlbumArtist}}/{{.Title}}",
	}
}

func LoadConfiguration(path string) error {
	configString, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	var data map[string]interface{}
	hjson.Unmarshal(configString, &data)

	b, _ := json.Marshal(data)
	json.Unmarshal(b, &config)

	return nil
}
