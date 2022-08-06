package facade

import (
	"context"
	"net/http"

	ory "github.com/ory/client-go"
)

type requestHeader string

const (
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

func (f facade) sessionMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		cookies := r.Header.Get("Cookie")

		session, _, err := f.ory.V0alpha2Api.ToSession(r.Context()).Cookie(cookies).Execute()
		if (err != nil && session == nil) || (err == nil && !*session.Active) {
			http.Redirect(w, r, "/.ory/self-service/login/browser", http.StatusSeeOther)
			return
		}

		ctx := withCookies(r.Context(), cookies)
		ctx = withSession(ctx, session)

		// continue to the requested page (in our case the Dashboard)
		next.ServeHTTP(w, r.WithContext(ctx))
	}

	return http.HandlerFunc(fn)
}
