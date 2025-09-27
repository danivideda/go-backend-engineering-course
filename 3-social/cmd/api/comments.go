package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/danivideda/go-backend-engineering-course/social/internal/store"
)

type CreateCommentPayload struct {
	PostID  int64  `json:"post_id" validate:"required"`
	Content string `json:"content" validate:"required,min=5,max=1000"`
}

func (app *application) createCommentHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var payload CreateCommentPayload

	if err := readJSON(w, r, &payload); err != nil {
		app.internalServerErrorResponse(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	comment := &store.Comment{
		PostID:  payload.PostID,
		UserID:  1,
		Content: payload.Content,
	}

	if _, err := app.store.Posts.GetByID(ctx, comment.PostID); err != nil {
		switch {
		case errors.Is(err, store.ErrNotFound):
			app.notFoundResponse(w, r, err)
		default:
			app.internalServerErrorResponse(w, r, err)
		}
	}

	if err := app.store.Comments.Create(ctx, comment); err != nil {
		app.internalServerErrorResponse(w, r, err)
		return
	}

	if err := writeJSON(w, http.StatusCreated, comment); err != nil {
		app.internalServerErrorResponse(w, r, err)
		return
	}
}

func (app *application) getCommentsByPostIDHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	postId := r.URL.Query().Get("postId")
	if postId == "" {
		app.badRequestResponse(w, r, errors.New("postId is empty"))
		return
	}

	postIdInt, err := strconv.ParseInt(postId, 10, 64)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	comments, err := app.store.Comments.GetByPostID(ctx, postIdInt)
	if err != nil {
		app.internalServerErrorResponse(w, r, err)
		return
	}

	if err = writeJSON(w, http.StatusOK, comments); err != nil {
		app.internalServerErrorResponse(w, r, err)
		return
	}
}
