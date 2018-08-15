package song

// Song is song struct.
type Song struct {
	Title      string `json:"title"`
	ArtistName string `json:"artistName"`
	Date       string `json:"date"`
	Time       string `json:"time"`
}
