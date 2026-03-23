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
	r.Route("/users", func(func(r chi.Router)) {
		r.Post("/", Handler.CreateUser)
		r.Get("/", Handler.GetAllUsers)
		r.Get("/{id}", Handler.GetUserByID)
		r.Put("/{id}", Handler.UpdateUserAgeByID)
		r.Delete("/{id}", Handler.DeleteUserByID)
		r.Delete("/", Handler.DeleteAllUsers)
	})
}
