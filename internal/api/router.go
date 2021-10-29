package api

import (
	"net/http"

	"github.com/EmotivesProject/common/middlewares"
	"github.com/EmotivesProject/common/response"
	"github.com/EmotivesProject/common/verification"
	"github.com/go-chi/chi"
)

func CreateRouter() http.Handler {
	r := chi.NewRouter()

	r.Use(middlewares.SimpleMiddleware())

	r.Route("/", func(r chi.Router) {
		r.Get("/healthz", response.Healthz)
	})

	r.With(verification.VerifyJTW()).Route("/image", func(r chi.Router) {
		r.Post("/", uploadImage)
	})

	r.With(verification.VerifyJTW()).Route("/user_profile", func(r chi.Router) {
		r.Post("/", uploadUserImage)
	})

	return r
}
