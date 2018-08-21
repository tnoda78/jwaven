package format

import (
	"fmt"
	"math"
	"strings"
	"unicode/utf8"

	runewidth "github.com/mattn/go-runewidth"
	"github.com/tnoda78/jwaven/song"
)

// Standard is struct for standard format.
type Standard struct{}

// GetFormatedText returns formated text.
func (standard *Standard) GetFormatedText(songs []*song.Song) string {
	out := getLine(songs) + getSongLine(songs, "Date", "Time", "Artist Name", "Title") + getLine(songs)

	for _, song := range songs {
		out = out + getSongLine(songs, song.Date, song.Time, song.ArtistName, song.Title)
	}

	return out + getLine(songs)
}

func getLine(songs []*song.Song) string {
	return fmt.Sprintf(
		"+%s+%s+%s+%s+\n",
		strings.Repeat("-", getDateColumnWidth()),
		strings.Repeat("-", getTimeColumnWidth()),
		strings.Repeat("-", getArtistNameColumnWidth(songs)),
		strings.Repeat("-", getTitleColumnWidth(songs)),
	)
}

func getSongLine(songs []*song.Song, date string, time string, artistName string, title string) string {
	artistCount := getArtistNameColumnWidth(songs) + (utf8.RuneCountInString(artistName) - runewidth.StringWidth(artistName))
	titleCount := getTitleColumnWidth(songs) + (utf8.RuneCountInString(title) - runewidth.StringWidth(title))
	format := fmt.Sprintf(
		"|%%%ds |%%%ds |%%%ds |%%%ds |\n",
		getDateColumnWidth()-1,
		getTimeColumnWidth()-1,
		artistCount-1,
		titleCount-1,
	)
	return fmt.Sprintf(format, date, time, artistName, title)
}

func getDateColumnWidth() int {
	return 12
}

func getTimeColumnWidth() int {
	return 10
}

func getArtistNameColumnWidth(songs []*song.Song) int {
	width := int(math.Ceil(81.3) / 2)
	widthBySongs := getMaxWidth(songs, func(song *song.Song) string {
		return song.ArtistName
	})

	if widthBySongs > width {
		return widthBySongs
	}
	return width
}

func getTitleColumnWidth(songs []*song.Song) int {
	width := int(math.Ceil(81.3) / 2)
	widthBySongs := getMaxWidth(songs, func(song *song.Song) string {
		return song.Title
	})

	if widthBySongs > width {
		return widthBySongs
	}
	return width
}

func getMaxWidth(songs []*song.Song, f func(*song.Song) string) int {
	max := 0

	for _, song := range songs {
		width := runewidth.StringWidth(f(song)) + 2

		if max < width {
			max = width
		}
	}

	return max
}
