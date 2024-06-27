package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/RobTov/hmblog-golang-backend/service/post"
	"github.com/RobTov/hmblog-golang-backend/service/user"
	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{addr: addr, db: db}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subrouter)

	postStore := post.NewStore(s.db)
	postHandler := post.NewHandler(postStore, userStore)
	postHandler.RegisterRoutes(subrouter)

	log.Printf("Server listening on: http://localhost%s\n", s.addr)

	return http.ListenAndServe(s.addr, router)
}
