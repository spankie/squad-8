package algo

import "testing"

func TestNumInList(t *testing.T) {
	s := []int{}
	num := 3
	isPresent := NumInList(s, num)
	if isPresent != false {
		t.Errorf("isPresent is true; expected false")
	}
	//s = []int{1,2,3,4,5,6}
	//num = 3
	//isPresent = NumInList(s, num)
	//if isPresent != true {
	//	t.Errorf("isPresent is false; expected true")
	//}
}