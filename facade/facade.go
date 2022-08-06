package facade

import (
	"fmt"
	"net/http"

	"github.com/carpedeez/store/config"
	"github.com/carpedeez/store/database"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	ory "github.com/ory/client-go"
)

type facade struct {
	querier database.Querier
	ory     *ory.APIClient
}

func Run(c config.FacadeConfig, q database.Querier) {
	f := facade{}

	oryC := ory.NewConfiguration()
	oryC.Servers = ory.ServerConfigurations{{URL: fmt.Sprintf("http://localhost:%s/", "4433")}}
	o := ory.NewAPIClient(oryC)

	f.querier = q
	f.ory = o

	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)

	handler := Handler(f, WithRouter(r), WithMiddleware("session", f.sessionMiddleware))

	go func() {
		err := http.ListenAndServe(fmt.Sprintf(":%s", c.Port), handler)
		if err != nil {
			fmt.Printf("failed to listen and serve api: %v", err)
		}
	}()
}
