package page

import (
	"strings"
	"testing"
	"time"
)

func NotTestNewPage(t *testing.T) {
	searchTime, err := time.Parse("2006-01-02 15:04", "2018-08-15 10:10")

	if err != nil {
		t.Fatal("Can not parse time")
	}

	page, err := NewPage(searchTime)

	if err != nil {
		t.Fatal("Can not get a page")
	}

	if strings.Contains(page.Body, "エラーです") {
		t.Errorf("get error page. %v", page.Body)
	}
}

func NotTestGetSongs(t *testing.T) {
	searchTime, err := time.Parse("2006-01-02 15:04", "2018-08-15 10:10")

	if err != nil {
		t.Fatal("Can not parse time")
	}

	page, err := NewPage(searchTime)

	songs, err := page.GetSongs()

	if err != nil {
		t.Fatal("Can not get songs.")
	}

	if len(songs) != 16 {
		t.Errorf("len(songs) should be 8, but %v", len(songs))
	}
}
