package format

import "github.com/tnoda78/jwaven/song"

type Format interface {
	GetFormatedText(songs []*song.Song) string
}
