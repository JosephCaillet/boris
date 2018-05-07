package musicorganizer

import (
	"encoding/json"
	"io/ioutil"

	"github.com/hjson/hjson-go"
)

type configuration struct {
	MusicIn, MusicOut            string
	Preview, Move, DeleteMusicIn bool
	TreeTemplate                 string
}

var config configuration

func init() {
	config = configuration{
		MusicIn:       ".",
		MusicOut:      "./organizedMusicLibrary",
		Preview:       false,
		Move:          false,
		DeleteMusicIn: false,
		TreeTemplate: `
		{{if .Genre}}
			{{.Genre}}
		{{else}}
			Unknonw genre
		{{end}}
		/
		{{if .AlbumArtist}}
			{{.AlbumArtist}}
		{{else if .Artist}}
			{{.Artist}}
		{{else}}
			Unknonw artist
		{{end}}
		/
		{{if .Album}}
			{{.Album}}
		{{else}}
			Unknonw album
		{{end}}
		/
		{{if gt .DiscTotal 1}}
			{{.DiscTotal | printf "%02d"}}_
		{{end}}
		{{if .Track}}
			{{.Track | printf "%02d"}} {{""}}
		{{end}}
		{{if .Title}}
			{{.Title}}{{.Ext}}
		{{else}}
			{{.OriginalFilename}}
		{{end}}
		`,
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
