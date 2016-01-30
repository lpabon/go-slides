// +build OMIT

package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type App struct {
	key string
}

func (a *App) SaveKey(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]
	a.key = key
	w.WriteHeader(http.StatusOK)
}

func (a *App) GetKey(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, a.key)
}

func main() {
	app := &App{}
	r := mux.NewRouter()
	r.Methods("GET").Path("/x").HandlerFunc(app.GetKey)
	r.Methods("POST").Path("/x/{key}").HandlerFunc(app.SaveKey)
	log.Fatal(http.ListenAndServe(":8080", r))
}
