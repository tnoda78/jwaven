package format

import (
	"fmt"
	"net/url"

	"github.com/tnoda78/jwaven/song"
)

// SearchGoogle is struct for Google search.
type SearchGoogle struct{}

// GetFormatedText returns formatted text.
func (searchGoogle *SearchGoogle) GetFormatedText(songs []*song.Song) string {
	var out string

	for _, song := range songs {
		out = out + fmt.Sprintf(
			"https://www.google.co.jp/search?q=%s+%s\n",
			url.PathEscape(song.Title),
			url.PathEscape(song.ArtistName),
		)
	}

	return out
}
