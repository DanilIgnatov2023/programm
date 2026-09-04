package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	var newYear time.Time
	if now.Month() == time.December && now.Day() > 31 {
		newYear = time.Date(now.Year()+2, time.January, 1, 0, 0, 0, 0, now.Location())
	} else if now.Month() == time.December && now.Day() == 31 {
		newYear = time.Date(now.Year()+1, time.January, 1, 0, 0, 0, 0, now.Location())
	} else {
		newYear = time.Date(now.Year()+1, time.January, 1, 0, 0, 0, 0, now.Location())
	}
	
	diff := newYear.Sub(now)
	
	days := int(diff.Hours() / 24)
	
	hours := int(diff.Hours()) % 24
	minutes := int(diff.Minutes()) % 60
	seconds := int(diff.Seconds()) % 60

	fmt.Printf("До Нового года осталось:\n")
	fmt.Printf("   %d дней, %d часов, %d минут, %d секунд\n", days, hours, minutes, seconds)
	fmt.Printf("   или %d полных дней\n", days)
	
}
