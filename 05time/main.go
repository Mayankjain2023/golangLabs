package main

import (
	"fmt"
	"time"
)

const date_format string = "01-02-2006 15:04:05 Monday"

func main(){
	fmt.Println("Getting started with learning time ...")
	currentTime := time.Now()

	formatTime := currentTime.Format(date_format) //it always have to be like this -> shittyy

	fmt.Println("Current time: ",currentTime);
	fmt.Println("Formatted time :",formatTime)

	randomDate := time.Date(2020,time.December,20,20,23,15,0,time.UTC)
	fmt.Println("Random time new :",randomDate)
	fmt.Println("Format random",randomDate.Format(date_format));



}


// Simple starter examples.
// do("Basic full date", "Mon Jan 2 15:04:05 MST 2006", "Wed Feb 25 11:06:39 PST 2015")
// do("Basic short date", "2006/01/02", "2015/02/25")

// The hour of the reference time is 15, or 3PM. The layout can express
// it either way, and since our value is the morning we should see it as
// an AM time. We show both in one format string. Lower case too.
// do("AM/PM", "3PM==3pm==15h", "11AM==11am==11h")

// When parsing, if the seconds value is followed by a decimal point
// and some digits, that is taken as a fraction of a second even if
// the layout string does not represent the fractional second.
// Here we add a fractional second to our time value used above.
// t, err = time.Parse(time.UnixDate, "Wed Feb 25 11:06:39.1234 PST 2015")
// if err != nil {
// 	panic(err)
// }
// It does not appear in the output if the layout string does not contain
// a representation of the fractional second.
// do("No fraction", time.UnixDate, "Wed Feb 25 11:06:39 PST 2015")

// Fractional seconds can be printed by adding a run of 0s or 9s after
// a decimal point in the seconds value in the layout string.
// If the layout digits are 0s, the fractional second is of the specified
// width. Note that the output has a trailing zero.
// do("0s for fraction", "15:04:05.00000", "11:06:39.12340")

// If the fraction in the layout is 9s, trailing zeros are dropped.
// do("9s for fraction", "15:04:05.99999999", "11:06:39.1234")
