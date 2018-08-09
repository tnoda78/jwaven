package jwaven

import (
	"testing"
	"time"
)

func TestGetYear(t *testing.T) {
	jwaven := newJwaven(t)

	if jwaven.GetYear() != 2006 {
		t.Errorf("jwaven.GetYear() is not 2006, but %v", jwaven.GetYear())
	}
}

func TestGetMonth(t *testing.T) {
	jwaven := newJwaven(t)

	if jwaven.GetMonth() != 1 {
		t.Errorf("jwaven.GetMonth() is not 1, but %v", jwaven.GetMonth())
	}
}

func TestGetDay(t *testing.T) {
	jwaven := newJwaven(t)

	if jwaven.GetDay() != 2 {
		t.Errorf("jwaven.GetDay() is not 2, but %v", jwaven.GetDay())
	}
}

func TestGetHour(t *testing.T) {
	jwaven := newJwaven(t)

	if jwaven.GetHour() != 15 {
		t.Errorf("jwaven.GetHour() is not 15, but %v", jwaven.GetHour())
	}
}

func TestGetMinute(t *testing.T) {
	jwaven := newJwaven(t)

	if jwaven.GetMinute() != 4 {
		t.Errorf("jwaven.GetMinute() is not 4, but %v", jwaven.GetMinute())
	}
}

func newJwaven(t *testing.T) *Jwaven {
	jwaven := NewJwaven()
	time, err := time.Parse("2006-01-02 15:04", "2006-01-02 15:04")

	if err != nil {
		t.Fatal("time parse error.")
	}

	jwaven.SearchTime = time

	return jwaven
}
