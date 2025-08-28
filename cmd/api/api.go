package main

import (
	"net/http"
	"paw-me-back/internal/docs"
	"paw-me-back/internal/env"
	"paw-me-back/internal/store"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/supertokens/supertokens-golang/supertokens"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"go.uber.org/zap"
)

type application struct {
	config config
	logger *zap.SugaredLogger
	store  store.Storage
}

type config struct {
	addr        string
	apiURL      string
	frontendURL string
	env         string
}

func (app *application) mount() http.Handler {
	r := chi.NewRouter()

	corsOptions := cors.Options{
		AllowedOrigins:   strings.Split(env.GetString("CORS_ALLOWED_ORIGIN", ""), ","),
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   append([]string{"Content-Type"}, supertokens.GetAllCORSHeaders()...),
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler(corsOptions))
	r.Use(supertokens.Middleware)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/swagger/*", httpSwagger.Handler(httpSwagger.URL("doc.json")))

		r.Route("/groups", func(r chi.Router) {
			r.Use(app.AuthMiddleware)
			r.Get("/", app.getGroupsHandler)
			r.Post("/", app.createGroupHandler)

			r.Route("/{groupId}", func(r chi.Router) {
				r.Use(app.groupsContextMiddleware)
				r.Use(app.checkGroupMembership)
				r.Get("/", app.getGroupHandler)
				r.Put("/members", app.addMembersToGroupHandler)
			})
		})
	})

	return r
}

func (app *application) run(mux http.Handler) error {
	docs.SwaggerInfo.Version = version
	docs.SwaggerInfo.Host = app.config.apiURL
	docs.SwaggerInfo.BasePath = "/api/v1"

	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	app.logger.Infow("Server has started at", "addr", app.config.addr, "env", app.config.env)

	return srv.ListenAndServe()
}
