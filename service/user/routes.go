package user

import (
	"net/http"

	"github.com/RobTov/hmblog-golang-backend/types"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {}).Methods(http.MethodGet)
	router.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {}).Methods(http.MethodPost)
	router.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {}).Methods(http.MethodPost)
	router.HandleFunc("/user/{id}", func(w http.ResponseWriter, r *http.Request) {}).Methods(http.MethodPut)
	router.HandleFunc("/user/{id}", func(w http.ResponseWriter, r *http.Request) {}).Methods(http.MethodDelete)
}
