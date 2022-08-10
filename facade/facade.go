package facade

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/carpedeez/facade/config"
	"github.com/carpedeez/facade/database"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	ory "github.com/ory/client-go"
	"github.com/rs/zerolog"
)

type facade struct {
	log     zerolog.Logger
	querier database.Querier
	ory     *ory.APIClient
}

func Run(wg *sync.WaitGroup, ctx context.Context, log zerolog.Logger, c config.FacadeConfig, q database.Querier) {
	defer wg.Done()

	oryC := ory.NewConfiguration()
	oryC.Servers = ory.ServerConfigurations{{URL: fmt.Sprintf("http://localhost:%s/", "4433")}}
	o := ory.NewAPIClient(oryC)

	f := facade{log, q, o}

	r := chi.NewRouter()
	r.Use(middleware.RealIP, middleware.RequestID, middleware.Recoverer, newRequestLogger(log))

	r.Get("/v0/redoc", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<!DOCTYPE html>
		<html>
		  <head>
			<title>Redoc</title>
			<!-- needed for adaptive design -->
			<meta charset="utf-8"/>
			<meta name="viewport" content="width=device-width, initial-scale=1">
			<link href="https://fonts.googleapis.com/css?family=Montserrat:300,400,700|Roboto:300,400,700" rel="stylesheet">
		
			<!--
			Redoc doesn't change outer page styles
			-->
			<style>
			  body {
				margin: 0;
				padding: 0;
			  }
			</style>
		  </head>
		  <body>
			<redoc spec-url='http://localhost:4000/v0/spec.json'></redoc>
			<script src="https://cdn.redoc.ly/redoc/latest/bundles/redoc.standalone.js"> </script>
		  </body>
		</html>`))
	})

	r.Get("/v0/spec.json", func(w http.ResponseWriter, r *http.Request) {
		swagger, err := GetSwagger()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		json, err := json.Marshal(swagger)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.Write(json)
	})

	handler := Handler(f, WithRouter(r), WithMiddleware("session", f.sessionMiddleware), WithServerBaseURL("/v0"))

	go func() {
		log.Info().Msgf("listening on http://localhost:%s/", c.Port)
		err := http.ListenAndServe(":"+c.Port, handler)
		if err != nil {
			log.Panic().Msg("failed to start listening and serving api")
		}
	}()

	<-ctx.Done()
}
