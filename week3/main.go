package main

import "fmt"

type Name string

func (n *Name) SayMyName(s string) {
	fmt.Printf("My name is %v\n", *n)
	if s != "" {
		*n = Name(s)
	}
}

func Receivers() {
	var franklyn Name = "Franklyn"
	franklyn.SayMyName("new name")
	// second call
	franklyn.SayMyName("")
	franklyn.SayMyName("")
	// SayMyName()
}

func Arr() {
	var arr [2]int64
	arr[0] = 1

	arr2 := [...]int64{1, 2, 3,4,5,6,7,8,9,0,9,8,7,6}
	//arr2[14] = 4 // not going to work because index is out of range
	fmt.Printf("len of arr2: %v\n", len(arr2))
	// slices
	//slice := []string{"spankie", "franklyn", "shuaib"}
	//slice = append(slice, "Victor")
	//slice[3] = "Emmanuel"
	//fmt.Printf("len: %v; cap: %v\n", len(slice), cap(slice))
	//slice[0] = "David"
	//slice[3] = "Victor"
	//slice[5] = "Emmanuel"
	//fmt.Printf("slice : %v\n", slice)

	//slice2 := make([]int64, 4)
	//slice2[0] = 201
	//slice2 = append(slice2, 300)
	//fmt.Println(slice2)
	//fmt.Printf("len: %v; cap: %v\n", len(slice2), cap(slice2))

	var s [][]int = make([][]int, 0, 8)
	fmt.Println(s, len(s), cap(s))
	//if s == nil {
	//	fmt.Println("nil!")
	//}
	s = append(s, []int{1,2}, []int{3,4}, []int{5,6}, []int{7,8})
	for _, v := range s {
		fmt.Println(v)
	}
	// len of n or n-1 index
	s = s[:len(s)-1]
	fmt.Printf("new slice: %v\n", s)
	s = s[:len(s)-1]
	fmt.Printf("new slice: %v\n", s)

	nums := []int{3,4, 5,6,7,8,9}
	nums = remove(nums, 3)
	fmt.Printf("new nums: %v\n", nums)
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func maps() {
	//var m map[int]string = map[int]string{}
	m := make(map[int]string)
	m[0] = "Emmanuel"
	m[1] = "Victor"
	m[2] = "Omoyemi"
	fmt.Println(m)
	for key, val := range m {
		fmt.Printf("%d : %s\n", key, val)
	}
	//name := m[0]
	//fmt.Println(name)

	key := 0
	name, ok := m[key]
	if ok {
		// it means that there is a key of 10 in the map
		fmt.Printf("key of %d has value of: %s\n", key, name)
	} else {
		fmt.Println("m does not have key 10")
	}
	//var m = map[string]string{
	//	"Bell Labs": `Vertex{
	//		40.68433, -74.39967,
	//	}`,
	//	"Google": `Vertex{
	//		37.42202, -122.08408,
	//	}`,
	//}
}

// structs

type Address struct {
	Street string
}

type Person struct {
	Name string
	Age int
	Height, Weight float32
	Address []Address
}

func (p Person) String() string {
	return fmt.Sprintf("My name is %s, I am %d years old, and I weigh %.2f\n",
		p.Name, p.Age, p.Weight)
}

func (p *Person) LooseWeight(w float32) {
	p.Weight -= w
}

func OOP() {
	Nnamdi := Person{
		Name: "Nnamdi",
		Age: 54,
		Height: 20.56,
		Weight: 300.45,
	}

	Franklyn := Person{
		Name: "Franklyn",
		Age: 18,
		Height: 23.45,
		Weight: 45,
	}

	persons := []Person{
		{
			Name: "New person",
			Age: 34,
			Height: 23,
			Weight: 45,
		},
		Franklyn,
		Nnamdi,
	}
	for _, person := range persons[:1] {
		fmt.Println(person)
		person.LooseWeight(5)
		fmt.Println(person)
	}

}

func Change(v []int) {
	v[0] = 50
}

func main() {
	//Arr()
	//maps()
	//OOP()
	//a := []int{1,2,3,4,5,6,7,8,9}
	//Change(a)
	//fmt.Println(a)
	//b := a
	//// | | | | | | |
	//fmt.Printf("[1]: address of slice a %p add of b %p \n", &a, &b)
	//a[0] = 30
	//a[8] = 35
	//fmt.Printf("[2]: address of slice a %p add of b %p \n", &a, &b)
	//a = append(a, 12)
	//a[0] = 45
	//fmt.Printf("[3]: address of slice a %p add of b %p \n", &a, &b)
	//fmt.Printf("address of slice %v add of Arr %v \n", a, b)

	//ghost := Cat{
	//	Basics:       Animal{
	//		Name: "Ghost",
	//		mean: true,
	//	},
	//	MeowStrength: 2,
	//}
	//var shout NoiseMaker
	ice := &Dog{
		Animal:       Animal{
			Name: "Ice",
			mean: false,
		},
		BarkStrength: 3,
	}

	//shout = ice
	//ice.MakeNoise()
	//ghost.MakeNoise()

	// interfaces
	//n := NoiseMaker()
	MakeSomeNoise(ice)
}