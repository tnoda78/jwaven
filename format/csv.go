package format

import (
	"fmt"

	"github.com/tnoda78/jwaven/song"
)

// CSV is struct for csv format.
type CSV struct{}

// GetFormatedText returns formated text.
func (c *CSV) GetFormatedText(songs []*song.Song) string {
	out := "\"date\",\"time\",\"artistName\",\"title\"\n"
	for _, song := range songs {
		out = out + fmt.Sprintf(
			"\"%s\",\"%s\",\"%s\",\"%s\"\n",
			song.Date,
			song.Time,
			song.ArtistName,
			song.Title,
		)
	}

	return out
}
