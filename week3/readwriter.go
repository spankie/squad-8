package main

import "fmt"

type ReadWriter interface {
	 Read() string
	 Write(out string)
}

type ChemistryBook struct{}

func (cli ChemistryBook) Read() string {
	return "covalent and electrovalent bonds"
}

func (cli ChemistryBook) Write(output string) {
	fmt.Println(output)
}

type BiologyBook struct{}

func (cli BiologyBook) Read() string {
	return "what is photosynthensis"
}

func (cli BiologyBook) Write(output string) {
	fmt.Println(output)
}

type CommandLine struct{}

func (cli CommandLine) Read() string {
	var input string

	fmt.Scanf("%s\n", &input)

	return input
}

func (cli CommandLine) Write(output string) {
	fmt.Printf("%v", output)
}

func Student(c ReadWriter) {
	val := c.Read() // i am a boy

	// write what was read
	c.Write("I just read: "+val)
}

func main() {
	Student(CommandLine{})
}