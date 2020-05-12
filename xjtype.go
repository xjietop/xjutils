package xjutils

import (
	"time"
)

type Time time.Time

const (
	timeFormart = "2006-01-02 15:04:05"
)

func (t *Time) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+timeFormart+`"`, string(data), time.Local)
	*t = Time(now)
	return
}

func (t Time) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(timeFormart)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, timeFormart)
	b = append(b, '"')
	return b, nil
}

func (t Time) String() string {
	return time.Time(t).Format(timeFormart)
}

func (t Time) StringDiy(sFormat string) string {
	return time.Time(t).Format(sFormat)
}

type Json string

func (j *Json) UnmarshalJSON(data []byte) (err error) {
	*j = Json(data)
	return
}

func (j Json) MarshalJSON() ([]byte, error) {
	b := []byte(j)
	return b, nil
}