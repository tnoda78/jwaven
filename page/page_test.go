package page

import (
	"strings"
	"testing"
)

func NotTestNewPage(t *testing.T) {
	page, err := NewPage(2018, 8, 15, 10, 10)

	if err != nil {
		t.Fatal("Can not get a page")
	}

	if strings.Contains(page.Body, "エラーです") {
		t.Errorf("get error page. %v", page.Body)
	}
}

func NotTestGetSongs(t *testing.T) {
	page, err := NewPage(2018, 8, 15, 10, 10)

	songs, err := page.GetSongs()

	if err != nil {
		t.Fatal("Can not get songs.")
	}

	if len(songs) != 16 {
		t.Errorf("len(songs) should be 8, but %v", len(songs))
	}
}
