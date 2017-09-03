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

	pomodoros := make(map[string]int)
	var dates []string

	var r = regexp.MustCompile(`^(\d{8})@.*(p)$`)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		//fmt.Println(line)
		parts := r.FindStringSubmatch(line)
		//fmt.Printf("parts: %q\n", parts)
		if len(parts) != 0 {
			if val, ok := pomodoros[parts[1]]; ok {
				pomodoros[parts[1]] = val + 1
			} else {
				pomodoros[parts[1]] = 1
				dates = append(dates, parts[1])
			}
		}

	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	sort.Strings(dates)
	for _, date := range dates {
		y, _ := strconv.Atoi(date[:4])
		m, _ := strconv.Atoi(date[4:6])
		d, _ := strconv.Atoi(date[6:8])
		l, _ := time.LoadLocation("UTC")
		//fmt.Println("parts: ", y, m, d)
		dt := time.Date(y, time.Month(m), d, 0, 0, 0, 0, l)
		layout := "Mon 2006-01-02"

		fmt.Printf("%v: %v\n", dt.Format(layout), pomodoros[date])
	}
}
