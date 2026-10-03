package main

import (
	"bufio"
	"fmt"
	"os"
)

func techniciansLog() {
	separator := "----------------------"
	jobLog := []string{}
	jobMap := make(map[string]int)

	scanner := bufio.NewScanner(os.Stdin)
	for i := 1; ; i++ {
		fmt.Printf("Job %d Input > ", i)
		scanner.Scan()
		text := scanner.Text()
		if text == "done" {
			break
		}
		if text == "" {
			fmt.Println("No jobs entered")
			return
		}
		jobLog = append(jobLog, text)
		jobMap[text]++
	}
	if len(jobLog) == 0 {
		fmt.Println("No jobs entered")
		return
	}

	fmt.Println(separator)
	fmt.Println("Technicians Log:")

	for i, job := range jobLog {
		output := fmt.Sprintf("%d. %s", i+1, job)
		fmt.Println(output)
	}
	fmt.Println(" ")
	fmt.Printf("Job count: %d\n", len(jobLog))
	fmt.Println(separator)
	fmt.Println("Job stats:")
	mostFrequent := []string{}
	mostFrequentCount := 0
	for job, count := range jobMap {
		if count > mostFrequentCount {
			mostFrequentCount = count
			mostFrequent = nil
			mostFrequent = append(mostFrequent, job)
		} else if count == mostFrequentCount {
			mostFrequent = append(mostFrequent, job)
		}
		fmt.Printf("%s: %d\n", job, count)
	}
	fmt.Println(separator)
	fmt.Println("Most frequent:")
	for _, job := range mostFrequent {
		fmt.Printf("%s: %d\n", job, mostFrequentCount)
	}
}

func main() {
	techniciansLog()
}
