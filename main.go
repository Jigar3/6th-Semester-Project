package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jigar3/hardware_counters/routes"
)

func main() {

	r := mux.NewRouter()
	headersOk := handlers.AllowedHeaders([]string{"Content-Type", "Access-Control-Allow-Origin", "*"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"})

	r.HandleFunc("/uploadFile", routes.GetCFileAndPerform).Methods("POST")

	if err := http.ListenAndServe(":3000", handlers.CORS(headersOk, originsOk, methodsOk)(r)); err != nil {
		log.Fatal(err)
	}

	// cache_performance := run("hello_world.c")

	// fmt.Println("The value of performance counters are ->")
	// for k, v := range cache_performance {
	// 	fmt.Println(k, " = ", v)
	// }
}

// func run(filename string) map[string]string {
// 	compileACProgram(filename)
// 	runValgrind(filename[:len(filename)-2])
// 	return getValuesFromCallgrindFile("cg_with_branch_sim.out")
// }
