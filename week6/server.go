package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

// spaHandler implements the http.Handler interface, so we can use it
// to respond to HTTP requests. The path to the static directory and
// path to the index file within that static directory are used to
// serve the SPA in the given static directory.
type spaHandler struct {
	staticPath string
	indexPath  string
}

// ServeHTTP inspects the URL path to locate a file within the static dir
// on the SPA handler. If a file is found, it will be served. If not, the
// file located at the index path on the SPA handler will be served. This
// is suitable behavior for serving an SPA (single page application).
func (h spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// get the absolute path to prevent directory traversal
	path, err := filepath.Abs(r.URL.Path)
	if err != nil {
		// if we failed to get the absolute path respond with a 400 bad request
		// and stop
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// prepend the path with the path to the static directory
	path = filepath.Join(h.staticPath, path)

	// check whether a file exists at the given path
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		// file does not exist, serve index.html
		http.ServeFile(w, r, filepath.Join(h.staticPath, h.indexPath))
		return
	} else if err != nil {
		// if we got an error (that wasn't that the file doesn't exist) stating the
		// file, return a 500 internal server error and stop
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// otherwise, use http.FileServer to serve the static dir
	http.FileServer(http.Dir(h.staticPath)).ServeHTTP(w, r)
}

type Todo struct {
	Name string
	Description string
}

var DB = []*Todo{}

func GetTodos(responseWriter http.ResponseWriter, rew *http.Request) {
	todos, err := json.Marshal(DB)
	if err != nil {
		log.Println(err)
		responseWriter.WriteHeader(http.StatusInternalServerError)
		responseWriter.Write([]byte("{'error': 'internal server error'}"))
		return
	}

	responseWriter.WriteHeader(http.StatusOK)
	responseWriter.Write(todos)
}

func CreateTodo(responseWriter http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println(err)
		responseWriter.WriteHeader(http.StatusInternalServerError)
		responseWriter.Write([]byte("{'error': 'internal server error'}"))
		return
	}
	todo := &Todo{}
	err = json.Unmarshal(body, todo)
	if err != nil {
		log.Println(err)
		responseWriter.WriteHeader(http.StatusBadRequest)
		responseWriter.Write([]byte("{'error': 'bad request'}"))
		return
	}
	DB = append(DB, todo)
	responseWriter.WriteHeader(http.StatusCreated)
	responseWriter.Write([]byte("{'message': 'todo created successfully'}"))
}

func RestMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		header := rw.Header()
		header.Add("Content-Type", "application/json")
		header.Add("Squad", "squad-8")
		next.ServeHTTP(rw, req)
	})
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		// an example API handler
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	})

	apiRouter := router.PathPrefix("/api").Subrouter()

	apiRouter.Use(RestMiddleware)

	apiRouter.HandleFunc("/todos", GetTodos).Methods(http.MethodGet)
	apiRouter.HandleFunc("/todos", CreateTodo).Methods(http.MethodPost)
	apiRouter.HandleFunc("/todos/:id", CreateTodo).Methods(http.MethodGet)
	// http://localhost:8000/api/todos/1

	spa := spaHandler{staticPath: "build", indexPath: "index.html"}
	router.PathPrefix("/").Handler(spa)
	// http://localhost:8000/api/todos?query=Spankie&name=David&squad=squad8&
	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("server started at ", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}
