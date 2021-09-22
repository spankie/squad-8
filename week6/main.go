package main

import (
	"fmt"
	"net/http"
	"os"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// handler root request
	/*response := `
		<h1>Hello there. Welcome to my awesome blog</h1>
		<table>
	<tr><th>Squad</th><th>Year</th></tr>
<tr><td>Eight</td><td>2021</td></tr>
</table>
	`
	 */
	responseJson := `
	{
	"name": "Emmanuel",
"squad" : "008",
"location": "Edo tech Park",
"Stack": "GOlang"
}
`
	header := w.Header()
	header.Add("Content-Type", "application/json")
	//header.Add()
	w.WriteHeader(http.StatusInternalServerError)
	// text/html
	// text/plain
	//
	w.Write([]byte(responseJson))

	// 2XX
}
// middleware

func BlogHandler(rw http.ResponseWriter, req *http.Request) {
	// handle blog route
	if req.Method != http.MethodPost {
		//rw.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	fmt.Printf("User agent: %s\n", req.UserAgent())
	fmt.Printf("accepts: %s\n", req.Header.Get("Accept"))
	blogPosts := ""
	blogResponse := `
	<h1>Welcome to my blog page</h1>
	%s
	<form method="POST" action="/blog/create">
		<input placeholder="Title" name="title" type="text"/>
		<br/><br/>
		<textarea name="body" ></textarea>
		<br/><br/>
		<input name="submit" type="submit"/>
	</form>
	`
	rw.WriteHeader(http.StatusOK)
	for i, v := range Posts {
		blogPosts += fmt.Sprintf("<h4>%d %s</h4><br/><p>%s</p><br/>", i, v.Title, v.Body)
	}
	rw.Write([]byte(fmt.Sprintf(blogResponse, blogPosts)))
}

type BlogPost struct {
	Title string
	Body string
}

var Posts = []BlogPost{}

func CreatBlogHandler(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("i want to create a blog")
	req.ParseForm()
	bp := BlogPost{
		Title: req.Form.Get("title"),
		Body: req.Form.Get("body"),
	}
	Posts = append(Posts, bp)

	rw.WriteHeader(http.StatusOK)
}

func start() {
	port := ":"
	port += os.Getenv("PORT")
	if port == ":" {
		port = ":3000"
	}
	fmt.Println("starting server on port: " + port)

	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/blog", BlogHandler)
	http.HandleFunc("/blog/create", CreatBlogHandler)

	http.ListenAndServe(port, nil)
}
