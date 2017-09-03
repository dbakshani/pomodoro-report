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
		//fmt.Printf("parts: %q\n", parts)
		// only process lines that match the regex
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

	var processingDate time.Time

	sort.Strings(dates)

	if len(dates) > 0 {
		// find the earliest date from the provided file
		//processingDate = convertToTime(dates[0])
		processingDate.UnmarshalText([]byte(dates[0]))
	} else {
		fmt.Println("No pomodoro data found.")
		os.Exit(2)
	}

	//	for _, date := range dates {
	//		//y, _ := strconv.Atoi(date[:4])
	//		//m, _ := strconv.Atoi(date[4:6])
	//		//d, _ := strconv.Atoi(date[6:8])
	//		//l, _ := time.LoadLocation("UTC")
	//		//dt := time.Date(y, time.Month(m), d, 0, 0, 0, 0, l)
	//		dt := convertToTime(date)
	//		layout := "Mon 2006-01-02"
	//
	//		fmt.Printf("%v: %v\t%v\n", dt.Format(layout), pomodoros[date], nil)
	//	}

	for processingDate.Before(time.Now().UTC()) {
		//fmt.Printf("%v: %v\n", processingDate, pomodoroForDate(pomodoros, processingDate))
		fmt.Printf("%v: %v\n", processingDate, pomodoros[processingDate])

		processingDate = processingDate.Add(time.Hour * 24)
	}
}

// convertToTime return a Time object from the provided date string in yyyymmdd format.
func convertToTime(date string) time.Time {
	y, _ := strconv.Atoi(date[:4])
	m, _ := strconv.Atoi(date[4:6])
	d, _ := strconv.Atoi(date[6:8])
	l, _ := time.LoadLocation("UTC")
	return time.Date(y, time.Month(m), d, 0, 0, 0, 0, l)
}
