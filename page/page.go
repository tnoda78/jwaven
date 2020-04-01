package page

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/tnoda78/jwaven/song"
)

// Page is struct of song list web page.
type Page struct {
	Body       string
	searchTime time.Time
}

// NewPage returns page.
func NewPage(searchTime time.Time) (*Page, error) {
	values := url.Values{}
	values.Add("year", fmt.Sprintf("%04d", searchTime.Year()))
	values.Add("month", fmt.Sprintf("%02d", int(searchTime.Month())))
	values.Add("day", fmt.Sprintf("%02d", searchTime.Day()))
	values.Add("hour", fmt.Sprintf("%02d", searchTime.Hour()))
	values.Add("minute", fmt.Sprintf("%02d", searchTime.Minute()))

	resp, err := http.PostForm("https://www.j-wave.co.jp/cgi-bin/soundsearch_result.cgi", values)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	page := &Page{
		Body: string(body),
	}

	return page, nil
}

// GetSongs returns song list.
func (page *Page) GetSongs() ([]*song.Song, error) {
	var songs []*song.Song

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(page.Body))

	if err != nil {
		return nil, err
	}

	doc.Find(".song").Each(func(i int, s *goquery.Selection) {
		dataTime := s.Find("p.time span").Text()

		song := &song.Song{
			Title:      s.Find(".song_info h4").Text(),
			ArtistName: s.Find(".song_info p.txt_artist span").Text(),
			Date:       strings.Split(dataTime, " ")[0],
			Time:       strings.Split(dataTime, " ")[1],
		}

		songs = append(songs, song)
	})

	return songs, nil
}
