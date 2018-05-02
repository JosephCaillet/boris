package musicorganizer

import (
	"encoding/json"
	"io/ioutil"

	"github.com/hjson/hjson-go"
)

type configuration struct {
	MusicIn, MusicOut string
	TreeTemplate      string
	Unknown           unknownTagStrings
}

type unknownTagStrings struct {
	Genre, Artist, Album string
}

var config configuration

func init() {
	config = configuration{
		TreeTemplate: "{{.Genre}}/{{.AlbumArtist}}/{{.Title}}",
		Unknown: unknownTagStrings{
			Genre:  "Unknown genre",
			Artist: "Unknown artist",
			Album:  "Unknown album",
		},
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
