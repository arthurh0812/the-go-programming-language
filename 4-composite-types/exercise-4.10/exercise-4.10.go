package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"time"

	github "github.com/arthurh0812/the-go-programming-language/4-composite-types/github"
)

const (
	OneDay = iota
	TwoDays
	ThreeDays
	FourDays
	FiveDays
	OneWeek
	OneMonth
	ThreeMonths
	OneYear
	ThreeYears
)

var durations = map[int]time.Duration{
	OneDay:      1 * 24 * time.Hour,
	TwoDays:     2 * 24 * time.Hour,
	ThreeDays:   3 * 24 * time.Hour,
	FourDays:    4 * 24 * time.Hour,
	FiveDays:    5 * 24 * time.Hour,
	OneWeek:     7 * 24 * time.Hour,
	OneMonth:    1 * 30 * 24 * time.Hour,
	ThreeMonths: 3 * 30 * 24 * time.Hour,
	OneYear:     365 * 24 * time.Hour,
	ThreeYears:  3 * 365 * 24 * time.Hour,
}

type sortByAge struct {
	slice []*github.Issue
}

func (s sortByAge) Len() int {
	return len(s.slice)
}

func (s sortByAge) Less(i, j int) bool {
	return s.slice[i].CreatedAt.After(s.slice[j].CreatedAt)
}

func (s sortByAge) Swap(i, j int) {
	s.slice[i], s.slice[j] = s.slice[j], s.slice[i]
}

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d issues:\n", result.TotalCount)

	// create a copy
	resultsByAge := make([]*github.Issue, len(result.Items))
	copy(resultsByAge, result.Items)

	sort.Sort(sortByAge{
		slice: resultsByAge,
	})

	printByAge(resultsByAge)
}

func printByAge(items []*github.Issue) {
	for _, item := range items {
		since := time.Since(item.CreatedAt)

		printDuration(since)

		fmt.Printf("\t#%-5d\t%9.9s\t%.55s\n", item.Number, item.User.Login, item.Title)
	}
}

func printDuration(since time.Duration) {
	switch {
	case since > durations[ThreeYears]:
		fmt.Printf("More than 3 years ago:")
	case since > durations[OneYear]:
		fmt.Printf("More than 1 year ago:")
	case since > durations[ThreeMonths]:
		fmt.Printf("More than 3 months ago:")
	case since > durations[OneMonth]:
		fmt.Printf("More than 1 month ago:")
	case since > durations[OneWeek]:
		fmt.Printf("More than 1 week ago:")
	case since > durations[FiveDays]:
		fmt.Printf("More than 5 days ago:")
	case since > durations[FourDays]:
		fmt.Printf("More than 4 days ago:")
	case since > durations[ThreeDays]:
		fmt.Printf("More than 3 days ago:")
	case since > durations[TwoDays]:
		fmt.Printf("More than 2 days ago:")
	case since > durations[OneDay]:
		fmt.Printf("More than 1 day ago:")
	case since < durations[OneDay]:
		fmt.Printf("Last 24 hours:")
	}
}
