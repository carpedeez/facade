package facade

import (
	"context"
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
	r.Use(middleware.RealIP)
	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestLogger(&apiLogger{log}))

	handler := Handler(f, WithRouter(r), WithMiddleware("session", f.sessionMiddleware))

	go func() {
		log.Info().Msgf("listening on http://localhost:%s/", c.Port)
		err := http.ListenAndServe(":"+c.Port, handler)
		if err != nil {
			log.Panic().Msg("failed to start listening and serving api")
		}
	}()

	<-ctx.Done()
}
