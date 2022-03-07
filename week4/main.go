package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"os"
	"time"

	"github.com/pkg/errors"
)

type Name string

type Err struct{ Name string }

func (e Err) Error() string {
	return "Error: " + e.Name
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

type User struct {
	ID          int    `json:"-"`
	Name        string `json:"name,omitempty"`
	Occupation  string `json:"occupation"`
	Description string `json:"description"`
}

/// {} curly brackets or braces
/// [] square brackets
// . < > angle brackets
// () parentheses

func Json() {
	/*
		jsonObject := `
		[
			{
				"key": "value",
				"key2": 10,
				"key3": 10.00,
				"key4": {
					"key5": "value",
					"key6": [
						"string", 10, {}
					]
				}
			}
		]
		`

		// convert json string to golang object
		var jsonData interface{}

		err := json.Unmarshal([]byte(jsonObject), &jsonData)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%+v\n", jsonData)
	*/
	// for _, m := range jsonData {
	// 	for k, v := range m {
	// 		fmt.Printf("%s : %v\n", k, v)
	// 	}
	// }

	// J S O N => Javascript Object Notation

	u1 := User{
		ID:          1,
		Description: "Sample description",
		// Name:        "John Doe",
		Occupation: "gardener",
	}
	// json_data, err := json.Marshal(u1)

	json_data, err := json.MarshalIndent(u1, "-", "	")

	if err != nil {

		log.Fatal(err)
	}

	fmt.Println(string(json_data))
	/*
		users := []User{
			{ID: 2, Name: "Roger Roe", Occupation: "driver"},
			{ID: 3, Name: "Lucy Smith", Occupation: "teacher"},
			{ID: 4, Name: "David Brown", Occupation: "programmer"},
		}

		json_data2, err := json.MarshalIndent(users, "", "  ")

		if err != nil {

			log.Fatal(err)
		}

		fmt.Println(string(json_data2))
	*/
}

func Json2() {
	obj := `{
	"id": 1,
  "name": "John Doe",
  "Occupation": "gardener",
"Description": "5"
}`
	/*
			obj1 := `
		[
		  {
		    "Id": 2,
		    "Name": "Roger Roe",
		    "Occupation": "driver"
		  },
		  {
		    "Id": 3,
		    "Name": "Lucy Smith",
		    "Occupation": "teacher"
		  },
		  {
		    "Id": 4,
		    "Name": "David Brown",
		    "Occupation": "programmer"
		  }
		]
		`
	*/
	var user1 *User = &User{}
	//user1 := map[string]interface{}{}
	err := json.Unmarshal([]byte(obj), &user1)
	if err != nil {
		// return error that says: bad request
		log.Fatal(err)
	}
	log.Printf("%+v\n", user1)
}

type Transaction struct {
	Description, Status  string
	CreatedAt, UpdatedAt time.Time
}

func Timestamp() {
	t := time.Now()
	//tr := Transaction{
	//	CreatedAt: t,
	//	UpdatedAt: t,
	//}
	//t.Year()
	timeUTC := t.UTC()
	// Nigeria time = 12:00noon  ---- 4pm
	// france time = 6:00pm
	// find my transaction between 12:00 WAT and 3:00 WAT
	fmt.Printf("timeWAT : %v\nTime UTC: %v\n", t, timeUTC)

	someonesBirthday := time.Date(1990, time.May, 10, 23, 12, 5, 3, time.UTC)

	day := 24 * time.Hour
	newBDay := someonesBirthday.Add(365 * 2 * day)
	t.Add(time.Second)
	fmt.Printf("someone's birthday: %v\nin two years time birthday = %v\n", someonesBirthday, newBDay)
	// Mon Jan 2 15:04:05 -0700 MST 2006
	fmt.Println(t.Format("Mon Jan 02 01 2006 MST"))
	// t.Sub(someonesBirthday) // returns duration
	t.Add(-time.Hour * 24)
}

//
//type Todo struct {
//	Name        string
//	Description string
//	ID          string
//}

type entry struct {
	Name string
	Done bool
}

type ToDo struct {
	User string
	List []entry
}

//var html = ``

func Templates() {
	//td := Todo{
	//	Name:        "Test templates",
	//	Description: "Let's test a template to see the magic.",
	//	ID: "1",
	//}
	/*
		t, err := template.New("todos").Parse("You have a task named \"{{ .ID}}. {{ .Name}}\" with description: \"{{ .Description}}\"")
		if err != nil {
			panic(err)
		}
		// You have a task named "Test templates"
		// You have a task named "1. Test templates"
		err = t.Execute(os.Stdout, td)
		if err != nil {
			panic(err)
		}
	*/
	todos := ToDo{
		User: "Omoyemi",
		List: []entry{
			{
				Name: "Write test",
				Done: false,
			},
			{
				Name: "Eat Lunch",
				Done: false,
			},
		},
	}
	/*
			{
				User: "Clinton",
				List: []entry{{
					Name: "See my girlfriend",
					Done: false,
				},
					{
						Name: "Go to dinner date",
						Done: true,
					},
				},
			},
		}
	*/

	tmpls := []string{"todo.tmpl"}
	t := template.Must(template.ParseFiles(tmpls...))

	err := t.Execute(os.Stdout, todos)
	if err != nil {
		panic(err)
	}
}

func main() {
	Json()
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
	//file, err := os.Open("./content/novel.txt") // For read access.
	//if err != nil {
	//	log.Fatal(err)
	//}
	//data := make([]byte, 100)
	//count, err := file.Read(data)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Printf("read %d bytes: %q\n", count, data[:count])
	//Json2()
	//Timestamp()
	// Templates()
}
