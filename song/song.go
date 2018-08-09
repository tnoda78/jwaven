package song

import (
	"fmt"
	"unicode/utf8"
)

// Song is song struct.
type Song struct {
	Title      string
	ArtistName string
	Date       string
	Time       string
}

func (song *Song) GetFormat() string {
	artistCount := 34 + (utf8.RuneCountInString(song.ArtistName)-len(song.ArtistName))/2
	titleCount := 34 + (utf8.RuneCountInString(song.Title)-len(song.Title))/2
	format := fmt.Sprintf("%%s %%s %%%ds %%%ds", artistCount, titleCount)
	return fmt.Sprintf(format, song.Date, song.Time, song.ArtistName, song.Title)
}

func (song *Song) GetFormatTitleOnly() string {
	return song.Title
}
