package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Howdy struct {
	l *log.Logger
}

func NewHowdy(l *log.Logger) *Howdy {
	return &Howdy{l}
}

func (h *Howdy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.l.Println("Howdy 👋")

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "ahh... no", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Howdy %s", data)
}
