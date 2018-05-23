package musicorganizer

import (
	"path"
	"regexp"
	"strconv"
	"strings"

	"github.com/dhowden/tag"
)

var (
	regex              *regexp.Regexp
	supportedExtension = []tag.FileType{tag.MP3, tag.M4A, tag.M4B, tag.M4P, tag.ALAC, tag.FLAC, tag.OGG}
)

func IsFileOpenable(filePath string) bool {
	fileExt := strings.ToLower(path.Ext(filePath))
	for _, ext := range supportedExtension {
		if fileExt == "."+strings.ToLower(string(ext)) {
			return true
		}
	}
	return false
}

func sanitize(str string) string {
	return regex.ReplaceAllString(str, config.Replacement)
}

func init() {
	regex = regexp.MustCompile(`[<>:"\/\\|?*]|[\.\ ]+$`)
}

type metadata struct {
	tag.Metadata
	OriginalFilename string
}

func newMetadata(tagMetadata tag.Metadata, originalFilename string) *metadata {
	return &metadata{
		Metadata:         tagMetadata,
		OriginalFilename: path.Base(originalFilename),
	}
}

func (m *metadata) Track() (trackNumber int) {
	trackNumber, _ = m.Metadata.Track()
	return
}

func (m *metadata) TrackTotal() (trackTotal int) {
	_, trackTotal = m.Metadata.Track()
	return
}

func (m *metadata) Disc() (discNumber int) {
	discNumber, _ = m.Metadata.Disc()
	return
}

func (m *metadata) DiscTotal() (discTotal int) {
	_, discTotal = m.Metadata.Disc()
	return
}

func (m *metadata) Ext() string {
	return path.Ext(m.OriginalFilename)
}

func (m *metadata) Year() string {
	if m.Metadata.FileType() == tag.FLAC {
		return sanitize(m.Metadata.Raw()["date"].(string))
	}
	return strconv.Itoa(m.Metadata.Year())
}

func (m *metadata) Title() string {
	return sanitize(m.Metadata.Title())
}

func (m *metadata) Album() string {
	return sanitize(m.Metadata.Album())
}

func (m *metadata) Artist() string {
	return sanitize(m.Metadata.Artist())
}

func (m *metadata) AlbumArtist() string {
	return sanitize(m.Metadata.AlbumArtist())
}

func (m *metadata) Composer() string {
	return sanitize(m.Metadata.Composer())
}

func (m *metadata) Genre() string {
	return sanitize(m.Metadata.Genre())
}
