package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/olahol/melody"
)

func main() {
	m := melody.New()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			http.ServeFile(w, r, "index.html")
			return
		}

		if r.URL.Path[0:3] == "/ws" {
			id, err := strconv.Atoi(r.URL.Path[len(r.URL.Path)-2 : len(r.URL.Path)])
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
			}

			fmt.Println("request with keys")
			m.HandleRequestWithKeys(w, r, map[string]any{"id": id})
			fmt.Println("after request with keys")
			return
		}

		http.ServeFile(w, r, "chan.html")
	})

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		m.BroadcastFilter(msg, func(q *melody.Session) bool {
			return q.Keys["id"] == 13
		})
	})

	http.ListenAndServe(":5000", nil)
}
