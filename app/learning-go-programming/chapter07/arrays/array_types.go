package arrays

import "fmt"

//PrintDays is an example of arrays in go
func PrintDays() error {
	var val [10]int
	var days [7]string
	var weekdays [5]string
	var truth [10]bool
	var histogram [5]map[string]int
	var histogram2 [5]map[string][]int

	fmt.Printf("%T: %v\n", val, val)
	fmt.Printf("%T: %v\n", days, days)
	fmt.Printf("%T: %v\n", weekdays, weekdays)
	fmt.Printf("%T: %v\n", truth, truth)
	fmt.Printf("%T: %v\n", histogram, histogram)
	fmt.Printf("%T: %v\n", histogram2, histogram2)

	for i, val := range days {
		fmt.Printf("Day %d = %s\n", i, val)
	}
	return nil
}
