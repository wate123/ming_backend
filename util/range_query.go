package util

import (
	"log"
	"strings"
)

var rangeQuery = map[string]string{
	"day":   "DAY(invdate)",
	"week":  "WEEK(invdate, 3)",
	"month": "MONTH(invdate)",
	"year":  "YEAR(invdate)"}

func RangeBy(timeRange string) string{
	if value, exist := rangeQuery[strings.ToLower(timeRange)]; exist {
		return value
	} else {
		log.Panic("range not found")
		return ""
	}
}
