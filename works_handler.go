package handler

import (
	"database/sql"
	"encoding/json"
	"internal/domain"
	"internal/usecase"
	"net/http"

	"strconv"

	"github.com/gorilla/mux"
)

type WorkHandler struct {
	WorkUsecase *usecase.WorkUsecase
	DB          *sql.DB
}

func NewWorkHandler(db *sql.DB) *WorkHandler {
	return &WorkHandler{DB: db}
}

func (h *WorkHandler) Creatework(w http.ResponseWriter, r *http.Request) {
	var work domain.Works
	err := json.NewDecoder(r.Body).Decode(&work)
	if err := json.NewDecoder(r.Body).Decode(&work); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.WorkUsecase.Creatework(&work); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = h.DB.Exec("INSERT INTO works (title, user_id) VALUES (?, ?)", work.Title, work.UserID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(work)
}

func (h *WorkHandler) GetworkByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	work, err := h.WorkUsecase.GetworkByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(work)
}

func (h *WorkHandler) GetWork(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var work domain.Work
	err := h.DB.QueryRow("SELECT id, title, user_id FROM works WHERE id = ?", id).Scan(&work.ID, &work.Title, &work.UserID)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Work not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(work)
}
