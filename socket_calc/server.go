package main

import (
    "fmt"
    "html"
    "log"
    "net/http"
	"github.com/gorilla/mux"
)
type Para struct {
    OP string `json:"op`
    A float64 `json:"a"`
    B float64`json:"b"`
}
var ops []Para
func main() {
	router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/add", add).Methods("GET")
    router.HandleFunc("/sub", sub).Methods("GET")
    router.HandleFunc("/mul", mul).Methods("GET")
	router.HandleFunc("/div", div).Methods("GET")
	log.Fatal(http.ListenAndServe(":8081", router))
}
func middleware(w http.ResponseWriter, req *http.Request){
	params := mux.Vars(req)
    for _, item := range ops {
		fmt.Println(item)
        // if item.ID == params["id"] {
        //     json.NewEncoder(w).Encode(item)
        //     return
        // }
    }
}
// func add(a,b float64)float64{return a+b}
// func sub(a,b float64)float64{return a-b}
// func mul(a,b float64)float64{return a*b}
// func div(a,b float64)float64{return a/b}