package model

import "time"

var (
	UserBirthdayFormat string = "2006-01-02"
	UserDefaultBirthday time.Time = time.Date(1900, 01, 01, 0, 0, 0, 0, time.UTC)
	UserDefaultGender string = "unknown"
)
