package scootin

import (
	"context"
	"encoding/json"
	"errors"
	"html/template"
	"net/http"
)

type Storage interface {
	FindScooters(ctx context.Context, lat1, lng1, lat2, lng2 float64, status string) ([]Scooter, error)
	UpdateScooter(ctx context.Context, scooter Scooter) error
	Close() error
}

type HTTPHandler struct {
	mux     *http.ServeMux
	storage Storage
}

func NewHTTPHandler(s Storage) (*HTTPHandler, error) {
	mux := http.NewServeMux()
	hdl := &HTTPHandler{mux: mux, storage: s}

	hdl.mux.HandleFunc("GET /", hdl.Elements)
	hdl.mux.HandleFunc("GET /docs", hdl.OpenAPI)
	hdl.mux.HandleFunc("GET /api/v1/scooters", hdl.FindScooters)
	hdl.mux.HandleFunc("POST /api/v1/scooters/{id}", hdl.UpdateScooter)

	return hdl, nil
}

func (h *HTTPHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mux.ServeHTTP(w, r)
}

func (h *HTTPHandler) FindScooters(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	input, err := NewFindScootersInput(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	scooters, err := h.storage.FindScooters(ctx, input.Latitude1, input.Longitude1, input.Latitude2, input.Longitude2, string(input.Status))
	if errors.Is(err, ErrNoScootersFound) {
		http.Error(w, "no scooters found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(scooters); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *HTTPHandler) UpdateScooter(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	input, err := NewUpdateScooterInput(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	scooter := Scooter{
		ID:        input.ID,
		Status:    string(input.Status),
		Latitude:  input.Latitude,
		Longitude: input.Longitude,
	}

	err = h.storage.UpdateScooter(ctx, scooter)
	if errors.Is(err, ErrScooterNotFound) {
		http.Error(w, "scooter not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}

func (h *HTTPHandler) Elements(w http.ResponseWriter, _ *http.Request) {
	template.Must(template.New("elements").Parse(elements)).Execute(w, "./docs")
}

func (h *HTTPHandler) OpenAPI(w http.ResponseWriter, _ *http.Request) {
	template.Must(template.New("docs").Parse(docs)).Execute(w, nil)
}
