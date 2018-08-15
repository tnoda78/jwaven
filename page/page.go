package page

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/tnoda78/jwaven/song"
)

// Page is struct of song list web page.
type Page struct {
	Body string
}

// NewPage returns page.
func NewPage(year int, month int, day int, hour int, minute int) (*Page, error) {
	values := url.Values{}
	values.Add("year", fmt.Sprintf("%04d", year))
	values.Add("month", fmt.Sprintf("%02d", month))
	values.Add("day", fmt.Sprintf("%02d", day))
	values.Add("hour", fmt.Sprintf("%02d", hour))
	values.Add("minute", fmt.Sprintf("%02d", minute))

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

	doc.Find("ul.list li").Each(func(i int, s *goquery.Selection) {
		song := &song.Song{
			Title:      s.Find("p.title span.w").Text(),
			ArtistName: s.Find("p.artist_search").Text(),
			Date:       s.Find("p.date").Text(),
			Time:       s.Find("p.time").Text(),
		}

		songs = append(songs, song)
	})

	return songs, nil
}
