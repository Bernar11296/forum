package handler

import (
	"fmt"
	"forum/models"
	"html/template"
	"net/http"
)

func (h *Handler) ErrorHandle(w http.ResponseWriter, status int, message string) {
	errData := models.Error{
		Status:     status,
		StatusText: http.StatusText(status),
		Message:    message,
	}
	tml, err := template.ParseFiles("./ui/error.html")
	if err != nil {
		fmt.Printf("error handler parse files: %s/n", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), 500)
		return
	}
	w.WriteHeader(status)
	if err := tml.Execute(w, errData); err != nil {
		fmt.Printf("error handler execute: %s/n", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), 500)
		return
	}
}
