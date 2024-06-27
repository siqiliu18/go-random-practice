package main

import (
	"fmt"
	"time"
)

func main() {
	currTime := time.Now()
	fmt.Printf("current time: %v\n", currTime)

	futureTime, _ := time.Parse(time.RFC3339, "2024-06-28T17:00:00-07:00")
	fmt.Printf("future time: %v\n", futureTime)

	// Sub returns the duration t-u. If the result exceeds the maximum (or minimum) value that can be stored in a Duration, the maximum (or minimum) duration will be returned. To compute t-d for a duration d, use t.Add(-d).
	diff := futureTime.Sub(currTime)
	fmt.Printf("time difference: %v\n", diff)

	diffInSecs := int(diff.Seconds())
	fmt.Printf("time diff in seconds: %v\n", diffInSecs)

	// base values
	oneDay := 60 * 60 * 24
	oneHour := 60 * 60
	oneMinute := 60

	diffDay := diffInSecs / oneDay
	diffHour := (diffInSecs % oneDay) / oneHour
	diffMinute := ((diffInSecs % oneDay) % oneHour) / oneMinute
	diffSec := ((diffInSecs % oneDay) % oneHour) % oneMinute

	fmt.Printf("diff in Day %v, Hour %v, Minute %v, Second %v: ", diffDay, diffHour, diffMinute, diffSec)
}
