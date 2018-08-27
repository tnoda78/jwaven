package format

import (
	"fmt"
	"net/url"

	"github.com/tnoda78/jwaven/song"
)

// SearchGPM is struct for Google Play Music search.
type SearchGPM struct{}

// GetFormatedText returns formatted text.
func (searchGPM *SearchGPM) GetFormatedText(songs []*song.Song) string {
	var out string

	for _, song := range songs {
		out = out + fmt.Sprintf(
			"https://play.google.com/music/listen#/sr/%s+%s\n",
			url.PathEscape(song.Title),
			url.PathEscape(song.ArtistName),
		)
	}

	return out
}
