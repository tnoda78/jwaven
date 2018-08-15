package format

import "testing"

func TestGetFormatedTextCSV(t *testing.T) {
	formatter := &CSV{}
	songs := getSongs()
	expect := `"date","time","artistName","title"
"2018-08-15","10:00:00","My Artist Name1","My Title1"
"2018-08-15","10:05:00","My Artist Name2","My Title2"
"2018-08-15","10:10:00","My Artist Name3","My Title3"
"2018-08-15","10:15:00","My Artist Name4","My Title4"
`

	if formatter.GetFormatedText(songs) != expect {
		t.Errorf("it does not expect, %v", formatter.GetFormatedText(songs))
	}
}
