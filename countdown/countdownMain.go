package main

import (
	"flag"
	"fmt"
	"go-random-practice/countdown/countdown"
	"os"
	"time"
)

func main() {
	fmt.Println("---------------")
	// ex: go build main.go
	// ./main -deadline=2024-05-03T17:00:00+01:00

	deadline := flag.String("deadline", "", "The deadline for the countdown timer in RFC3339 format (e.g. 2019-12-25T15:00:00+01:00")
	flag.Parse()

	if *deadline == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	v, err := time.Parse(time.RFC3339, *deadline)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	countdown.NewTickFunc(v)
}
