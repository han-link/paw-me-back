package main

import (
	"context"
	"errors"
	"net/http"
	"paw-me-back/internal/store"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/supertokens/supertokens-golang/recipe/emailpassword"
	"github.com/supertokens/supertokens-golang/recipe/session"
)

func (app *application) AuthMiddleware(next http.Handler) http.Handler {
	return session.VerifySession(nil, func(w http.ResponseWriter, r *http.Request) {
		sessionContainer := session.GetSessionFromRequestContext(r.Context())

		if sessionContainer == nil {
			app.internalServerError(w, r, errors.New("session not found"))
			return
		}

		userInfo, err := emailpassword.GetUserByID(sessionContainer.GetUserID())

		if err != nil {
			app.internalServerError(w, r, err)
		}

		superTokenId, err := uuid.Parse(sessionContainer.GetUserID())

		if err != nil {
			app.internalServerError(w, r, err)
		}

		ctx := r.Context()

		user, err := app.store.Users.GetOrCreateBySuperTokenID(ctx, superTokenId, strings.Split(userInfo.Email, "@")[0], userInfo.Email)
		if err != nil {
			app.unauthorizedErrorResponse(w, r, err)
			return
		}

		ctx = context.WithValue(ctx, userCtx, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (app *application) groupsContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		idParam := chi.URLParam(r, "groupId")
		groupId, err := uuid.Parse(idParam)

		if err != nil {
			app.internalServerError(w, r, err)
			return
		}

		ctx := r.Context()

		group, err := app.store.Groups.GetByID(ctx, groupId)
		if err != nil {
			switch {
			case errors.Is(err, store.ErrNotFound):
				app.notFoundResponse(w, r, err)
			default:
				app.internalServerError(w, r, err)
			}
			return
		}

		ctx = context.WithValue(ctx, groupCtx, group)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (app *application) CheckGroupMembership(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}
