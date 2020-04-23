package main

import (
	"log"
	"net/http"
)

// func main() {
// 	rtr := mux.NewRouter()

// 	rtr.HandleFunc("/{topic}", func(w http.ResponseWriter, r *http.Request) {
// 		vars := mux.Vars(r)
// 		w.Write([]byte("Topic: " + vars["topic"]))
// 	})

// 	http.Handle("/", rtr)

// 	if err := http.ListenAndServe(":3005", nil); err != nil {
// 		log.Fatal(err)
// 	}
// }

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	if err := http.ListenAndServe(":3005", nil); err != nil {
		log.Fatal(err)
	}
}
