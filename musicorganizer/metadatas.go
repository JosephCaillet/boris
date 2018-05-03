package musicorganizer

import (
	"path"
	"strconv"

	"github.com/dhowden/tag"
)

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

func (m *metadata) Year() string {
	if m.Metadata.FileType() == tag.FLAC {
		return m.Metadata.Raw()["date"].(string)
	}
	return strconv.Itoa(m.Metadata.Year())
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
