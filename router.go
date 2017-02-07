package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
)

// Server informations
// TODO server config file
type Server struct {
	Port string
}

//Env structure hold db interface and other middleware structures
type Env struct {
	// No need database for now so commented
	// db dbstorage.DataStore
}

// Run starts the server
func run(httpHandlers http.Handler) {
	server := Server{":8080"}
	runHTTP(httpHandlers, server)
}

func runHTTP(handlers http.Handler, s Server) {
	log.Println("Starting local server on port", s.Port)

	log.Fatal(http.ListenAndServe(s.Port, handlers))
}

////////////////   ROUTES  ///////////////////

//DoRoutes creates routes and supply them with middleware
func doRoutes() http.Handler {
	//env := &Env{}
	r := httprouter.New()
	r.NotFound = http.HandlerFunc(notIndex)
	r.ServeFiles("/static/*filepath", http.Dir("static"))
	r.GET("/", index)
	return toHandler(r)
}

//Router to Handler
func toHandler(r *httprouter.Router) http.Handler {
	// Cors for cross-origin resources sharing needed
	// In dev phase ng2 :3000 port and api 8080
	return http.Handler(cors.Default().Handler(r))
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	index, err := ioutil.ReadFile("static/index.html")
	if err != nil {
		panic(err)
	}
	fmt.Fprint(w, string(index))
}

func notIndex(w http.ResponseWriter, r *http.Request) {
	index, err := ioutil.ReadFile("static/index.html")
	if err != nil {
		panic(err)
	}
	fmt.Fprint(w, string(index))
}
