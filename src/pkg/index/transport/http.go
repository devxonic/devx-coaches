package transport

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	api "github.com/ibad69/go-devxcoaches/src/pkg"
	"github.com/ibad69/go-devxcoaches/src/pkg/index"
	repo "github.com/ibad69/go-devxcoaches/src/pkg/index/store"
)

type handler struct {
	indexService index.Service
}

func Activate(ro chi.Router, db *pgxpool.Pool) {
	service := index.New(repo.New(db))
	NewHandler(service, ro)
}

func NewHandler(s index.Service, ro chi.Router) {
	h := handler{indexService: s}
	ro.Route("/api", func(r chi.Router) {
		r.Post("/addStudent", h.AddStudent)
		r.Post("/addClass", h.AddClass)
		r.Post("/addSubject", h.AddSubject)
		r.Post("/addBatch", h.AddBatch)
		r.Get("/getStudents", h.GetStudents)
		r.Get("/getClass", h.GetClass)
		r.Get("/getSubjects", h.GetSubjects)
		r.Get("/getBatches", h.GetBatches)
	})
}

func (h *handler) AddStudent(w http.ResponseWriter, r *http.Request) {
	var s api.Student
	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	results, err := h.indexService.AddStudent(r.Context(), s)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
	}
	json.NewEncoder(w).Encode(results)
}

func (h *handler) AddClass(w http.ResponseWriter, r *http.Request) {
	var s api.Class
	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	results, err := h.indexService.AddClass(r.Context(), s)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
	}
	json.NewEncoder(w).Encode(results)
}

func (h *handler) AddSubject(w http.ResponseWriter, r *http.Request) {
	var s api.Subject
	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	results, err := h.indexService.AddSubject(r.Context(), s)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
	}
	json.NewEncoder(w).Encode(results)
}

func (h *handler) AddBatch(w http.ResponseWriter, r *http.Request) {
	var s api.Batch
	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	results, err := h.indexService.AddBatch(r.Context(), s)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
	}
	json.NewEncoder(w).Encode(results)
}

func (h *handler) GetStudents(w http.ResponseWriter, r *http.Request) {
	var s api.Student
	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	results, err := h.indexService.GetStudents(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
	}
	json.NewEncoder(w).Encode(results)
}

func (h *handler) GetClass(w http.ResponseWriter, r *http.Request) {
	var s api.Student
	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	results, err := h.indexService.GetClass(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
	}
	json.NewEncoder(w).Encode(results)
}

func (h *handler) GetSubjects(w http.ResponseWriter, r *http.Request) {
	var s api.Student
	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	results, err := h.indexService.GetSubjects(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
	}
	json.NewEncoder(w).Encode(results)
}

func (h *handler) GetBatches(w http.ResponseWriter, r *http.Request) {
	var s api.Batch
	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	results, err := h.indexService.GetBatches(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
	}
	json.NewEncoder(w).Encode(results)
}
