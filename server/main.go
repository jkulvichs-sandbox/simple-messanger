package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type Message struct {
	Text string `json:"text"`
}

func main() {
	messages := make([]Message, 0)

	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://*"},
		AllowedMethods: []string{"GET", "POST"},
	}))

	r.Get("/api/echo", func(w http.ResponseWriter, r *http.Request) {
		data, _ := ioutil.ReadAll(r.Body)
		w.Write(data)
	})
	r.Post("/api/messages", func(w http.ResponseWriter, r *http.Request) {
		data, _ := ioutil.ReadAll(r.Body)
		fmt.Println("MSG SAVED:", string(data))
		msg := Message{}
		if err := json.Unmarshal(data, &msg); err != nil {
			fmt.Errorf("can't unmarshal: %s", err)
		}
		messages = append(messages, msg)
	})
	r.Get("/api/messages/{id}", func(w http.ResponseWriter, r *http.Request) {
		str_id := chi.URLParam(r, "id")
		id, err := strconv.Atoi(str_id)
		fmt.Println("MSG SENT:", str_id)
		if err != nil {
			log.Fatalf("can't parse id: %s", err)
		}
		if id >= 0 && id < len(messages) {
			msg := messages[id]
			data, _ := json.Marshal(msg)
			w.Write(data)
		}
	})

	fmt.Println("Ready")
	http.ListenAndServe(":5000", r)
}
