package main

import (
	"fmt"
	"github.com/pkg/errors"
	"log"
	"os"
)

type Name string

type Err struct{Name string}

func (e Err) Error() string {
	return "Error: "+ e.Name
}

func Run() (string, error) {
	f := Err{"cannot read file"}

	return f.Error(), errors.Wrap(f, "error from run")
}

func CallRun() {
	//_, err := Run()
	//if errors.Is(err, io.EOF) {
	//if t, ok := err.(io.); ok {
	//	fmt.Println(err)
	// fmt.Println(errors.Unwrap(err))
	// panic
	// return error
	// handle the error (create the file)
	// log the error
	//}
}

func main() {
	//out, err := os.Create("novel.txt") //give any name
	//if err != nil {
	//	fmt.Printf("error occured creating novel.txt: %v\n", err)
	//}
	//out.Write([]byte("here is a file content\twe want to put into novel.txt"))
	//defer out.Close()
	//
	//in, err := os.Create("novel1.txt") //give any name
	//if err != nil {
	//	fmt.Printf("error occured creating novel.txt: %v\n", err)
	//}
	//in.Write([]byte("This is the file we want to read from"))
	//defer in.Close()
	//
	//n, err := io.Copy(out, in)
	//if err != nil {
	//	log.Println(err)
	//}
	//log.Printf("we copied %d bytes from in to out\n", n)
	file, err := os.Open("./content/novel.txt") // For read access.
	if err != nil {
		log.Fatal(err)
	}
	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("read %d bytes: %q\n", count, data[:count])

}

