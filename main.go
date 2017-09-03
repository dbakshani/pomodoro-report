package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"time"
)

/**
* This tool takes a file containing pomodoro data in the following format, and generates a
* report to stdout containing the date, daily total and weekly total of pomodoros.
*
* The expected format for the pomodoro data is:
*
*		20170901@1020 : p
*		20170901@1024 : s
*		20170901@1049 : p
*		20170901@1054 : s
*		20170901@1121 : p
*		20170901@1147 : p
*		20170901@1153 : s
*		20170901@1508 : p
*		20170901@1535 : p
*
* The times are ignored, as are lines that don't end with "p" (in my usage, "s" represents
* a short break, and "b" represents a long break).
*
**/

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("The only parameter needed is the name of the pomodoros file.")
		os.Exit(1)
	}
	file, err := os.Open(args[1])
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	defer file.Close()

	// mapping of date to pomodoro count
	pomodoros := make(map[time.Time]int)

	// array of dates for sorting results
	var dates []string

	var r = regexp.MustCompile(`^(\d{8})@.*(p)$`)

	// process file line-by-line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		//fmt.Println(line)
		parts := r.FindStringSubmatch(line)
		// only process lines that match the regex; i.e. have pomodoro data
		if len(parts) != 0 {
			t := convertToTime(parts[1])
			if val, ok := pomodoros[t]; ok {
				pomodoros[t] = val + 1
			} else {
				pomodoros[t] = 1
				dates = append(dates, t.Format(time.RFC3339))
			}
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	sort.Strings(dates)

	var processingDate time.Time
	if len(dates) > 0 {
		// find the earliest date from the provided file
		processingDate.UnmarshalText([]byte(dates[0]))
	} else {
		fmt.Println("No pomodoro data found.")
		os.Exit(2)
	}

	layout := "Mon 2006-01-02"
	var nextDay time.Time
	var weeklyTotal int

	// print report for each day from earliest date until today
	for processingDate.Before(time.Now().UTC()) {
		nextDay = processingDate.Add(time.Hour * 24)

		weeklyTotal = weeklyTotal + pomodoros[processingDate]

		// last day of current week or last day of data; print weekly total
		if processingDate.Weekday() > nextDay.Weekday() ||
			nextDay.After(time.Now().UTC()) {
			fmt.Printf("%v: %v\t%v\n", processingDate.Format(layout), pomodoros[processingDate], weeklyTotal)
			weeklyTotal = 0
		} else {
			fmt.Printf("%v: %v\n", processingDate.Format(layout), pomodoros[processingDate])
		}
		processingDate = nextDay
	}
}

// convertToTime returns a Time object from the provided date string in "yyyymmdd" format.
func convertToTime(date string) time.Time {
	y, _ := strconv.Atoi(date[:4])
	m, _ := strconv.Atoi(date[4:6])
	d, _ := strconv.Atoi(date[6:8])
	l, _ := time.LoadLocation("UTC")
	return time.Date(y, time.Month(m), d, 0, 0, 0, 0, l)
}
