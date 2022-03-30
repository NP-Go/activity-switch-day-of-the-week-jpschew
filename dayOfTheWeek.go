package main

import (
	"fmt"
	"time"
)

func main() {
	//Insert code here

	//Now().Weekday() give mon, tue ... sun
	//Now().Day() give the date of the month
	day := time.Now().Weekday()
	fmt.Printf("%v, %T\n", day, day)
	// fmt.Printf("%v, %T\n", time.Wednesday, time.Wednesday)

	switch day {
	case time.Monday, time.Wednesday, time.Friday, time.Sunday:
		fmt.Println("The day is odd")
	case time.Tuesday, time.Thursday, time.Saturday:
		fmt.Println("The day is even")
	default:
		fmt.Print("No matching day")
	}
}
