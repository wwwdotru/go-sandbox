package arrays

import "fmt"

// ArrayInit is an example of how to initialize arrays
func ArrayInit() error {

	var val = [10]int{12, 353, 564, 6454}
	var val2 = [10]int{0: 12, 1: 353, 2: 564, 3: 6454, 6: 9999} // Note last index, you can skip positions
	var msg = [12]rune{0: 'h', 2: 'E'}
	var days = [7]string{
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday",
		"Saturday",
		"Sunday",
	}

	var days2 = [...]string{
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Sunday",
	}

	var truth = [12]bool{true}

	var histogram = [5]map[string]int{
		map[string]int{"A": 12, "b": 1},
		map[string]int{"cat": 12, "dog": 1, "rat": 848},
	}

	var board = [4][2]int{
		{33, 12},
		{343, 23},
		{343, 23},
		{343, 23},
	}

	fmt.Printf("%T: %v\n", val, val)
	fmt.Printf("%T: %v\n", val2, val2)
	fmt.Printf("%T: %v\n", msg, msg)
	fmt.Printf("%T: %v\n", days, days)
	fmt.Printf("%T: %v\n", days2, days2)
	fmt.Printf("%T: %v\n", truth, truth)
	fmt.Printf("%T: %v\n", histogram, histogram)
	fmt.Printf("%T: %v\n", board, board)

	return nil
}
