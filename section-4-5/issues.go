package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"time"

	"gopl.io/ch4/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		createdAt := item.CreatedAt
		currentTime := time.Now()
		timeDiff := createdAt.Sub(currentTime)
		diffInDays := math.Abs(timeDiff.Hours() / 24)
		var timeMessage string
		if diffInDays <= 30 {
			timeMessage = fmt.Sprintln("Less than a month")
		} else if diffInDays <= 364 {
			timeMessage = fmt.Sprintln("Less than a Year")
		}
		if diffInDays >= 365 {
			timeMessage = fmt.Sprintln("More than a year")
		}
		fmt.Printf("#%-5d %9.9s %.55s %s\n",
			item.Number, item.User.Login, item.Title, timeMessage)
	}
}
