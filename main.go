package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
)

func main() {
	file, err := os.Open("pomodoros.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	pomodoros := make(map[string]int)
	var dates []string

	var r = regexp.MustCompile(`^(\d{8})@.*(p)$`)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		parts := r.FindStringSubmatch(line)
		fmt.Printf("parts: %q\n", parts)
		if len(parts) != 0 {
			if val, ok := pomodoros[parts[1]]; ok {
				pomodoros[parts[1]] = val + 1
			} else {
				pomodoros[parts[1]] = 1
				dates = append(dates, parts[1])
			}
		}

	}
	sort.Strings(dates)
	//for date, count := range pomodoros {
	for _, date := range dates {
		fmt.Printf("%v: %v\n", date, pomodoros[date])
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
