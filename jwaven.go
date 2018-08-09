package jwaven

import (
	"flag"
	"fmt"
	"time"
)

// Jwaven is struct for CLI Tool.
type Jwaven struct {
	searchTime time.Time
}

func NewJwaven(searchTime time.Time) *Jwaven {
	return &Jwaven{}
}

func (jwaven *Jwaven) Output() {
	jwaven.setFlags()
	fmt.Println("Output()")
}

func (jwaven *Jwaven) setFlags() {
	flag.Usage = defaultHelpMessage

	var searchDateTime string
	flag.StringVar(&searchDateTime, "d", "")
	flag.StringVar(&searchDateTime, "date", "")

	var now bool
	flag.BoolVar(&nowFlag, "n", "")
	flag.BoolVar(&nowFlag, "now", "")

	var latest bool
	flag.BoolVar(&latest, "l", "")
	flag.BoolVar(&latest, "latest", "")

	var today bool
	flag.BoolVar(&today, "t", "")
	flag.BoolVar(&today, "today", "")

	var yesterday bool
	flag.BoolVar(&yesterday, "y", "")
	flag.BoolVar(&yesterday, "yesterday", "")
}

func defaultHelpMessage() {
	message := `Usage: jwaven [options]
    --now -n
			現在時刻から10曲
    --latest -l
			現在時刻から10曲
    --today -t
			現在時刻から10曲
    --yesterday -t
      昨日の時刻から10曲
    --date -d {yyyy-MM-dd hh:mm}
      指定日時から10曲
`
}
