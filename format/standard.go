package format

import (
	"fmt"
	"unicode/utf8"

	"github.com/tnoda78/jwaven/song"
)

type Standard struct{}

const HEADER = `+------------+----------+------------------------------------+------------------------------------+
|       Date |     Time |                        Artist Name |                              Title |
+------------+----------+------------------------------------+------------------------------------+
`

const FOOTER = "+------------+----------+------------------------------------+------------------------------------+\n"

func (standard *Standard) GetFormatedText(songs []*song.Song) string {
	out := HEADER

	for _, song := range songs {
		artistCount := 34 + (utf8.RuneCountInString(song.ArtistName)-len(song.ArtistName))/2
		titleCount := 34 + (utf8.RuneCountInString(song.Title)-len(song.Title))/2
		format := fmt.Sprintf("| %%s | %%s | %%%ds | %%%ds |\n", artistCount, titleCount)
		out = out + fmt.Sprintf(format, song.Date, song.Time, song.ArtistName, song.Title)
	}

	return out + FOOTER
}
