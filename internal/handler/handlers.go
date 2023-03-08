package handler

import (
	"forum/internal/service"
	"net/http"
	"text/template"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) tmlExecute(w http.ResponseWriter, path string, data interface{}) error {
	tml, err := template.ParseFiles(path)
	if err != nil {
		h.ErrorHandle(w, 500, http.StatusText(500))
		return err
	}
	if err := tml.Execute(w, data); err != nil {
		h.ErrorHandle(w, 500, http.StatusText(500))
		return err
	}
	return nil

}

func (h *Handler) InitRoutes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", h.middleWare(h.home))
	mux.HandleFunc("/auth/sign-up", h.signUp)
	mux.HandleFunc("/auth/sign-in", h.signIn)
	mux.HandleFunc("/log-out", h.logOut)
	mux.HandleFunc("/post/create", h.middleWare(h.createPost))
	mux.HandleFunc("/my-posts", h.middleWare(h.myPosts))
	mux.HandleFunc("/my-favourites", h.middleWare(h.myFavourites))
	mux.HandleFunc("/post/", h.middleWare(h.post))
	mux.HandleFunc("/like-post", h.middleWare(h.likePost))
	mux.HandleFunc("/dislike-post", h.middleWare(h.dislikePost))
	mux.HandleFunc("/like-comment", h.middleWare(h.likeComment))
	mux.HandleFunc("/dislike-comment", h.middleWare(h.dislikeComment))
	mux.Handle("/ui/css/", http.StripPrefix("/ui/css/", http.FileServer(http.Dir("./ui/css/"))))
	return mux
}
