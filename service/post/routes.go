package post

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/RobTov/hmblog-golang-backend/service/auth"
	"github.com/RobTov/hmblog-golang-backend/types"
	"github.com/RobTov/hmblog-golang-backend/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type Handler struct {
	PostStore types.PostStore
	UserStore types.UserStore
}

func NewHandler(postStore types.PostStore, userStore types.UserStore) *Handler {
	return &Handler{PostStore: postStore, UserStore: userStore}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/posts", h.handleGetPosts).Methods(http.MethodGet)
	router.HandleFunc("/posts/{user_id}", h.handleGetPostsByUserID).Methods(http.MethodGet)
	router.HandleFunc("/post", auth.WithJWTAuth(h.handleCreatePost, h.UserStore)).Methods(http.MethodPost)
	router.HandleFunc("/post/{id}", auth.WithJWTAuth(h.handleUpdatePost, h.UserStore)).Methods(http.MethodPut)
	router.HandleFunc("/post/{id}", h.handleDeletePost).Methods(http.MethodDelete)

}

func (h *Handler) handleGetPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := h.PostStore.GetPosts()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, posts)
}

func (h *Handler) handleGetPostsByUserID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	stringUserID, ok := vars["user_id"]
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("userID not provided"))
		return
	}

	userID, _ := strconv.Atoi(stringUserID)

	posts, err := h.PostStore.GetPostsByUserID(uint(userID))
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, posts)
}

func (h *Handler) handleCreatePost(w http.ResponseWriter, r *http.Request) {
	var userID = auth.GetUserIDFromContext(r.Context())
	var payload types.PostCreateAndUpdatePayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	err := h.PostStore.CreatePost(types.Post{
		Title:            payload.Title,
		ShortDescription: payload.ShortDescription,
		Body:             payload.Body,
		Image:            payload.Image,
		UserID:           userID,
	})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, nil)
}

func (h *Handler) handleUpdatePost(w http.ResponseWriter, r *http.Request) {
	userID := auth.GetUserIDFromContext(r.Context())
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid postID"))
		return
	}
	postID, _ := strconv.Atoi(id)

	var payload types.PostCreateAndUpdatePayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	err := h.PostStore.UpdatePost(types.Post{
		ID:               uint(postID),
		Title:            payload.Title,
		ShortDescription: payload.ShortDescription,
		Body:             payload.Body,
		Image:            payload.Image,
		UserID:           userID,
	})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, nil)
}

func (h *Handler) handleDeletePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid postID"))
		return
	}

	postID, _ := strconv.Atoi(id)

	err := h.PostStore.DeletePost(types.Post{
		ID: uint(postID),
	})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, nil)
}
