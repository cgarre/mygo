/*
one - This is to use merge intervals
input intervals struct
*/
package main

import (
	"fmt"
)

/*
Interval - the interval struct
*/
type Interval struct {
	Start int
	End   int
}

/*
Intervals - Intervals slice
*/
type Intervals []Interval

func printMerge(interv Intervals) {
	currStart := interv[0].Start
	currEnd := interv[0].End

	for i := 1; i < len(interv); i++ {
		if interv[i].Start <= currEnd {
			currEnd = interv[i].End
		} else {
			fmt.Println(currStart, ",", currEnd)
			currStart = interv[i].Start
			currEnd = interv[i].End
		}
	}

	fmt.Println(currStart, ",", currEnd)
}

func main() {

	interv := Intervals{
		{2, 4},
		{3, 6},
		{8, 12},
		{14, 16},
		{16, 22},
		{21, 30}}

	printMerge(interv)

}
