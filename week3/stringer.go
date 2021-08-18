package main

//Interfaces are named collections of method signatures.

type Stringer interface {
	String() string // declare
	End(s string)
}

type Anything func()

func (a Anything) String() string {
	return ""
}
func (a Anything) End(s string) {
// anything
	// one plus one
}

// anything that implements all the methods inside an interface can be used as that interface

func Caller(s Anything) {

}

func main() {
	var a Anything = func() {}
	Caller(a)
}