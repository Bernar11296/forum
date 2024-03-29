package handler

import (
	"forum/models"
	"net/http"
	"strings"
)

func (h *Handler) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		h.ErrorHandle(w, http.StatusNotFound, http.StatusText(404))
		return
	}
	var user models.User
	var posts **[]models.Post
	var err error
	user = r.Context().Value(ctxUserKey).(models.User)
	switch r.Method {
	case http.MethodGet:
		r.ParseForm()
		category := strings.Join(r.Form["category"], " ")
		if category != "" {
			posts, err = h.services.GetPostByCategory(category)
			if err != nil {
				h.ErrorHandle(w, 500, http.StatusText(500))
				return
			}
		} else {
			posts, err = h.services.GetAllPost()
			if err != nil {
				h.ErrorHandle(w, 500, http.StatusText(500))
				return
			}
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	pageData := models.PageData{
		Username: user.Username,
		Posts:    **posts,
	}
	if err := h.tmlExecute(w, "./ui/index.html", pageData); err != nil {
		return
	}
}
