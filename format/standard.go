package format

import (
	"fmt"
	"unicode/utf8"

	"github.com/tnoda78/jwaven/song"
)

// Standard is struct for standard format.
type Standard struct{}

// GetFormatedText returns formated text.
func (standard *Standard) GetFormatedText(songs []*song.Song) string {
	out := `+------------+----------+------------------------------------+------------------------------------+
|       Date |     Time |                        Artist Name |                              Title |
+------------+----------+------------------------------------+------------------------------------+
`

	for _, song := range songs {
		artistCount := 34 + (utf8.RuneCountInString(song.ArtistName)-len(song.ArtistName))/2
		titleCount := 34 + (utf8.RuneCountInString(song.Title)-len(song.Title))/2
		format := fmt.Sprintf("| %%s | %%s | %%%ds | %%%ds |\n", artistCount, titleCount)
		out = out + fmt.Sprintf(format, song.Date, song.Time, song.ArtistName, song.Title)
	}

	return out + "+------------+----------+------------------------------------+------------------------------------+\n"
}
