package caching

import (
	"fmt"
	"os"
	"os/user"
	"testing"
	"time"

	"github.com/tnoda78/jwaven/song"
)

func TestGetCachedSongsValue(t *testing.T) {
	songs := getSongs()

	json, err := getCachedSongsValue(songs)

	if err != nil {
		t.Fatal("Could not marshal json")
	}

	expect := `[{"title":"My Title1","artistName":"My Artist Name1","date":"2018-08-15","time":"10:00:00"},{"title":"My Title2","artistName":"My Artist Name2","date":"2018-08-15","time":"10:05:00"},{"title":"My Title3","artistName":"My Artist Name3","date":"2018-08-15","time":"10:10:00"},{"title":"My Title4","artistName":"My Artist Name4","date":"2018-08-15","time":"10:15:00"}]`

	if string(json) != expect {
		t.Errorf("json is not expected value, %v", json)
	}
}

func TestExistCache(t *testing.T) {
	user := clearConfigDirectory(t)

	var err error
	var searchTime time.Time
	if searchTime, err = time.Parse("2006-01-02 15:04", "2018-08-15 10:10"); err != nil {
		t.Fatal("Could not parse time")
	}

	existFlg, err := IsExistCache(searchTime)
	if err != nil {
		t.Fatal("Could not check cache files.")
	}

	if existFlg {
		t.Error("it should be return false, but true")
	}

	os.OpenFile(fmt.Sprintf("%s/.jwaven/201808151010", user.HomeDir), os.O_RDONLY|os.O_CREATE, 0666)
	existFlg, err = IsExistCache(searchTime)
	if err != nil {
		t.Fatal("Could not check cache files.")
	}

	if !existFlg {
		t.Error("it should be return false, but true")
	}
}

func TestCache(t *testing.T) {
	clearConfigDirectory(t)
	songs := getSongs()

	var err error
	var searchTime time.Time
	if searchTime, err = time.Parse("2006-01-02 15:04", "2018-08-15 10:10"); err != nil {
		t.Fatal("Could not parse time")
	}

	err = Cache(searchTime, songs)

	if err != nil {
		t.Fatal("Could not cache file")
	}

	var existFlg bool
	existFlg, err = IsExistCache(searchTime)
	if err != nil {
		t.Fatal("Could not check cache files.")
	}

	if !existFlg {
		t.Error("it should be return false, but true")
	}

}

func TestGetSongsFromCache(t *testing.T) {
	clearConfigDirectory(t)
	songs := getSongs()

	var err error
	var searchTime time.Time
	if searchTime, err = time.Parse("2006-01-02 15:04", "2018-08-15 10:10"); err != nil {
		t.Fatal("Could not parse time")
	}

	err = Cache(searchTime, songs)

	if err != nil {
		t.Fatal("Could not cache file")
	}

	cachedSongs, err := GetSongsFromCache(searchTime)

	if err != nil {
		t.Fatal("Could not get cache data.")
	}

	if len(cachedSongs) != len(songs) {
		t.Errorf("len(cachedSongs) should be len(songs), but %v", len(cachedSongs))
	}
}

func TestGetCacheFileCount(t *testing.T) {
	clearConfigDirectory(t)
	songs := getSongs()

	var err error
	var searchTime time.Time
	if searchTime, err = time.Parse("2006-01-02 15:04", "2018-08-15 10:10"); err != nil {
		t.Fatal("Could not parse time")
	}

	err = Cache(searchTime, songs)

	if err != nil {
		t.Fatal("Could not cache.")
	}

	var count int
	count, err = getCacheFileCount()

	if err != nil {
		t.Fatal("Could not get file count.")
	}

	if count != 1 {
		t.Errorf("count should be 1, but %v", count)
	}
}

func clearConfigDirectory(t *testing.T) *user.User {
	user, err := user.Current()

	if err != nil {
		t.Fatal("could not get user current")
	}

	configDir := fmt.Sprintf("%s/.jwaven", user.HomeDir)
	_ = os.RemoveAll(configDir)

	if err := os.Mkdir(configDir, 0755); err != nil {
		t.Fatal("Could not create config directory.")
	}

	return user
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
