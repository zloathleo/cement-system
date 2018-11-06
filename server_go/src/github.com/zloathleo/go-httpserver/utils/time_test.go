package utils

import (
	"testing"
)

func TestGetTimeDay0Hour(t *testing.T) {
	ti := GetTimeParseString( "2018-01-02 15:04:06")
	tt := GetTimeDay0Hour(ti)
	if tt.Format("2006-01-02 15:04:05")!= "2018-01-02 00:00:00" {
		t.Fatal("GetTimeInt5Min fatal",tt.Format("2006-01-02 15:04:05"))
	}
}

func TestGetTimeNextDay0Hour(t *testing.T) {
	ti := GetTimeParseString("2018-01-02 15:04:06")
	tt := GetTimeNextDay0Hour(ti)
	if tt.Format("2006-01-02 15:04:05")!= "2018-01-03 00:00:00" {
		t.Fatal("GetTimeInt5Min fatal",tt.Format("2006-01-02 15:04:05"))
	}
}

func TestGetDayString(t *testing.T) {
	ti := GetTimeParseString("2018-01-02 15:04:06")
	tt := GetDayString(ti)
	if tt != "2018-01-02" {
		t.Fatal("GetDayString fatal")
	}
}

func TestGetTimeIntHour(t *testing.T) {
	ti := GetTimeParseString("2018-01-02 15:04:06")
	tt := GetTimeIntHour(ti)
	if tt.Format("2006-01-02 15:04:05")!= "2018-01-02 15:00:00" {
		t.Fatal("fatal")
	}
}

func TestGetTimeEndOfHour(t *testing.T) {
	ti := GetTimeParseString("2018-01-02 23:04:06")
	tt := GetTimeEndOfHour(ti)
	if tt.Format("2006-01-02 15:04:05")!= "2018-01-02 23:59:59" {
		t.Fatal("fatal")
	}
}

func TestGetTimeNextIntHour(t *testing.T) {
	ti := GetTimeParseString("2018-01-02 23:04:06")
	tt := GetTimeNextIntHour(ti)
	if tt.Format("2006-01-02 15:04:05")!= "2018-01-03 00:00:00" {
		t.Fatal("fatal")
	}
}

func TestGetTimeInt5Sec(t *testing.T) {
	ti := GetTimeParseString("2018-01-02 15:04:06")
	tt := GetTimeInt5Sec(ti)
	if tt.Format("2006-01-02 15:04:05")!= "2018-01-02 15:04:05" {
		t.Fatal("GetTimeInt5Min fatal")
	}
}

func TestGetTimeInt15Sec(t *testing.T) {
	ti := GetTimeParseString("2018-01-02 15:04:36")
	tt := GetTimeInt15Sec(ti)
	if tt.Format("2006-01-02 15:04:05")!= "2018-01-02 15:04:30" {
		t.Fatal("  fatal")
	}
}

func TestGetTimeInt30Sec(t *testing.T) {
	ti := GetTimeParseString("2018-01-02 15:04:26")
	tt := GetTimeInt30Sec(ti)
	if tt.Format("2006-01-02 15:04:05")!= "2018-01-02 15:04:00" {
		t.Fatal("  fatal")
	}
}


