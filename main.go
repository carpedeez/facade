package main

import (
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	ory "github.com/ory/client-go"

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

	f := facade.Facade{Querier: &database.Querier{DB: db}}
	handler := facade.Handler(f, facade.WithRouter(r))

	// Kratos Testing

	c := ory.NewConfiguration()
	c.Servers = ory.ServerConfigurations{{URL: fmt.Sprintf("http://localhost:%s/", "4433")}}

	app := &App{
		ory: ory.NewAPIClient(c),
	}

	r.Handle("/", app.sessionMiddleware(app.dashboardHandler()))

	// Kratos testing

	go func() {
		err = http.ListenAndServe(fmt.Sprintf(":%s", config.FacadeConfig.Port), handler)
		if err != nil {
			fmt.Printf("failed to listen and serve api: %v", err)
		}
	}()

	<-ctx.Done()
}

// Kratos Testing

type App struct {
	ory *ory.APIClient
}

// save the cookies for any upstream calls to the Ory apis
func withCookies(ctx context.Context, v string) context.Context {
	return context.WithValue(ctx, "req.cookies", v)
}

func getCookies(ctx context.Context) string {
	return ctx.Value("req.cookies").(string)
}

// save the session to display it on the dashboard
func withSession(ctx context.Context, v *ory.Session) context.Context {
	return context.WithValue(ctx, "req.session", v)
}

func getSession(ctx context.Context) *ory.Session {
	return ctx.Value("req.session").(*ory.Session)
}

func (app *App) sessionMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		log.Printf("handling middleware request\n")

		// set the cookies on the ory client
		var cookies string

		// this example passes all request.Cookies
		// to `ToSession` function
		//
		// However, you can pass only the value of
		// ory_session_projectid cookie to the endpoint
		cookies = request.Header.Get("Cookie")

		// check if we have a session
		session, _, err := app.ory.V0alpha2Api.ToSession(request.Context()).Cookie(cookies).Execute()
		if (err != nil && session == nil) || (err == nil && !*session.Active) {
			// this will redirect the user to the managed Ory Login UI
			http.Redirect(writer, request, "/login", http.StatusSeeOther)
			return
		}

		ctx := withCookies(request.Context(), cookies)
		ctx = withSession(ctx, session)

		// continue to the requested page (in our case the Dashboard)
		next.ServeHTTP(writer, request.WithContext(ctx))
	}
}

func (app *App) dashboardHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		tmpl, err := template.New("index.html").ParseFiles("index.html")
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		session, err := json.Marshal(getSession(request.Context()))
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		err = tmpl.ExecuteTemplate(writer, "index.html", string(session))
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

// Kratos Testing
