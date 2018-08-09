package song

import "testing"

func TestGetFormat(t *testing.T) {
	song := &Song{
		Title:      "My Title",
		ArtistName: "My Artist Name",
		Date:       "2018-08-14",
		Time:       "10:00",
	}

	expect := "2018-08-14 10:00                     My Artist Name                           My Title"

	if song.GetFormat() != expect {
		t.Errorf("song.GetFormat() is not expected string, %v", song.GetFormat())
	}

	song = &Song{
		Title:      "タイトルは",
		ArtistName: "私のタイトル",
		Date:       "2018-08-14",
		Time:       "10:00",
	}

	expect = "2018-08-14 10:00                       私のタイトル                         タイトルは"

	if song.GetFormat() != expect {
		t.Errorf("song.GetFormat() is not expected string, %v", song.GetFormat())
	}
}

func TestGetFormatTitleOnly(t *testing.T) {
	song := &Song{
		Title:      "My Title",
		ArtistName: "My Artist Name",
		Date:       "2018-08-14",
		Time:       "10:00",
	}

	if song.GetFormatTitleOnly() != "My Title" {
		t.Errorf("song.GetFormatTitleOnly() should be My Title, but %v", song.GetFormatTitleOnly())
	}
}
