package routes

import (
	"aslango/pkg/api"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func LinkRoute(r chi.Router)  {
	r.Get("/create/*", getShortLink)
	r.Get("/*", goShortLink)
}

func getShortLink(w http.ResponseWriter, r *http.Request)  {
}

func goShortLink(w http.ResponseWriter, r *http.Request) {
	type ErrorResponse struct {
		Status int `bson:"status" json:"status"`
		Error string `bson:"error" json:"error"`
	}

	if link, err := api.GetOriginalLink(r.RequestURI[4:]); err != nil {
		json.NewEncoder(w).Encode(ErrorResponse{
			Status: 404,
			Error: "Short link not found!",
		})
	} else {
		http.Redirect(w, r, *link, http.StatusMovedPermanently)
	}
}