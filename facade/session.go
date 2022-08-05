package facade

import (
	"context"
	"fmt"
	"net/http"

	ory "github.com/ory/client-go"
)

type requestHeader string

var (
	requestCookies requestHeader = "req.cookies"
	requestSession requestHeader = "req.session"
)

// save the cookies for any upstream calls to the Ory apis
func withCookies(ctx context.Context, v string) context.Context {
	return context.WithValue(ctx, requestCookies, v)
}

// save the session to display it on the dashboard
func withSession(ctx context.Context, v *ory.Session) context.Context {
	return context.WithValue(ctx, requestSession, v)
}

func getSession(ctx context.Context) *ory.Session {
	return ctx.Value(requestSession).(*ory.Session)
}

func (facade facade) needsSessionMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// implement checking if sessionMiddleware needs to be done (it is not needed on most GETs, but is definitely on all POST, PATCH, DELETE, etc. requests)
		next.ServeHTTP(w, r)
	}
}

func (facade facade) sessionMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		fmt.Printf("handling middleware request\n")

		// this example passes all request.Cookies
		// to `ToSession` function
		//
		// However, you can pass only the value of
		// ory_session_projectid cookie to the endpoint
		cookies := request.Header.Get("Cookie")

		// check if we have a session
		session, _, err := facade.ory.V0alpha2Api.ToSession(request.Context()).Cookie(cookies).Execute()
		fmt.Printf("error: %v, session: %+v", err, session)
		if (err != nil && session == nil) || (err == nil && !*session.Active) {
			// this will redirect the user to the managed Ory Login UI
			http.Redirect(writer, request, "/.ory/self-service/login/browser", http.StatusSeeOther)
			return
		}

		ctx := withCookies(request.Context(), cookies)
		ctx = withSession(ctx, session)

		fmt.Printf("cookies: %v", cookies)

		// continue to the requested page (in our case the Dashboard)
		next.ServeHTTP(writer, request.WithContext(ctx))
	}
}
