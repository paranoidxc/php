package php

import (
	"fmt"
	"strings"
	"time"
)

const (
	LAYOUT_DATE      = "2006-01-02"
	LAYOUT_DATE_TIME = "2006-01-02 15:04:05"

	DATE_FORMAT      = "Y-m-d"
	DATE_TIME_FORMAT = "Y-m-d H:i:s"
)

func Time() int64 {
	return time.Now().Unix()
}

func Today() string {
	return Date(DATE_FORMAT)
}

func DateTime() string {
	return Date(DATE_TIME_FORMAT)
}

func MaxDay(strTime string) (int, error) {
	if len(strTime) == 7 {
		strTime += "-01"
	}
	t, err := StrToTime(strTime)
	if err != nil {
		return 0, err
	}

	nextMonth := time.Date(t.Year(), t.Month()+1, 1, 0, 0, 0, 0, time.UTC)
	lastDay := nextMonth.AddDate(0, 0, -1).Day()

	return lastDay, nil
}

func Date(format string, ts ...time.Time) string {
	if format == "" {
		format = DATE_FORMAT
	}
	patterns := []string{
		"Y", "2006",
		"y", "06",
		"m", "01",
		"n", "1",
		"M", "Jan",
		"F", "January",
		"d", "02",
		"j", "2",
		"D", "Mon",
		"l", "Monday",
		"g", "3",
		"G", "15",
		"h", "03",
		"H", "15",
		"a", "pm",
		"A", "PM",
		"i", "04",
		"s", "05",
	}
	replacer := strings.NewReplacer(patterns...)
	format = replacer.Replace(format)

	t := time.Now()
	if len(ts) > 0 {
		t = ts[0]
	}
	return t.Format(format)
}

func StrToTime(strTime string) (time.Time, error) {
	if strTime == "" {
		return time.Time{}, nil
	}

	zoneName, offset := time.Now().Zone()
	zoneValue := offset / 3600 * 100

	if zoneValue > 0 {
		strTime += fmt.Sprintf(" +%04d", zoneValue)
	} else {
		strTime += fmt.Sprintf(" -%04d", zoneValue)
	}

	if zoneName != "" {
		strTime += " " + zoneName
	}

	return StrToZoneTime(strTime)
}

func StrToZoneTime(strTime string) (t time.Time, err error) {
	if strTime == "" {
		return time.Time{}, err
	}

	layouts := []string{
		"2006-01-02 15:04:05 -0700 MST",
		"2006-01-02 15:04:05 -0700",
		"2006-01-02 15:04:05",
		"2006/01/02 15:04:05 -0700 MST",
		"2006/01/02 15:04:05 -0700",
		"2006/01/02 15:04:05",
		"2006-01-02 -0700 MST",
		"2006-01-02 -0700",
		"2006-01-02",
		"2006/01/02 -0700 MST",
		"2006/01/02 -0700",
		"2006/01/02",
		"2006-01-02 15:04:05 -0700 -0700",
		"2006/01/02 15:04:05 -0700 -0700",
		"2006-01-02 -0700 -0700",
		"2006/01/02 -0700 -0700",
		time.ANSIC,
		time.UnixDate,
		time.RubyDate,
		time.RFC822,
		time.RFC822Z,
		time.RFC850,
		time.RFC1123,
		time.RFC1123Z,
		time.RFC3339,
		time.RFC3339Nano,
		time.Kitchen,
		time.Stamp,
		time.StampMilli,
		time.StampMicro,
		time.StampNano,
	}

	for _, layout := range layouts {
		t, err = time.Parse(layout, strTime)
		if err == nil {
			return t, err
		}
	}

	return t, err
}
