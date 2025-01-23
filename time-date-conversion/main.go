package main

import (
	"fmt"
	"time"
)

// format -- yyyy-mm-dd--hh-mm-ss  (fixed format for date and time ans time is in 24 hr format)

// the 2006-01-02 , 15 :04 :05 --> this is the date in which the golang is annoced and if we want our type of format then we had to enter this format only

func main() {
	fmt.Println("Time and Date Conversion")

	curr_time := time.Now()
	fmt.Println(curr_time)

	// val := time.Now().Hour()
	// val := time.Now().Month()
	// val := time.Now().Year()
	// val := time.Now().Minute()
	// val := time.Now().Nanosecond()
	// val := time.Now().Day()
	// val := time.Now().Weekday()
	// val,val2,val3 := time.Now().Clock()
	// val , val2 , val3 := time.Now().Date()
	// val , val2 := time.Now().Zone()
	// val  := time.Kitchen


	// fmt.Println(curr_time.YearDay() )


	formatted := curr_time.Format("2006-01-02  , 15 :04 :05 , Monday") 
	// if we wnat the data the we had to only enter the this format only which is the official announce date of the golnag through tihs we can acces any values and in place of monady it will print the current day

	// formatted := curr_time.Format("2006-01-02  , 3 :04 :05")  // this is fo the 12 hour format values  ans the upper one shows the 24 hour format

	fmt.Println(formatted)
}