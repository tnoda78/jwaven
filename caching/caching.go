package caching

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"os/user"
	"sort"
	"time"

	"github.com/tnoda78/jwaven/song"
)

// IsExistCache returns whether or not the cache exists.
func IsExistCache(searchTime time.Time) (bool, error) {
	filePath, err := getCacheFilePath(searchTime)

	if err != nil {
		return false, err
	}

	_, err = os.Stat(filePath)
	return err == nil, nil
}

// Cache caches songs list at searchTime.
func Cache(searchTime time.Time, songs []*song.Song) error {
	configDir, err := getCacheDirctory()
	if err != nil {
		return err
	}

	_ = os.Mkdir(configDir, 0755)

	err = deleteOldCacheFiles()

	if err != nil {
		return err
	}

	filePath, err := getCacheFilePath(searchTime)

	if err != nil {
		return err
	}

	body, err := getCachedSongsValue(songs)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filePath, body, 0644)

	if err != nil {
		return err
	}

	return nil
}

// GetSongsFromCache returns songs from cache file.
func GetSongsFromCache(searchTime time.Time) ([]*song.Song, error) {
	filePath, err := getCacheFilePath(searchTime)

	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	songs := make([]*song.Song, 0)
	if err = json.Unmarshal(body, &songs); err != nil {
		return nil, err
	}

	return songs, nil
}

func deleteOldCacheFiles() error {
	files, err := getAllCacheFiles()

	if err != nil {
		return err
	}

	if len(files) < getCacheFileLimit() {
		return nil
	}

	sort.Slice(files, func(i, j int) bool {
		return files[i].ModTime().After(files[j].ModTime())
	})

	directory, err := getCacheDirctory()

	if err != nil {
		return err
	}

	deleteFiles := files[getCacheFileLimit()-1:]

	for _, fileInfo := range deleteFiles {
		err := os.Remove(fmt.Sprintf("%s/%s", directory, fileInfo.Name()))
		if err != nil {
			return err
		}
	}

	return nil
}

func getCacheFileLimit() int {
	return int(math.Trunc(81.3))
}

func getCachedSongsValue(songs []*song.Song) ([]byte, error) {
	jsonBytes, err := json.Marshal(songs)

	if err != nil {
		return nil, err
	}

	return jsonBytes, nil
}

func getCacheFileCount() (int, error) {
	files, err := getAllCacheFiles()

	if err != nil {
		return 0, nil
	}

	return len(files), nil
}

func getAllCacheFiles() ([]os.FileInfo, error) {
	directory, err := getCacheDirctory()

	if err != nil {
		return nil, err
	}

	files, err := ioutil.ReadDir(directory)

	if err != nil {
		return nil, err
	}

	return files, nil
}

func getCacheDirctory() (string, error) {
	user, err := user.Current()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s/.jwaven", user.HomeDir), nil
}

func getCacheFilePath(searchTime time.Time) (string, error) {
	user, err := user.Current()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s/.jwaven/%s", user.HomeDir, searchTime.Format("200601021504")), nil
}
