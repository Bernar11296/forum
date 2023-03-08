package handler

import (
	"context"
	"errors"
	"forum/models"
	"net/http"
	"time"
)

const ctxUserKey ctxKey = 0

type ctxKey int8

func (h *Handler) middleWare(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		c, err := r.Cookie("session_token")
		if err != nil {
			switch {
			case errors.Is(err, http.ErrNoCookie):
				handler.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxUserKey, models.User{})))
				return
			case errors.Is(err, c.Valid()):
				h.ErrorHandle(w, http.StatusBadRequest, "invalid cookie value")
			}
			h.ErrorHandle(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
			return
		}
		user, err = h.services.Authorization.GetUserByToken(c.Value)
		if err != nil {
			handler.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxUserKey, models.User{})))
			return
		}
		if user.TokenDuration.Before(time.Now()) {
			if err := h.services.DeleteToken(user.Token); err != nil {
				h.ErrorHandle(w, http.StatusInternalServerError, http.StatusText(500))
				return
			}
			http.Redirect(w, r, "/auth/sign-in", http.StatusSeeOther)
			return
		}
		handler.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxUserKey, user)))
	}
}
