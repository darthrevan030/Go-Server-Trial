package user

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	repo Repository
}

func NewHandler(repo Repository) *Handler {
	return &Handler{repo: repo}
}

func (h Handler) CreateUser(w http.ResponseWriter, r *http.Request) {

}

func (h Handler) GetUserByID(w http.ResponseWriter, r *http.Request) {

}

func (h Handler) GetAllUsers(w http.ResponseWriter, r *http.Request) {

}

func (h Handler) UpdateUserAgeByID(w http.ResponseWriter, r *http.Request) {

}

func (h Handler) DeleteUserByID(w http.ResponseWriter, r *http.Request) {

}

func (h Handler) DeleteAllUsers(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) RegisterRoutes(r chi.Router) {
	r.Route("/users", func(r chi.Router)) {
		r.Post("/", h.CreateUser)
		r.Get("/", h.GetAllUsers)
		r.Get("/{id}", h.GetUserByID)
		r.Put("/{id}", h.UpdateUserAgeByID)
		r.Delete("/{id}", h.DeleteUserByID)
		r.Delete("/", h.DeleteAllUsers)
	})
}
