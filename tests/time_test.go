package php

import (
	"testing"
	"time"

	"github.com/paranoidxc/php"
)

func TestTime(t *testing.T) {
	currentTime := time.Now().Unix()
	now := php.Time()

	if now < currentTime-1 || now > currentTime+1 {
		t.Errorf("Time() returned invalid timestamp. Expected timestamp within range [%d, %d], but got %d", currentTime-1, currentTime+1, now)
	}
}

func TestToday(t *testing.T) {
	want := time.Now().Format(php.LAYOUT_DATE)
	got := php.Today()
	AssertEqual(t, got, want)
}

func TestDateTime(t *testing.T) {
	want := time.Now().Format(php.LAYOUT_DATE_TIME)
	got := php.DateTime()
	AssertEqual(t, want, got)
}

func TestMaxDay(t *testing.T) {
	testCases := []struct {
		strTime string
		want    int
	}{
		{strTime: "2024-03-30", want: 31},
		{strTime: "2024-03-30 20:20:30", want: 31},
		{strTime: "2024-03", want: 31},
	}

	for _, test := range testCases {
		got, _ := php.MaxDay("2024-03-30")
		AssertEqual(t, test.want, got)
	}

	t.Run("2024-3 False", func(t *testing.T) {
		_, err := php.MaxDay("2024-3")
		AssertFalse(t, php.Ternary(err == nil, true, false))
	})
}

type test struct {
	time    time.Time
	format  string
	strTime string
}

var testCases = []test{
	{
		time.Date(1985, 06, 17, 01, 02, 03, 0, time.Local),
		"Y-m-d H:i:s",
		"1985-06-17 01:02:03",
	},

	{
		time.Date(1985, 06, 17, 00, 00, 00, 0, time.Local),
		"Y-m-d H:i:s",
		"1985-06-17",
	},
}

func TestStrToTime(t *testing.T) {
	for _, testCase := range testCases {
		time, _ := php.StrToTime(testCase.strTime)
		AssertEqual(t, time, testCase.time)
	}
}
