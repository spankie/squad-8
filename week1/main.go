package main

// collection of packages - check
// go.mod at it's root

import (
	"fmt"

	ct "github.com/daviddengcn/go-colortext"
)

func main() {
	ct.Foreground(ct.Green, false)
	fmt.Println("Green text starts here...")
	ct.ChangeColor(ct.Red, true, ct.Green, false)
	fmt.Println("some other text")
	ct.ResetColor()
	text := fmt.Sprintf("testing gofmt")
	fmt.Println(text)
}
