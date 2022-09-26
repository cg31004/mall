package handler

import (
	"regexp"
	"time"

	"github.com/gin-gonic/gin"
)

func GetStartTime(ctx *gin.Context) (time.Time, bool) {
	timeStr, existed := ctx.GetQuery("startTime")
	if !existed {
		return time.Time{}, false
	}
	return parseTime(timeStr)
}

func GetEndTime(ctx *gin.Context) (time.Time, bool) {
	timeStr, existed := ctx.GetQuery("endTime")
	if !existed {
		return time.Time{}, false
	}
	return parseTime(timeStr)
}

func parseTime(timeStr string) (time.Time, bool) {
	t, err := time.Parse(time.RFC3339, timeStr)
	if err != nil {
		return time.Time{}, false
	}
	return t, true
}

var timezonePattern = regexp.MustCompile(`^UTC[+-]\d{2}:\d{2}`)

func ValidateTimezone(timezone string) bool {
	return timezonePattern.Match([]byte(timezone))
}
