package main

import (
	"net/http"

	"github.com/danivideda/go-backend-engineering-course/social/internal/store"
)

type CreateCommentPayload struct {
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

	post := getPostFromCtx(ctx)
	comment := &store.Comment{
		PostID:  post.ID,
		UserID:  1,
		Content: payload.Content,
	}

	if err := app.store.Comments.Create(ctx, comment); err != nil {
		app.internalServerErrorResponse(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusCreated, comment); err != nil {
		app.internalServerErrorResponse(w, r, err)
		return
	}
}

func (app *application) getCommentsHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	post := getPostFromCtx(ctx)

	comments, err := app.store.Comments.GetByPostID(ctx, post.ID)
	if err != nil {
		app.internalServerErrorResponse(w, r, err)
		return
	}

	if err = app.jsonResponse(w, http.StatusOK, comments); err != nil {
		app.internalServerErrorResponse(w, r, err)
		return
	}
}
