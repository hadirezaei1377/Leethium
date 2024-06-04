package goAPI

import (
	"encoding/json"
	"net/http"
	"sync"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var (
	users = make(map[string]User)
	mutex = &sync.Mutex{}
)

func addUser(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mutex.Lock()
	users[user.ID] = user
	mutex.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mutex.Lock()
	if _, ok := users[user.ID]; ok {
		users[user.ID] = user
		mutex.Unlock()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
	} else {
		mutex.Unlock()
		http.Error(w, "User not found", http.StatusNotFound)
	}
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	mutex.Lock()
	if _, ok := users[id]; ok {
		delete(users, id)
		mutex.Unlock()
		w.WriteHeader(http.StatusOK)
	} else {
		mutex.Unlock()
		http.Error(w, "User not found", http.StatusNotFound)
	}
}

func main() {
	http.HandleFunc("/add", addUser)
	http.HandleFunc("/update", updateUser)
	http.HandleFunc("/delete", deleteUser)

	http.ListenAndServe(":8080", nil)
}
