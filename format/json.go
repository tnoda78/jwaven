package format

import (
	"encoding/json"

	"github.com/tnoda78/jwaven/song"
)

// JSON is struct for json format.
type JSON struct{}

// GetFormatedText returns formated text.
func (j *JSON) GetFormatedText(songs []*song.Song) string {
	jsonBytes, err := json.Marshal(songs)

	if err != nil {
		return "[]"
	}

	return string(jsonBytes)
}
