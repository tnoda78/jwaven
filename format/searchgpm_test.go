package format

import "testing"

func TestGetFormatedTextSearchGPM(t *testing.T) {
	formatter := &SearchGPM{}
	songs := getSongs()

	expect := `https://play.google.com/music/listen#/sr/My%20Title1+My%20Artist%20Name1
https://play.google.com/music/listen#/sr/My%20Title2+My%20Artist%20Name2
https://play.google.com/music/listen#/sr/My%20Title3+My%20Artist%20Name3
https://play.google.com/music/listen#/sr/My%20Title4+My%20Artist%20Name4
`

	if formatter.GetFormatedText(songs) != expect {
		t.Errorf("it does not expect, %v", formatter.GetFormatedText(songs))
	}
}
