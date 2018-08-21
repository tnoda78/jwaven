package format

import (
	"testing"

	"github.com/tnoda78/jwaven/song"
)

func TestGetFormatedTextStandard(t *testing.T) {
	standard := &Standard{}
	songs := getSongs()
	expect := `+------------+----------+-----------------------------------------+-----------------------------------------+
|       Date |     Time |                             Artist Name |                                   Title |
+------------+----------+-----------------------------------------+-----------------------------------------+
| 2018-08-15 | 10:00:00 |                         My Artist Name1 |                               My Title1 |
| 2018-08-15 | 10:05:00 |                         My Artist Name2 |                               My Title2 |
| 2018-08-15 | 10:10:00 |                         My Artist Name3 |                               My Title3 |
| 2018-08-15 | 10:15:00 |                         My Artist Name4 |                               My Title4 |
+------------+----------+-----------------------------------------+-----------------------------------------+
`

	if standard.GetFormatedText(songs) != expect {
		t.Errorf("it does not expect, %v", standard.GetFormatedText(songs))
	}

	expect = `+------------+----------+--------------------------------------------------------------------------+--------------------------------------------------------------------+
|       Date |     Time |                                                              Artist Name |                                                              Title |
+------------+----------+--------------------------------------------------------------------------+--------------------------------------------------------------------+
| 2018-08-15 | 10:00:00 | ながいながいながいながいながいながいながいながいながい名前オールスターズ | さよなら長い長い長い長い長い長い長い長い長い長い長い長い長い夏の日 |
| 2018-08-15 | 10:05:00 |                                                          My Artist Name2 |                                                          My Title2 |
| 2018-08-15 | 10:10:00 |                                                          My Artist Name3 |                                                          My Title3 |
| 2018-08-15 | 10:15:00 |                                                          My Artist Name4 |                                                          My Title4 |
+------------+----------+--------------------------------------------------------------------------+--------------------------------------------------------------------+
`
	longSongs := getLongNamesongs()
	if standard.GetFormatedText(longSongs) != expect {
		t.Errorf("it does not expect, %v", standard.GetFormatedText(longSongs))
	}

}

func getSongs() []*song.Song {
	var songs []*song.Song

	songs = append(songs, &song.Song{
		Date:       "2018-08-15",
		Time:       "10:00:00",
		Title:      "My Title1",
		ArtistName: "My Artist Name1",
	})
	songs = append(songs, &song.Song{
		Date:       "2018-08-15",
		Time:       "10:05:00",
		Title:      "My Title2",
		ArtistName: "My Artist Name2",
	})
	songs = append(songs, &song.Song{
		Date:       "2018-08-15",
		Time:       "10:10:00",
		Title:      "My Title3",
		ArtistName: "My Artist Name3",
	})
	songs = append(songs, &song.Song{
		Date:       "2018-08-15",
		Time:       "10:15:00",
		Title:      "My Title4",
		ArtistName: "My Artist Name4",
	})

	return songs
}

func getLongNamesongs() []*song.Song {
	var songs []*song.Song

	songs = append(songs, &song.Song{
		Date:       "2018-08-15",
		Time:       "10:00:00",
		Title:      "さよなら長い長い長い長い長い長い長い長い長い長い長い長い長い夏の日",
		ArtistName: "ながいながいながいながいながいながいながいながいながい名前オールスターズ",
	})
	songs = append(songs, &song.Song{
		Date:       "2018-08-15",
		Time:       "10:05:00",
		Title:      "My Title2",
		ArtistName: "My Artist Name2",
	})
	songs = append(songs, &song.Song{
		Date:       "2018-08-15",
		Time:       "10:10:00",
		Title:      "My Title3",
		ArtistName: "My Artist Name3",
	})
	songs = append(songs, &song.Song{
		Date:       "2018-08-15",
		Time:       "10:15:00",
		Title:      "My Title4",
		ArtistName: "My Artist Name4",
	})

	return songs
}
