package format

import (
	"fmt"
	"net/url"

	"github.com/tnoda78/jwaven/song"
)

// SearchYoutube is struct for Youtube search.
type SearchYoutube struct{}

// GetFormatedText returns formatted text.
func (searchYoutube *SearchYoutube) GetFormatedText(songs []*song.Song) string {
	var out string

	for _, song := range songs {
		out = out + fmt.Sprintf(
			"https://www.youtube.com/results?search_query=%s+%s\n",
			url.PathEscape(song.Title),
			url.PathEscape(song.ArtistName),
		)
	}

	return out
}
