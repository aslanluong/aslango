package routes

import (
	"aslango/pkg/api"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/go-chi/chi/v5"
)

func LinkRoute(r chi.Router) {
	r.Use(cacheMiddleware)
	r.Get("/create/*", getShortLink)
	r.Get("/*", goShortLink)
}

type ErrorResponse struct {
	Status int    `bson:"status" json:"status"`
	Error  string `bson:"error" json:"error"`
}

type ShortLinkResponse struct {
	Status    int    `json:"status"`
	ShortLink string `json:"short_link"`
}

func getShortLink(w http.ResponseWriter, r *http.Request) {
	originalLink := chi.URLParam(r, "*")
	if match, _ := regexp.MatchString("https://.+", originalLink); !match {
		json.NewEncoder(w).Encode(ErrorResponse{
			Status: 400,
			Error:  "This link is not supported yet!",
		})
		return
	}

	link, err := api.GenerateShortLink(originalLink)
	if err != nil {
		fmt.Println("error")
		return
	}
	shortLink := r.Host + strings.Replace(r.RequestURI, chi.RouteContext(r.Context()).RoutePath, "/"+link.ShortLink, 1)
	json.NewEncoder(w).Encode(ShortLinkResponse{
		Status:    200,
		ShortLink: shortLink,
	})
}

func goShortLink(w http.ResponseWriter, r *http.Request) {
	if link, err := api.GetOriginalLink(chi.URLParam(r, "*")); err != nil {
		json.NewEncoder(w).Encode(ErrorResponse{
			Status: 404,
			Error:  "Short link not found!",
		})
	} else {
		api.UpdateLinkActive(link.LinkId)
		http.Redirect(w, r, link.OriginalLink, http.StatusMovedPermanently)
	}
}

func cacheMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "max-age=43200,s-maxage=43200")
		next.ServeHTTP(w, r)
	})
}
