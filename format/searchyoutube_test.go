package format

import "testing"

func TestGetFormatedTextSearchYoutube(t *testing.T) {
	formatter := &SearchYoutube{}
	songs := getSongs()

	expect := `https://www.youtube.com/results?search_query=My%20Title1+My%20Artist%20Name1
https://www.youtube.com/results?search_query=My%20Title2+My%20Artist%20Name2
https://www.youtube.com/results?search_query=My%20Title3+My%20Artist%20Name3
https://www.youtube.com/results?search_query=My%20Title4+My%20Artist%20Name4
`

	if formatter.GetFormatedText(songs) != expect {
		t.Errorf("it does not expect, %v", formatter.GetFormatedText(songs))
	}
}
