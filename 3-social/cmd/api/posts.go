package main

import (
	"net/http"

	"github.com/danivideda/go-backend-engineering-course/social/internal/store"
)

type CreatePostPayload struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Tags []string `json:"tags"`
}

func (app *application) createPostHandler(w http.ResponseWriter, r *http.Request) {
	var payload CreatePostPayload
	ctx := r.Context()

	if err := readJSON(w, r, &payload); err != nil {
		writeJSONError(w, http.StatusBadRequest, err.Error())
		return
	}

	post := &store.Post{
		Title:   payload.Title,
		Content: payload.Content,
		Tags: payload.Tags,
		// TODO: change after authentication
		UserID:  1,
	}

	if err := app.store.Posts.Create(ctx, post); err != nil {
		writeJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := writeJSON(w, http.StatusCreated, post); err != nil {
		writeJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}
}
