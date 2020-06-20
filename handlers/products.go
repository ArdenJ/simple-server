package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ardenj/graceful-http/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		p.getProducts(w, r)
		break
	case http.MethodPost:
		p.addProducts(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// getProducts is an internal function that implements the GetProducts method on the Products type '../data/products.go'
func (p *Products) getProducts(w http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	// d, err := json.Marshal(lp)
	err := lp.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to marshal JSON", http.StatusInternalServerError)
	}

	// w.Write(d)
}

func (p *Products) marGetProducts(w http.ResponseWriter, r *http.Request) {
	lp := data.MarGetProducts()
	d, err := json.Marshal(lp)
	// err := lp.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to marshal JSON", http.StatusInternalServerError)
	}

	w.Write(d)
}

func (p *Products) addProducts(w http.ResponseWriter, r *http.Request) {
	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "Unable to unMarshal Request body", http.StatusBadRequest)
	}

	p.l.Printf("Prod %#v", prod)
	data.AddProduct(prod)
}
