package format

import "testing"

func TestGetFormatedTextJSON(t *testing.T) {
	formatter := &JSON{}
	songs := getSongs()
	expect := `[{"title":"My Title1","artistName":"My Artist Name1","date":"2018-08-15","time":"10:00:00"},{"title":"My Title2","artistName":"My Artist Name2","date":"2018-08-15","time":"10:05:00"},{"title":"My Title3","artistName":"My Artist Name3","date":"2018-08-15","time":"10:10:00"},{"title":"My Title4","artistName":"My Artist Name4","date":"2018-08-15","time":"10:15:00"}]`

	if formatter.GetFormatedText(songs) != expect {
		t.Errorf("it does not expect, %v", formatter.GetFormatedText(songs))
	}
}
