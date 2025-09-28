package main

import (
	"net/http"
	"strconv"

	"github.com/danivideda/go-backend-engineering-course/social/internal/store"
	"github.com/go-chi/chi/v5"
)

func (app *application) getUserHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userId, err := strconv.ParseInt(chi.URLParam(r, "userID"), 10, 64)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	user, err := app.store.Users.GetByID(ctx, userId)
	if err != nil {
		switch err {
		case store.ErrNotFound:
			app.notFoundResponse(w, r, err)
		default:
			app.internalServerErrorResponse(w, r, err)
		}

		return
	}

	if err = app.jsonResponse(w, http.StatusOK, user); err != nil {
		app.internalServerErrorResponse(w, r, err)
		return
	}
}
