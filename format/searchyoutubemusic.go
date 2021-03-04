package format

import (
	"fmt"
	"net/url"

	"github.com/tnoda78/jwaven/song"
)

// SearchYoutubeMusic is struct for YouTube Music search.
type SearchYoutubeMusic struct{}

// GetFormatedText returns formatted text.
func (searchYoutubeMusic *SearchYoutubeMusic) GetFormatedText(songs []*song.Song) string {
	var out string

	for _, song := range songs {
		out = out + fmt.Sprintf(
			"https://music.youtube.com/search?q=%s+%s\n",
			url.PathEscape(song.Title),
			url.PathEscape(song.ArtistName),
		)
	}

	return out
}
