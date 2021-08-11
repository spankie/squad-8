package main

import (
	"errors"
	"fmt"
	"strings"
)

func Speak(s ...Name) {
	for _, v := range s {
		fmt.Printf("%s - ", v)
	}
}

type Age int // custom type or type definition

type Name = string // type alias

func Day1() {
	var name string // declaring a variable

	name = "Spankie" // assigning to a variable

	fmt.Printf("my name is : %s\n", name)

	bob := Age(24) // assigning a value to a variable

	// s := "Golang" // declaring and assigning a string
	f := 5
	f -= 2
	fmt.Printf("type of bob is %T", bob)
	s := []string{"Hello", "World"}
	Speak(s...)

	var n Name = "jumping"
	Speak(string(n))

	fmt.Println("\n----------")
	// control structures
	// for loops
	names := []string{"Spankie", "Bob", "Alice", "Nnamdi", "Emmanuel"}
	for i := len(names) - 1; i >= 0; i -= 2 {
		fmt.Println(names[i])
	}
	// switch


	switch names[0] {
	case "name":
		fmt.Println("n is name")
		Speak("name")
	case "spankie":
		fmt.Println("n is spankie")
	default:
		fmt.Println("n is neither name nor spankie")
	}

	var t interface{} = 64
	switch t.(type) {
	case string:
		fmt.Println("t is of type string")
	case int:
		fmt.Println("t is of type int")
	case int64:
		fmt.Println("t is of type int64")
	default:
		fmt.Println("t is neither int nor string")
	}
}

func Day2() {
	// pointers and values
	/*
	val := `this is a very very long text that should take maybe up to 10KB a quick brown fox jumped over a dog or something" // 5 bytes
	this is a very very long text that should take maybe up to 10KB a quick brown fox jumped over a dog or something" // 5 bytes
	this is a very very long text that should take maybe up to 10KB a quick brown fox jumped over a dog or something" // 5 bytes
	this is a very very long text that should take maybe up to 10KB a quick brown fox jumped over a dog or something" // 5 bytes
	this is a very very long text that should take maybe up to 10KB a quick brown fox jumped over a dog or something" // 5 bytes
	this is a very very long text that should take maybe up to 10KB a quick brown fox jumped over a dog or something" // 5 bytes
	this is a very very long text that should take maybe up to 10KB a quick brown fox jumped over a dog or something" // 5 bytes
	this is a very very long text that should take maybe up to 10KB a quick brown fox jumped over a dog or something" // 5 bytes
	this is a very very long text that should take maybe up to 10KB a quick brown fox jumped over a dog or something" // 5 bytes
	this is a very very long text that should take maybe up to 10KB a quick brown fox jumped over a dog or something" // 5 bytes`
	*/
	//fmt.Println(&val)
	//val = "Valuing"
	//key := 64

	//var valPointer *string = &val
	//var keyPointer *int = &key
	//fmt.Printf("%v : %v : %v", valPointer, &val, &valPointer)
	//fmt.Printf("size of valPointer : %v\n", unsafe.Sizeof(valPointer))
	//fmt.Printf("size of val : %v\n", unsafe.Sizeof(val))
	//var k = &valPointer
	//fmt.Printf("&val : %v\n &valPointer : %v", &val, *&valPointer)
	// if else
	// for range TODO: edge case when using for range
	houses := [...]string{"Bungalow", "SemiDetached", "Detached", "Duplex", "terrace"}
	//|1|2|3|4|5|
	f := []*string{}
	for i, _ := range houses {
		//fmt.Println(&v)
		f = append(f, &houses[i])
		//f = append(f, &v)
		//fmt.Println(&houses[i])
	}
	for _, v := range f {
		fmt.Println(*v)
	}
}

func Login(username, password string) (string, error) {
	//fmt.Println("you are logged in")
	if username == "spankie" && password == "password" {
		return "You are logged in", nil
	}
	// send an email to user that you just logged in
	return "", errors.New("invalid username or password")
}

func SaveUser(user string) error {
	// save to database
	return nil
}

func Signup(name, email string) (func(), error) {
	if !strings.Contains(email, "@") {
		return  nil, fmt.Errorf("invalid email address")
	}
	age := 54
	err := SaveUser(name)
	if err != nil {
		// do something
	}
	return func() {
		fmt.Printf("%s is now registered, you can use %s as your username to login, you are %v years old\n", name, email, age)
	}, nil
}

func GetUser(name string, age int, callback func(name string)) {
	// plenty stuff
	// does some more things...
	if age > 20 {
		callback(name)
	} else {
		fmt.Println("you are not of the required age")
	}
}

func PlayMusic() (favInstrument string) {
	inst := [...]string{"guitar", "piano", "drums"}
	defer func(){
		favInstrument = "nothing"
		fmt.Printf("I have changed the return value: %s\n", favInstrument)
	}()
	for i, v := range inst {
		if i == 30 {
			fmt.Println("I like ", v)
			favInstrument = v
			return
		}
	}
	// fmt.Println("could not see any musical instrument i liked")
	return "violin"
}

func Roster(squad8 []string) {
	//defer func() {
	//	err := recover()
	//	fmt.Printf("recover: this is what happened: %v\n", err)
	//}()
	for i := 0; i <= len(squad8); i++ {
		fmt.Printf("%d: %s\n", i, squad8[i])
	}
}

func Add(a, b int) {
	fmt.Println(a+b)
}

func Adder(y int, callback func(int, int)) {
	callback(y, 1)
}

// main

func F() func(int, int) {
	return func(a, b int) {
		fmt.Println("definition of a function")
	}
}

/*     function declaration        */
func factorial(n int) uint64 {
	var factVal uint64 = 1 // uint64 is the set of all unsigned 64-bit integers.
	// Range: 0 through 18446744073709551615.
	if(n < 0){
		fmt.Print("Factorial of negative number doesn't exist.")
	}else{
		for i:=1; i<=n; i++ {
			factVal *= uint64(i)  // mismatched types int64 and int
		}
	}
	return factVal  /* return from function*/
}

// Fac Recursive function
func Fac(n int) int {
	if n == 0 {
		return 1
	}

	return n * Fac(n-1)
}

//  Fac(1) // return n * 1 = 1
// Fac(2) // 2 * 1
// Fac(3) // 3 * 2
// Fac(4) // 4 * 6
// main

func main() {
	// Adder
	fmt.Println(Fac(4))
	// defer
	//f := PlayMusic()
	//fmt.Println(f)

	// panic and recover
	//defer func() {
	//	err := recover()
	//	fmt.Println("2 recover: ", err)
	//}()
	//fmt.Println("we are talking about panic")
	////panic("err trying to run the code")
	//Roster([]string{"Clinton", "Shuaib", "Nnamdi", "Chinonso", "Emmanuel", "Omoyemi", "Chukwuebuka"})
	//fmt.Println("this should not be called if we panic")

	// multiple parameters and return values
	//s, err := Login("nnamdi", "password")
	//if err != nil {
	//	fmt.Printf("error: %v\n", err)
	//}
	//fmt.Println(s)

	// returning function
	//f, err := Signup("Clinton", "clinton@gmail.com")
	//if err != nil {
	//	fmt.Printf("error: %v\n", err)
	//}
	//f(64)

	// functions callbacks
	//auth := func(name string) {
	//	// do authentication
	//	fmt.Printf("Hi %s, you are eligible to travel\n", name)
	//}
	//GetUser("Adedeji", 23, auth)

	// closures
	//age := 30
	//fmt.Printf("first age: %v\n", age)
	//f, err := Signup("Chinonso", "chinonso@gmail.com")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//f()
}
