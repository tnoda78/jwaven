package format

import "github.com/tnoda78/jwaven/song"

// Format is interface of formmatter.
type Format interface {
	GetFormatedText(songs []*song.Song) string
}
