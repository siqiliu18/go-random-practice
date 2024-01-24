package doordash

import (
	"fmt"
	"time"
)

// To find all the stores that are open in a user’s delivery radius, we need to check the store’s operating hours.
// We currently store this information in Elasticsearch but the query performance for our use-case is turning out to be
// not very efficient. The format that we store it in Elasticsearch is something like this:

// {
//   "store": {
//     "store_name": "Tasty Pizzas",
//     "store_id": "123345",
//     "operating_hours": [
//       {
//         "hours.open": "mon 10:00 am",
//         "hours.close": "mon 2:00 pm"
//       },
//       {
//         "hours.open": "tue 10:00 am",
//         "hours.close": "tue 2:00 pm"
//       },
//       {
//         "hours.open": "wed 10:00 am",
//         "hours.close": "wed 2:00 pm"
//       },
//       {
//         "hours.open": "thu 10:00 am",
//         "hours.close": "thu 2:00 pm"
//       },
//       {
//         "hours.open": "fri 10:00 am",
//         "hours.close": "fri 2:00 pm"
//       },
//       {
//         "hours.open": "sat 10:00 am",
//         "hours.close": "sat 2:00 pm"
//       }
//     ]
//   }
// }
// We want to experiment with improving the performance of fetching open stores by converting the operating hours into encoded tokens.
// These would be of fixed length of five.

// The first digit would represent the day, Monday maps to number 1, Tuesday to 2 and so on.
// The next four digits would represent the hours in 24 hours format.

// Examples:

// Monday, 10:00am transforms to 11000
// Monday, 2:00pm transforms to 11400 (as 2pm -> 14:00 in 24 hr format)

// Write a function that takes in a start_time and end_time and returns a list of all encoded tokens of in that range at 5 minute intervals.
// Ensure you validate the input.

// Ex: Input: ("mon 10:00 am", "mon 11:00 am")
// Output: [“11000”, “11005”, “11010”, “11015”, “11020”, “11025”, “11030”, “11035”, “11040”, “11045”, “11050”, “11055”, “11100”]

// Input: ("mon 10:15 am", "mon 11:00 am")
// Output: [“11015”, “11020”, “11025”, “11030”, “11035”, “11040”, “11045”, “11050”, “11055”, “11100”]

// Input: ("mon 00:00 am", "mon 00:00 am")
// Output: ["10000"]

// Input: ("mon 00:00 am", "mon 11:59 am")
// Output: ["10000", ...., 11155]

// Input: ("mon 00:00 am", "mon 2:00 pm")
// Output: ["10000", ...., 11400]

// Input: ("mon 00:00 am", "wed 1:59 pm")
// Output: ["10000", ...., "31355"]

// Input: ("mon 00:00 am", "wed 00:00 am")
// Output: ["10000", ...., "22355", "30000"]

/* 1. input validation, "mom 00:00 pm"
    - pm && 00:00

Note:   check last 2 digits = 55
        check 2nd and 3rd digit = 23 -> increase the 1st digit

*/

func TimeFormats() string {
	beginTime, _ := time.Parse("2006-01 15:04", "2024-01 00:00")
	// fmt.Println(beginTime)
	// fmt.Println(time.Friday.String())

	// Time := time.Date(2020, 11, 14, 10, 45, 16, 0, time.UTC)

	// local, _ := time.LoadLocation("")
	// now := time.Now().In(local)
	// beginOfToday := now.Truncate(24 * time.Hour)
	// fmt.Println(beginOfToday)

	// sunOffset := beginOfToday.Weekday() - time.Sunday
	// fmt.Println(sunOffset)

	// sunDate := beginOfToday.AddDate(0, 0, -int(sunOffset))
	// fmt.Println(sunDate)

	// friOffSet := beginOfToday.Weekday() - time.Friday
	// fmt.Println(friOffSet)

	// friDate := beginOfToday.AddDate(0, 0, -int(friOffSet))
	// fmt.Println(friDate)

	targetDay := findDate("fri")
	targetDayTime := addTime(targetDay, true, 3)
	fmt.Println(targetDayTime)

	return beginTime.GoString()
}

var dayMap = map[string]time.Weekday{
	"mon": time.Monday,
	"tue": time.Tuesday,
	"wed": time.Wednesday,
	"thu": time.Thursday,
	"fri": time.Friday,
	"sat": time.Saturday,
	"sun": time.Sunday,
}

// findDate returns the beginning of the input day of current week
func findDate(day string) time.Time {
	// source: https://stackoverflow.com/questions/25254443/return-local-beginning-of-day-time-object
	local, _ := time.LoadLocation("")
	now := time.Now().In(local)
	beginOfToday := now.Truncate(24 * time.Hour)

	// source: https://stackoverflow.com/questions/67275624/get-the-most-recent-monday-0530-unix-time-in-go
	dayOffset := beginOfToday.Weekday() - dayMap[day]
	targetDay := beginOfToday.AddDate(0, 0, -int(dayOffset))
	return targetDay
}

func addTime(day time.Time, hms bool, gap int64) time.Time {
	if hms {
		return day.Add(time.Hour * time.Duration(gap))
	}
	return day.Add(time.Minute * time.Duration(gap))
}
