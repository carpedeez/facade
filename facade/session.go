package facade

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	ory "github.com/ory/client-go"
)

type requestCtxKey string

const (
	requestCookies requestCtxKey = "req.cookies"
	requestSession requestCtxKey = "req.session"
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

func getFakeSession() *ory.Session {
	t := time.Now()
	vemailID := "testtest-test-test-test-test--vemail"
	method := "password"
	state := ory.IdentityState("active")
	active := true
	authenticatedAt := t
	authenticatorAssuranceLevel := ory.AuthenticatorAssuranceLevel("aal1")
	authenticationMethods := []ory.SessionAuthenticationMethod{
		{
			Aal:         &authenticatorAssuranceLevel,
			CompletedAt: &authenticatedAt,
			Method:      &method,
		},
	}
	expiresAt := t.Add(24 * time.Hour)
	id := "testtest-test-test-test-testtest--id"
	identity := ory.Identity{
		CreatedAt:      &t,
		Credentials:    &map[string]ory.IdentityCredentials{},
		Id:             "testtest-test-test-test-test----user",
		MetadataPublic: nil,
		RecoveryAddresses: []ory.RecoveryAddress{
			{
				CreatedAt: &t,
				Id:        "testtest-test-test-test-test--remail",
				UpdatedAt: &t,
				Value:     "test@facadeapp.dev",
				Via:       "email",
			},
		},
		SchemaId:       "preset://email",
		SchemaUrl:      "http://localhost:4455/schemas/cHJlc2V0Oi8vZW1haWw",
		State:          &state,
		StateChangedAt: &t,
		Traits: map[string]any{
			"email": "test@facadeapp.dev",
			"name": map[string]any{
				"first": "Todd",
				"last":  "Something",
			},
		},
		UpdatedAt: &time.Time{},
		VerifiableAddresses: []ory.VerifiableIdentityAddress{
			{
				CreatedAt: &t,
				Id:        &vemailID,
				Status:    "sent",
				UpdatedAt: &t,
				Value:     "test@facadeapp.dev",
				Verified:  false,
				Via:       "email",
			},
		},
	}
	issuedAt := t
	session := &ory.Session{
		Active:                      &active,
		AuthenticatedAt:             &authenticatedAt,
		AuthenticationMethods:       authenticationMethods,
		AuthenticatorAssuranceLevel: &authenticatorAssuranceLevel,
		ExpiresAt:                   &expiresAt,
		Id:                          id,
		Identity:                    identity,
		IssuedAt:                    &issuedAt,
	}
	return session
}

func fakeSessionMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		cookies := r.Header.Get("Cookie")
		ctx := withCookies(r.Context(), cookies)

		session := getFakeSession()

		ctx = withSession(ctx, session)

		// continue to the requested page (in our case the Dashboard)
		next.ServeHTTP(w, r.WithContext(ctx))
	}

	return http.HandlerFunc(fn)
}

func (f facade) FakeMe(w http.ResponseWriter, r *http.Request) {
	session := getFakeSession()

	bytes, _ := json.Marshal(session)
	w.Write(bytes)
}
