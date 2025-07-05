package formatting

import (
	"time"
)

func Format2TimestampStr(epochTime int64) (resp string) {
	return time.Unix(epochTime, 0).In(LoadTimezone("Asia/Jakarta")).Format(time.RFC3339)
}

func LoadTimezone(timeZone string) *time.Location {
	tz, err := time.LoadLocation(timeZone)
	if err != nil {
		tz = time.Local
	}

	return tz
}
