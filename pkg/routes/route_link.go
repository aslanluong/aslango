package routes

import (
	"aslango/pkg/api"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func LinkRoute(r chi.Router) {
	r.Get("/create/*", getShortLink)
	r.Get("/*", goShortLink)
}

type ErrorResponse struct {
	Status int    `bson:"status" json:"status"`
	Error  string `bson:"error" json:"error"`
}

func getShortLink(w http.ResponseWriter, r *http.Request) {
	// domalnRegex := `https:\/\/(?:(?!(?:10|127)(?:\.\d{1,3}){3})(?!(?:169\.254|192\.168)(?:\.\d{1,3}){2})(?!172\.(?:1[6-9]|2\d|3[0-1])(?:\.\d{1,3}){2})(?:[1-9]\d?|1\d\d|2[01]\d|22[0-3])(?:\.(?:1?\d{1,2}|2[0-4]\d|25[0-5])){2}(?:\.(?:[1-9]\d?|1\d\d|2[0-4]\d|25[0-4]))|(?![-_])(?:[-\w\u00a1-\uffff]{0,63}[^-_]\.)+(?:[a-z\u00a1-\uffff]{2,}\.?))(?::\d{2,5})?(?:[/?#]\S*)?`
}

func goShortLink(w http.ResponseWriter, r *http.Request) {
	if link, err := api.GetOriginalLink(r.RequestURI[4:]); err != nil {
		json.NewEncoder(w).Encode(ErrorResponse{
			Status: 404,
			Error:  "Short link not found!",
		})
	} else {
		api.UpdateLinkActive(link.LinkId)
		http.Redirect(w, r, link.OriginalLink, http.StatusMovedPermanently)
	}
}
