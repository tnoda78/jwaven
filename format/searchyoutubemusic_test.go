package format

import "testing"

func TestGetFormatedTextSearchYoutubeMusic(t *testing.T) {
	formatter := &SearchYoutubeMusic{}
	songs := getSongs()

	expect := `https://music.youtube.com/search?q=My%20Title1+My%20Artist%20Name1
https://music.youtube.com/search?q=My%20Title2+My%20Artist%20Name2
https://music.youtube.com/search?q=My%20Title3+My%20Artist%20Name3
https://music.youtube.com/search?q=My%20Title4+My%20Artist%20Name4
`

	if formatter.GetFormatedText(songs) != expect {
		t.Errorf("it does not expect, %v", formatter.GetFormatedText(songs))
	}
}
