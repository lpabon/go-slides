// +build OMIT

package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/heketi/utils"
	"log"
	"net/http"
	"sync"
)

// start JSON structs OMIT
type App struct {
	keys map[string]string
	lock sync.Mutex
}

type AppValue struct {
	Value string `json:"value"`
}

type AppSaveRequest struct {
	AppValue
}

type AppGetResponse struct {
	AppValue
}

type AppGetAllResponse struct {
	Keys map[string]string `json:"keys"`
}

// end JSON structs OMIT

// New OMIT
func NewApp() *App {
	a := &App{}
	a.keys = make(map[string]string)

	return a
}

// Func Save OMIT
func (a *App) SaveKey(w http.ResponseWriter, r *http.Request) {
	var msg AppSaveRequest
	vars := mux.Vars(r)
	key := vars["key"]

	// Unmarshal using github.com/heketi/utils
	err := utils.GetJsonFromRequest(r, &msg)
	if err != nil {
		http.Error(w, "request unable to be parsed", 422)
		return
	}

	// Check information in JSON request
	if len(key) == 0 || len(msg.Value) == 0 {
		http.Error(w, "Missing infomration", http.StatusBadRequest)
		return
	}

	a.lock.Lock()
	a.keys[key] = msg.Value
	a.lock.Unlock()

	w.WriteHeader(http.StatusCreated)
}

// End Func Save OMIT

// Func Get OMIT
func (a *App) GetKey(w http.ResponseWriter, r *http.Request) {
	var msg AppGetResponse
	vars := mux.Vars(r)
	key := vars["key"]

	if len(key) == 0 {
		http.Error(w, "Key missing", http.StatusBadRequest)
		return
	}

	a.lock.Lock()
	defer a.lock.Unlock()
	if value, ok := a.keys[key]; ok {
		msg.Value = value
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(msg); err != nil {
			panic(err)
		}
	} else {
		http.Error(w, "Unknown Key: "+key, http.StatusNotFound)
		return
	}
}

// End Func Get OMIT

// Func GetAll OMIT
func (a *App) GetAllKeys(w http.ResponseWriter, r *http.Request) {
	var msg AppGetAllResponse

	a.lock.Lock()
	defer a.lock.Unlock()

	msg.Keys = a.keys
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(msg); err != nil {
		panic(err)
	}
}

// End Func GetAll OMIT

// Main OMIT
func main() {
	app := NewApp()
	r := mux.NewRouter()
	r.Methods("POST").Path("/x/{key}").HandlerFunc(app.SaveKey)
	r.Methods("GET").Path("/x/{key}").HandlerFunc(app.GetKey)
	r.Methods("GET").Path("/x").HandlerFunc(app.GetAllKeys)
	log.Fatal(http.ListenAndServe(":8080", r))
}

// End Main OMIT
