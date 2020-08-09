package xjutils

import (
	"encoding/binary"
	"strconv"
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
	now := time.Unix(int64(binary.BigEndian.Uint32(data)/1000), 0)
	*t = EsTime(now)
	return
}

func (t EsTime) MarshalJSON() ([]byte, error) {
	if time.Time(t).IsZero() {
		return []byte("0"), nil
	} else {
		i := time.Time(t).Unix() * 1000
		str := strconv.FormatInt(i, 10)
		return []byte(str), nil
	}
}
func Int64ToBytes(i int64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(i))
	return buf
}

func BytesToInt64(buf []byte) int64 {
	return int64(binary.BigEndian.Uint64(buf))
}

//func (t EsTime) MarshalJSON() ([]byte, error) {
//var buf = make([]byte, len(timeFormart)+2)
//if time.Time(t).IsZero(){
//	buf = []byte(`"1970-01-01 00:00:00"`)
//}else{
//	i := time.Time(t).Unix()*1000
//	binary.BigEndian.PutUint64(buf, uint64(i))
//}
//return buf, nil
//
//	b := make([]byte, 0, len(timeFormart)+2)
//	b = append(b, '"')
//	b = time.Time(t).AppendFormat(b, timeFormart)
//	b = append(b, '"')
//	if string(b)==`"0001-01-01 00:00:00"`{
//		b = []byte(`"1970-01-01 00:00:00"`)
//	}
//	return b, nil
//}

type EsDate time.Time

func (t *EsDate) UnmarshalJSON(data []byte) (err error) {
	now := time.Unix(int64(binary.BigEndian.Uint32(data)/1000), 0)
	*t = EsDate(now)
	return
}

func (t EsDate) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(dateFormart)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, dateFormart)
	b = append(b, '"')
	if string(b) == `"0001-01-01"` {
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
