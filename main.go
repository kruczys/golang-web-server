package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"sync/atomic"
)

type SharkAttack struct {
	Date     string `json:"date"`
	Country  string `json:"country"`
	Name     string `json:"name"`
	Activity string `json:"activity"`
	Age      string `json:"age"`
	Injury   string `json:"injury"`
}

var (
	posts atomic.Value
)

func main() {
	filename, err := os.Open("global-shark-attack.json")

	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := filename.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	data, err := io.ReadAll(filename)

	if err != nil {
		log.Fatal(err)
	}

	var records []SharkAttack
	err = json.Unmarshal(data, &records)

	var initialPosts []SharkAttack
	for i := 0; i < 10; i++ {
		initialPosts = append(initialPosts, records[rand.Intn(len(records))])
	}

	posts.Store(initialPosts)

	http.HandleFunc("/posts", handlePosts)
	// http.HandleFunc("/posts/", handlePost)
	fmt.Println("Server is listening at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handlePosts(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		currentPosts := posts.Load().([]SharkAttack)
		json.NewEncoder(w).Encode(currentPosts)
    case http.MethodPost:
        var post SharkAttack
        if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }
        currentPosts := posts.Load().([]SharkAttack)
        newPosts := append(currentPosts, post)
        posts.Store(newPosts)
        w.WriteHeader(http.StatusCreated)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handlePost(w http.ResponseWriter, r *http.Request) {
    id := r.URL.Path[len("/posts/"):]
    currentPosts := posts.Load().([]SharkAttack)
    switch r.Method {
        case http.MethodGet:
            fmt.Println("naura")
            fmt.Println("test")
    }
}
