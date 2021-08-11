package algo

// NumInList write a function that takes in a slice of int and a number
// the function should return true if the num is in the slice
// else return false if the number is not in the slice
// NumInList([]int{123}, 2) returns true
// NumInList([]int{123}, 5) returns false
// NumInList([]int{123}, 0) returns false
// NumInList([]int{}, 2) returns false
func NumInList(list []int, num int) bool {
	Nummap := make(map[int]struct{})

	for _, v := range list {
		if _, ok := Nummap[v]; !ok {
			Nummap[v] = struct{}{}
		}
	}
	_, ok := Nummap[num]

	return ok
}