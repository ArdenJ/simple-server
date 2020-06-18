package handlers

import (
	"fmt"
	"log"
	"net/http"
)

type Boop struct {
	l *log.Logger
}

func NewBoop(l *log.Logger) *Boop {
	return &Boop{l}
}

func (b *Boop) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "boop")
}
