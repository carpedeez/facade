package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/carpedeez/store/config"
	"github.com/carpedeez/store/database"
	"github.com/carpedeez/store/facade"
)

func main() {
	ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt)

	config, err := config.GetConfig()
	if err != nil {
		fmt.Printf("failed to get config: %v", err)
		os.Exit(1)
	}

	db, err := database.Open(ctx, config.DBConfig)
	if err != nil {
		fmt.Printf("failed to open database: %v", err)
		os.Exit(1)
	}

	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)

	f := facade.Facade{DB: db}
	handler := facade.HandlerWithOptions(f, facade.ChiServerOptions{
		BaseURL:    "",
		BaseRouter: r,
		ErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
		},
	})

	err = http.ListenAndServe(fmt.Sprintf(":%s", config.FacadeConfig.Port), handler)
	if err != nil {
		fmt.Printf("failed to listen and serve api: %v", err)
	}
}
