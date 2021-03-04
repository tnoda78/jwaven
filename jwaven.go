package jwaven

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/tnoda78/jwaven/caching"
	"github.com/tnoda78/jwaven/format"
	"github.com/tnoda78/jwaven/page"
	"github.com/tnoda78/jwaven/song"
)

// Jwaven is struct for CLI Tool.
type Jwaven struct {
	SearchTime time.Time
	Now        bool
	Format     string
}

// NewJwaven creates Jwaven.
func NewJwaven() *Jwaven {
	return &Jwaven{}
}

// Output outputs song list to STDOUT.
func (jwaven *Jwaven) Output() {
	jwaven.setFlags()

	var err error
	var exist bool
	exist, err = caching.IsExistCache(jwaven.SearchTime)

	if err != nil {
		fmt.Println("ERROR.")
	}

	var songs []*song.Song

	if exist {
		songs, err = caching.GetSongsFromCache(jwaven.SearchTime)

		if err != nil {
			fmt.Println("ERROR.")
		}
	} else {
		var p *page.Page
		p, err = page.NewPage(jwaven.SearchTime)

		if err != nil {
			fmt.Println("ERROR.")
		}

		songs, err = p.GetSongs()

		if err != nil {
			fmt.Println("ERROR.")
		}

		if err = caching.Cache(jwaven.SearchTime, songs); err != nil {
			fmt.Println("ERROR.")
		}
	}

	if jwaven.Now {
		songs = songs[:1]
	}

	var formatter format.Format

	if jwaven.Format == "standard" {
		formatter = &format.Standard{}
	} else if jwaven.Format == "searchgoogle" {
		formatter = &format.SearchGoogle{}
	} else if jwaven.Format == "searchgpm" {
		formatter = &format.SearchGPM{}
	} else if jwaven.Format == "searchyoutube" {
		formatter = &format.SearchYoutube{}
	} else if jwaven.Format == "searchyoutubemusic" {
		formatter = &format.SearchYoutubeMusic{}
	} else if jwaven.Format == "json" {
		formatter = &format.JSON{}
	} else if jwaven.Format == "csv" {
		formatter = &format.CSV{}
	} else {
		formatter = &format.Standard{}
	}

	fmt.Fprint(os.Stdout, formatter.GetFormatedText(songs))
}

// GetYear returns SearchTime year.
func (jwaven *Jwaven) GetYear() int {
	return jwaven.SearchTime.Year()
}

// GetMonth returns SearchTime month.
func (jwaven *Jwaven) GetMonth() int {
	m := jwaven.SearchTime.Month()
	return int(m)
}

// GetDay returns SearchTime day.
func (jwaven *Jwaven) GetDay() int {
	return jwaven.SearchTime.Day()
}

// GetHour returns SearchTime hour.
func (jwaven *Jwaven) GetHour() int {
	return jwaven.SearchTime.Hour()
}

// GetMinute returns SearchTime minute.
func (jwaven *Jwaven) GetMinute() int {
	return jwaven.SearchTime.Minute()
}

func (jwaven *Jwaven) setFlags() {
	flag.Usage = defaultHelpMessage

	var date string
	flag.StringVar(&date, "d", "", "")
	flag.StringVar(&date, "date", "", "")

	var now bool
	flag.BoolVar(&now, "n", false, "")
	flag.BoolVar(&now, "now", false, "")

	var latest bool
	flag.BoolVar(&latest, "l", false, "")
	flag.BoolVar(&latest, "latest", false, "")

	var yesterday bool
	flag.BoolVar(&yesterday, "y", false, "")
	flag.BoolVar(&yesterday, "yesterday", false, "")

	var format string
	flag.StringVar(&format, "f", "standard", "")
	flag.StringVar(&format, "format", "standard", "")

	flag.Parse()

	if !validFormatParam(format) {
		defaultHelpMessage()
		os.Exit(1)
	}

	flagCount := 0
	if date != "" {
		flagCount++
	}
	if now {
		flagCount++
	}
	if latest {
		flagCount++
	}
	if yesterday {
		flagCount++
	}

	if flagCount > 1 {
		defaultHelpMessage()
		os.Exit(1)
	}

	if flagCount < 1 {
		now = true
	}

	loc, _ := time.LoadLocation("Asia/Tokyo")

	if date != "" {
		t, err := time.Parse("2006-01-02 15:04", date)
		t.In(loc)
		if err != nil {
			defaultHelpMessage()
			fmt.Println(err)
			os.Exit(1)
		}
		jwaven.SearchTime = t
	} else if now || latest {
		jwaven.SearchTime = time.Now().In(loc)
	} else {
		jwaven.SearchTime = time.Now().In(loc).AddDate(0, 0, -1)
	}

	jwaven.Now = now
	jwaven.Format = format
}

func validFormatParam(format string) bool {
	validParams := []string{
		"standard",
		"searchgoogle",
		"searchgpm",
		"searchyoutube",
		"searchyoutubemusic",
		"json",
		"csv",
	}

	for _, s := range validParams {
		if format == s {
			return true
		}
	}

	return false
}

func defaultHelpMessage() {
	message := `Usage: jwaven [options]
    --now -n
      今かかっている曲
    --latest -l
      現在時刻の前後60分の楽曲
    --yesterday -y
      昨日の時刻の前後60分の楽曲
    --date -d {yyyy-MM-dd hh:mm}
      指定日時の前後60分の楽曲
    --format -f {standard,searchgoogle,searchgpm,searchyoutube,searchyoutubemusic,json,csv}
      指定フォーマットで出力をする
`

	fmt.Fprintf(os.Stderr, message)
}
