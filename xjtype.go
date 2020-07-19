package xjutils

import (
	"encoding/binary"
	"time"
)

type Time time.Time

const (
	timeFormart = "2006-01-02 15:04:05"
	dateFormart = "2006-01-02"
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

type EsTime time.Time
func (t *EsTime) UnmarshalJSON(data []byte) (err error) {
	now := time.Unix(int64(binary.BigEndian.Uint32(data)/1000),0)
	*t = EsTime(now)
	return
}


func (t EsTime) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(timeFormart)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, timeFormart)
	b = append(b, '"')
	if string(b)==`"0001-01-01 00:00:00"`{
		b = []byte(`""`)
	}
	return b, nil
}

type EsDate time.Time
func (t *EsDate) UnmarshalJSON(data []byte) (err error) {
	now := time.Unix(int64(binary.BigEndian.Uint32(data)/1000),0)
	*t = EsDate(now)
	return
}


func (t EsDate) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(dateFormart)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, dateFormart)
	b = append(b, '"')
	if string(b)==`"0001-01-01"`{
		b = []byte(`""`)
	}
	return b, nil
}

type Date time.Time
func (t *Date) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+dateFormart+`"`, string(data), time.Local)
	*t = Date(now)
	return
}

func (t Date) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(dateFormart)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, dateFormart)
	b = append(b, '"')
	return b, nil
}

func (t Date) String() string {
	return time.Time(t).Format(dateFormart)
}

func (t Date) StringDiy(sFormat string) string {
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
