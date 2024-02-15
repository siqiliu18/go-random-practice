package doordash

import (
	"fmt"
	"strconv"
	"strings"
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

func TimeFormats(start, end string) string {
	beginTime, _ := time.Parse("2006-01 15:04", "2024-01 00:00")

	startArr := strings.Split(start, " ")
	endArr := strings.Split(end, " ")

	startTime := strings.Split(startArr[1], ":")
	endTime := strings.Split(endArr[1], ":")

	startHour := startTime[0]
	sh, _ := strconv.Atoi(startHour)
	startMin := startTime[1]
	sm, _ := strconv.Atoi(startMin)

	endHour := endTime[0]
	eh, _ := strconv.Atoi(endHour)
	endMin := endTime[1]
	em, _ := strconv.Atoi(endMin)

	startDay := startArr[0]
	endDay := endArr[0]

	startAP := startArr[2]
	endAP := endArr[2]

	startTimeDay := createTime(findDate(startDay), sh, sm, startAP)
	endTimeDay := createTime(findDate(endDay), eh, em, endAP)

	fmt.Println(startTimeDay)
	fmt.Println(endTimeDay)

	fmt.Println("----------------")

	diff := startTimeDay.Sub(endTimeDay)
	fmt.Println(diff < 0)
	diff = endTimeDay.Sub(startTimeDay)
	fmt.Println(diff > 0)

	fmt.Println("----------------")
	for currTime := startTimeDay; endTimeDay.Sub(currTime) >= 0; currTime = currTime.Add(time.Minute * time.Duration(5)) {
		// source - https://stackoverflow.com/a/33119937/9954367
		fmt.Println(dayIntMap[currTime.Weekday()] + currTime.Format("1504"))
	}

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

var dayIntMap = map[time.Weekday]string{
	time.Monday:    "1",
	time.Tuesday:   "2",
	time.Wednesday: "3",
	time.Thursday:  "4",
	time.Friday:    "5",
	time.Saturday:  "6",
	time.Sunday:    "7",
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

func createTime(day time.Time, hour, minute int, apm string) time.Time {
	if apm == "pm" {
		return day.Add(time.Hour*time.Duration(hour) + time.Hour*time.Duration(12) + time.Minute*time.Duration(minute))
	}
	// d := day.Add(time.Hour * time.Duration(12))
	// fmt.Println(d)
	return day.Add(time.Hour*time.Duration(hour) + time.Minute*time.Duration(minute))
}
