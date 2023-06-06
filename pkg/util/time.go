package util

import (
	"time"

	"github.com/sirupsen/logrus"
)

var (
	Timer timer
)

const (
	DateFormat     = "2006-01-02"
	DateTimeFormat = "2006-01-02 15:04:05"
)

type timer struct{}

func (timer) Date2String(t time.Time) string {
	return t.Format(DateFormat)
}

func (timer) Time2String(t time.Time) string {
	return t.Format(DateTimeFormat)
}

func (timer) String2Time(s string) time.Time {
	t, err := time.ParseInLocation(DateTimeFormat, s, time.Local)
	if err != nil {
		logrus.WithField("String2Time", s).Warn(err)
	}
	return t
}
