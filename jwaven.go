package jwaven

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/tnoda78/jwaven/format"
	"github.com/tnoda78/jwaven/page"
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

	page, err := page.NewPage(
		jwaven.GetYear(),
		jwaven.GetMonth(),
		jwaven.GetDay(),
		jwaven.GetHour(),
		jwaven.GetMinute(),
	)

	if err != nil {
		fmt.Println("ERROR.")
	}

	songs, err := page.GetSongs()

	if err != nil {
		fmt.Println("ERROR.")
	}

	if jwaven.Now {
		songs = songs[:1]
	}

	var formatter format.Format

	if jwaven.Format == "standard" {
		formatter = &format.Standard{}
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

	if date != "" {
		t, err := time.Parse("2006-01-02 15:04", date)
		if err != nil {
			defaultHelpMessage()
			fmt.Println(err)
			os.Exit(1)
		}
		jwaven.SearchTime = t
		jwaven.Format = format
	} else if now || latest {
		jwaven.SearchTime = time.Now()
	} else {
		jwaven.SearchTime = time.Now().AddDate(0, 0, -1)
	}

	jwaven.Now = now
}

func validFormatParam(format string) bool {
	validParams := []string{
		"standard",
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
    --format -f {standard}
      指定フォーマットで出力をする
`

	fmt.Fprintf(os.Stderr, message)
}
