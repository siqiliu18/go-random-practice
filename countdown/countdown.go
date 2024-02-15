package countdown

import (
	"fmt"
	"time"
)

type ct struct {
	t, d, h, m, s int
}

// https://freshman.tech/golang-timer/
func GetTimeRemaining(t time.Time) ct {
	currTime := time.Now()
	diff := t.Sub(currTime)

	// fmt.Println(diff)

	total := int(diff.Seconds())

	fmt.Println(total)

	days := int(total / (60 * 60 * 24))
	hours := int(total / (60 * 60) % 24)
	minutes := int(total/60) % 60
	second := int(total % 60)

	return ct{t: total, d: days, h: hours, m: minutes, s: second}
}

func TickFunc(v time.Time) {
	for range time.Tick(1 * time.Second) {
		timeRemaining := GetTimeRemaining(v)

		if timeRemaining.t <= 0 {
			fmt.Println("Countdown reached!")
			break
		}

		fmt.Printf("Days: %d, Hours: %d, Minutes: %d, Seconds: %d\n", timeRemaining.d, timeRemaining.h, timeRemaining.m, timeRemaining.s)
	}
}

// https://pkg.go.dev/time#NewTicker
func NewTickFunc(v time.Time) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for range ticker.C {
		timeRemaining := GetTimeRemaining(v)

		if timeRemaining.t <= 0 {
			fmt.Println("Countdown reached!")
			break
		}

		fmt.Printf("Days: %d, Hours: %d, Minutes: %d, Seconds: %d\n", timeRemaining.d, timeRemaining.h, timeRemaining.m, timeRemaining.s)

	}
}
