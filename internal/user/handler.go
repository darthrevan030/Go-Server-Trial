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

	// decode req body
	var req UserRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest) // 400
		return
	}

	// Above is the same as:
	// err := json.NewDecoder(r.Body).Decode(&req)
	// if err != nil {
	// 	http.Error(w, "Invalid request body", http.StatusBadRequest)
	// 	return
	// }

	// call the repository
	id, err := h.repo.CreateUser(User{
		Name:    req.Name,
		Age:     req.Age,
		Country: req.Country,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) // 500
		return
	}

	// send response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	json.NewEncoder(w).Encode(UserResponse{
		Data: map[string]string{"id": id},
	})

}

func (h Handler) GetUserByID(w http.ResponseWriter, r *http.Request) {

	// decode the req
	id := chi.URLParam(r, "id") // get id from url

	// call the repository
	user, err := h.repo.GetUserByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound) // 404
		return
	}

	// send response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200 
	json.NewEncoder(w).Encode(UserResponse{
		Data: user,
	})
}

func (h Handler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	// call the repository 
	users, err := h.repo.GetAllUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	return
	}

	// send response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	json.NewEncoder(w).Encode(UserResponse{
		Data: users,
	})
}

func (h Handler) UpdateUserAgeByID(w http.ResponseWriter, r *http.Request) {
	
	// decode the request
	id := chi.URLParam(r, "id") 

	var req UserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
	http.Error(w, "Invalid request body", http.StatusBadRequest) // 400
	return
}

	// call the repo
	result, err := h.repo.UpdateUserAgeByID(id, req.Age)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) // 500
		return
	}

	// send response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	json.NewEncoder(w).Encode(UserResponse{
		Data: result,
	})
}

func (h Handler) DeleteUserByID(w http.ResponseWriter, r *http.Request) {

	// decode the req body
	id := chi.URLParam(r, "id")

	// call the repo
	result, err := h.repo.DeleteUserByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) // 500
		return
	}

	// send the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	json.NewEncoder(w).Encode(UserResponse{
		Data: result,
	})
}

func (h Handler) DeleteAllUsers(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) RegisterRoutes(r chi.Router) {
	r.Route("/users", func(r chi.Router) {
		r.Post("/", h.CreateUser)
		r.Get("/", h.GetAllUsers)
		r.Get("/{id}", h.GetUserByID)
		r.Put("/{id}", h.UpdateUserAgeByID)
		r.Delete("/{id}", h.DeleteUserByID)
		r.Delete("/", h.DeleteAllUsers)
	})
}
