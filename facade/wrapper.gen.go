// Package facade provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/discord-gophers/goapi-gen version v0.2.2 DO NOT EDIT.
package facade

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/discord-gophers/goapi-gen/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

// Error defines model for Error.
type Error struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

// GetDisplay defines model for GetDisplay.
type GetDisplay struct {
	Description string    `json:"description"`
	ID          string    `json:"id"`
	Items       []GetItem `json:"items"`
	PhotoURL    string    `json:"photoURL"`
	Title       string    `json:"title"`
	UserID      string    `json:"userID"`
}

// GetItem defines model for GetItem.
type GetItem struct {
	DisplayID      string `json:"displayID"`
	ExternalLink   string `json:"externalLink"`
	ID             string `json:"id"`
	PhotoURL       string `json:"photoURL"`
	SocialPostLink string `json:"socialPostLink"`
	UserID         string `json:"userID"`
}

// PatchDisplay defines model for PatchDisplay.
type PatchDisplay struct {
	Description *string `json:"description,omitempty"`
	Title       *string `json:"title,omitempty"`
}

// PatchItem defines model for PatchItem.
type PatchItem struct {
	ExternalLink   *string `json:"externalLink,omitempty"`
	PhotoURL       *string `json:"photoURL,omitempty"`
	SocialPostLink *string `json:"socialPostLink,omitempty"`
}

// PostDisplay defines model for PostDisplay.
type PostDisplay struct {
	Description string `json:"description"`
	Title       string `json:"title"`
}

// PostItem defines model for PostItem.
type PostItem struct {
	ExternalLink string `json:"externalLink"`
}

// Session defines model for Session.
type Session struct {
	Active                bool   `json:"active"`
	AuthenticatedAt       string `json:"authenticated_at"`
	AuthenticationMethods []struct {
		CompletedAt string `json:"completed_at"`
		Method      string `json:"method"`
	} `json:"authentication_methods"`
	AuthenticatorAssuranceLevel string `json:"authenticator_assurance_level"`
	ExpiresAt                   string `json:"expires_at"`
	ID                          string `json:"id"`
	Identity                    struct {
		CreatedAt         string `json:"created_at"`
		ID                string `json:"id"`
		RecoveryAddresses []struct {
			CreatedAt string `json:"created_at"`
			ID        string `json:"id"`
			UpdatedAt string `json:"updated_at"`
			Value     string `json:"value"`
			Via       string `json:"via"`
		} `json:"recovery_addresses"`
		SchemaID       string `json:"schema_id"`
		SchemaURL      string `json:"schema_url"`
		State          string `json:"state"`
		StateChangedAt string `json:"state_changed_at"`
		Traits         struct {
			Email   string `json:"email"`
			Website string `json:"website"`
		} `json:"traits"`
		UpdatedAt           string `json:"updated_at"`
		VerifiableAddresses []struct {
			CreatedAt string `json:"created_at"`
			ID        string `json:"id"`
			Status    string `json:"status"`
			UpdatedAt string `json:"updated_at"`
			Value     string `json:"value"`
			Verified  bool   `json:"verified"`
			Via       string `json:"via"`
		} `json:"verifiable_addresses"`
	} `json:"identity"`
	IssuedAt string `json:"issued_at"`
}

// Upload defines model for Upload.
type Upload struct {
	Component string `json:"component"`
	File      string `json:"file"`
}

// DisplayID defines model for displayID.
type DisplayID string

// ItemID defines model for itemID.
type ItemID string

// BadRequest defines model for BadRequest.
type BadRequest Error

// CreateDisplayJSONBody defines parameters for CreateDisplay.
type CreateDisplayJSONBody PostDisplay

// UpdateDisplayJSONBody defines parameters for UpdateDisplay.
type UpdateDisplayJSONBody PatchDisplay

// CreateItemJSONBody defines parameters for CreateItem.
type CreateItemJSONBody PostItem

// UpdateItemJSONBody defines parameters for UpdateItem.
type UpdateItemJSONBody PatchItem

// CreateDisplayJSONRequestBody defines body for CreateDisplay for application/json ContentType.
type CreateDisplayJSONRequestBody CreateDisplayJSONBody

// Bind implements render.Binder.
func (CreateDisplayJSONRequestBody) Bind(*http.Request) error {
	return nil
}

// UpdateDisplayJSONRequestBody defines body for UpdateDisplay for application/json ContentType.
type UpdateDisplayJSONRequestBody UpdateDisplayJSONBody

// Bind implements render.Binder.
func (UpdateDisplayJSONRequestBody) Bind(*http.Request) error {
	return nil
}

// CreateItemJSONRequestBody defines body for CreateItem for application/json ContentType.
type CreateItemJSONRequestBody CreateItemJSONBody

// Bind implements render.Binder.
func (CreateItemJSONRequestBody) Bind(*http.Request) error {
	return nil
}

// UpdateItemJSONRequestBody defines body for UpdateItem for application/json ContentType.
type UpdateItemJSONRequestBody UpdateItemJSONBody

// Bind implements render.Binder.
func (UpdateItemJSONRequestBody) Bind(*http.Request) error {
	return nil
}

// Response is a common response struct for all the API calls.
// A Response object may be instantiated via functions for specific operation responses.
// It may also be instantiated directly, for the purpose of responding with a single status code.
type Response struct {
	body        interface{}
	Code        int
	contentType string
}

// Render implements the render.Renderer interface. It sets the Content-Type header
// and status code based on the response definition.
func (resp *Response) Render(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", resp.contentType)
	render.Status(r, resp.Code)
	return nil
}

// Status is a builder method to override the default status code for a response.
func (resp *Response) Status(code int) *Response {
	resp.Code = code
	return resp
}

// ContentType is a builder method to override the default content type for a response.
func (resp *Response) ContentType(contentType string) *Response {
	resp.contentType = contentType
	return resp
}

// MarshalJSON implements the json.Marshaler interface.
// This is used to only marshal the body of the response.
func (resp *Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(resp.body)
}

// MarshalXML implements the xml.Marshaler interface.
// This is used to only marshal the body of the response.
func (resp *Response) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.Encode(resp.body)
}

// MeJSON200Response is a constructor method for a Me response.
// A *Response is returned with the configured status code and content type from the spec.
func MeJSON200Response(body Session) *Response {
	return &Response{
		body:        body,
		Code:        200,
		contentType: "application/json",
	}
}

// UploadFileJSON200Response is a constructor method for a UploadFile response.
// A *Response is returned with the configured status code and content type from the spec.
func UploadFileJSON200Response(body string) *Response {
	return &Response{
		body:        body,
		Code:        200,
		contentType: "application/json",
	}
}

// UploadFileJSON400Response is a constructor method for a UploadFile response.
// A *Response is returned with the configured status code and content type from the spec.
func UploadFileJSON400Response(body Error) *Response {
	return &Response{
		body:        body,
		Code:        400,
		contentType: "application/json",
	}
}

// CreateDisplayJSON201Response is a constructor method for a CreateDisplay response.
// A *Response is returned with the configured status code and content type from the spec.
func CreateDisplayJSON201Response(body GetDisplay) *Response {
	return &Response{
		body:        body,
		Code:        201,
		contentType: "application/json",
	}
}

// CreateDisplayJSON400Response is a constructor method for a CreateDisplay response.
// A *Response is returned with the configured status code and content type from the spec.
func CreateDisplayJSON400Response(body Error) *Response {
	return &Response{
		body:        body,
		Code:        400,
		contentType: "application/json",
	}
}

// GetDisplayJSON200Response is a constructor method for a GetDisplay response.
// A *Response is returned with the configured status code and content type from the spec.
func GetDisplayJSON200Response(body GetDisplay) *Response {
	return &Response{
		body:        body,
		Code:        200,
		contentType: "application/json",
	}
}

// UpdateDisplayJSON200Response is a constructor method for a UpdateDisplay response.
// A *Response is returned with the configured status code and content type from the spec.
func UpdateDisplayJSON200Response(body GetDisplay) *Response {
	return &Response{
		body:        body,
		Code:        200,
		contentType: "application/json",
	}
}

// UpdateDisplayJSON400Response is a constructor method for a UpdateDisplay response.
// A *Response is returned with the configured status code and content type from the spec.
func UpdateDisplayJSON400Response(body Error) *Response {
	return &Response{
		body:        body,
		Code:        400,
		contentType: "application/json",
	}
}

// CreateItemJSON201Response is a constructor method for a CreateItem response.
// A *Response is returned with the configured status code and content type from the spec.
func CreateItemJSON201Response(body GetItem) *Response {
	return &Response{
		body:        body,
		Code:        201,
		contentType: "application/json",
	}
}

// CreateItemJSON400Response is a constructor method for a CreateItem response.
// A *Response is returned with the configured status code and content type from the spec.
func CreateItemJSON400Response(body Error) *Response {
	return &Response{
		body:        body,
		Code:        400,
		contentType: "application/json",
	}
}

// GetItemJSON200Response is a constructor method for a GetItem response.
// A *Response is returned with the configured status code and content type from the spec.
func GetItemJSON200Response(body GetItem) *Response {
	return &Response{
		body:        body,
		Code:        200,
		contentType: "application/json",
	}
}

// UpdateItemJSON200Response is a constructor method for a UpdateItem response.
// A *Response is returned with the configured status code and content type from the spec.
func UpdateItemJSON200Response(body GetItem) *Response {
	return &Response{
		body:        body,
		Code:        200,
		contentType: "application/json",
	}
}

// UpdateItemJSON400Response is a constructor method for a UpdateItem response.
// A *Response is returned with the configured status code and content type from the spec.
func UpdateItemJSON400Response(body Error) *Response {
	return &Response{
		body:        body,
		Code:        400,
		contentType: "application/json",
	}
}

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get Session
	// (GET /@me)
	Me(w http.ResponseWriter, r *http.Request) *Response
	// Upload file
	// (POST /assets)
	UploadFile(w http.ResponseWriter, r *http.Request) *Response
	// Create display
	// (POST /d)
	CreateDisplay(w http.ResponseWriter, r *http.Request) *Response
	// Delete display
	// (DELETE /d/{displayID})
	DeleteDisplay(w http.ResponseWriter, r *http.Request, displayID DisplayID) *Response
	// Get display
	// (GET /d/{displayID})
	GetDisplay(w http.ResponseWriter, r *http.Request, displayID DisplayID) *Response
	// Update display
	// (PATCH /d/{displayID})
	UpdateDisplay(w http.ResponseWriter, r *http.Request, displayID DisplayID) *Response
	// Create Item
	// (POST /d/{displayID}/i)
	CreateItem(w http.ResponseWriter, r *http.Request, displayID DisplayID) *Response
	// Delete item
	// (DELETE /d/{displayID}/i/{itemID})
	DeleteItem(w http.ResponseWriter, r *http.Request, displayID DisplayID, itemID ItemID) *Response
	// Get item
	// (GET /d/{displayID}/i/{itemID})
	GetItem(w http.ResponseWriter, r *http.Request, displayID DisplayID, itemID ItemID) *Response
	// Update item
	// (PATCH /d/{displayID}/i/{itemID})
	UpdateItem(w http.ResponseWriter, r *http.Request, displayID DisplayID, itemID ItemID) *Response
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler          ServerInterface
	Middlewares      map[string]func(http.Handler) http.Handler
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// Me operation middleware
func (siw *ServerInterfaceWrapper) Me(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := siw.Handler.Me(w, r)
		if resp != nil {
			if resp.body != nil {
				render.Render(w, r, resp)
			} else {
				w.WriteHeader(resp.Code)
			}
		}
	})

	// Operation specific middleware
	handler = siw.Middlewares["session"](handler).ServeHTTP
	
	handler(w, r.WithContext(ctx))
}

// UploadFile operation middleware
func (siw *ServerInterfaceWrapper) UploadFile(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := siw.Handler.UploadFile(w, r)
		if resp != nil {
			if resp.body != nil {
				render.Render(w, r, resp)
			} else {
				w.WriteHeader(resp.Code)
			}
		}
	})

	// Operation specific middleware
	handler = siw.Middlewares["session"](handler).ServeHTTP

	handler(w, r.WithContext(ctx))
}

// CreateDisplay operation middleware
func (siw *ServerInterfaceWrapper) CreateDisplay(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := siw.Handler.CreateDisplay(w, r)
		if resp != nil {
			if resp.body != nil {
				render.Render(w, r, resp)
			} else {
				w.WriteHeader(resp.Code)
			}
		}
	})

	// Operation specific middleware
	handler = siw.Middlewares["session"](handler).ServeHTTP

	handler(w, r.WithContext(ctx))
}

// DeleteDisplay operation middleware
func (siw *ServerInterfaceWrapper) DeleteDisplay(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// ------------- Path parameter "displayID" -------------
	var displayID DisplayID

	if err := runtime.BindStyledParameter("simple", false, "displayID", chi.URLParam(r, "displayID"), &displayID); err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{err, "displayID"})
		return
	}

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := siw.Handler.DeleteDisplay(w, r, displayID)
		if resp != nil {
			if resp.body != nil {
				render.Render(w, r, resp)
			} else {
				w.WriteHeader(resp.Code)
			}
		}
	})

	// Operation specific middleware
	handler = siw.Middlewares["session"](handler).ServeHTTP

	handler(w, r.WithContext(ctx))
}

// GetDisplay operation middleware
func (siw *ServerInterfaceWrapper) GetDisplay(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// ------------- Path parameter "displayID" -------------
	var displayID DisplayID

	if err := runtime.BindStyledParameter("simple", false, "displayID", chi.URLParam(r, "displayID"), &displayID); err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{err, "displayID"})
		return
	}

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := siw.Handler.GetDisplay(w, r, displayID)
		if resp != nil {
			if resp.body != nil {
				render.Render(w, r, resp)
			} else {
				w.WriteHeader(resp.Code)
			}
		}
	})

	handler(w, r.WithContext(ctx))
}

// UpdateDisplay operation middleware
func (siw *ServerInterfaceWrapper) UpdateDisplay(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// ------------- Path parameter "displayID" -------------
	var displayID DisplayID

	if err := runtime.BindStyledParameter("simple", false, "displayID", chi.URLParam(r, "displayID"), &displayID); err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{err, "displayID"})
		return
	}

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := siw.Handler.UpdateDisplay(w, r, displayID)
		if resp != nil {
			if resp.body != nil {
				render.Render(w, r, resp)
			} else {
				w.WriteHeader(resp.Code)
			}
		}
	})

	// Operation specific middleware
	handler = siw.Middlewares["session"](handler).ServeHTTP

	handler(w, r.WithContext(ctx))
}

// CreateItem operation middleware
func (siw *ServerInterfaceWrapper) CreateItem(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// ------------- Path parameter "displayID" -------------
	var displayID DisplayID

	if err := runtime.BindStyledParameter("simple", false, "displayID", chi.URLParam(r, "displayID"), &displayID); err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{err, "displayID"})
		return
	}

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := siw.Handler.CreateItem(w, r, displayID)
		if resp != nil {
			if resp.body != nil {
				render.Render(w, r, resp)
			} else {
				w.WriteHeader(resp.Code)
			}
		}
	})

	// Operation specific middleware
	handler = siw.Middlewares["session"](handler).ServeHTTP

	handler(w, r.WithContext(ctx))
}

// DeleteItem operation middleware
func (siw *ServerInterfaceWrapper) DeleteItem(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// ------------- Path parameter "displayID" -------------
	var displayID DisplayID

	if err := runtime.BindStyledParameter("simple", false, "displayID", chi.URLParam(r, "displayID"), &displayID); err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{err, "displayID"})
		return
	}

	// ------------- Path parameter "itemID" -------------
	var itemID ItemID

	if err := runtime.BindStyledParameter("simple", false, "itemID", chi.URLParam(r, "itemID"), &itemID); err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{err, "itemID"})
		return
	}

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := siw.Handler.DeleteItem(w, r, displayID, itemID)
		if resp != nil {
			if resp.body != nil {
				render.Render(w, r, resp)
			} else {
				w.WriteHeader(resp.Code)
			}
		}
	})

	// Operation specific middleware
	handler = siw.Middlewares["session"](handler).ServeHTTP

	handler(w, r.WithContext(ctx))
}

// GetItem operation middleware
func (siw *ServerInterfaceWrapper) GetItem(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// ------------- Path parameter "displayID" -------------
	var displayID DisplayID

	if err := runtime.BindStyledParameter("simple", false, "displayID", chi.URLParam(r, "displayID"), &displayID); err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{err, "displayID"})
		return
	}

	// ------------- Path parameter "itemID" -------------
	var itemID ItemID

	if err := runtime.BindStyledParameter("simple", false, "itemID", chi.URLParam(r, "itemID"), &itemID); err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{err, "itemID"})
		return
	}

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := siw.Handler.GetItem(w, r, displayID, itemID)
		if resp != nil {
			if resp.body != nil {
				render.Render(w, r, resp)
			} else {
				w.WriteHeader(resp.Code)
			}
		}
	})

	handler(w, r.WithContext(ctx))
}

// UpdateItem operation middleware
func (siw *ServerInterfaceWrapper) UpdateItem(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// ------------- Path parameter "displayID" -------------
	var displayID DisplayID

	if err := runtime.BindStyledParameter("simple", false, "displayID", chi.URLParam(r, "displayID"), &displayID); err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{err, "displayID"})
		return
	}

	// ------------- Path parameter "itemID" -------------
	var itemID ItemID

	if err := runtime.BindStyledParameter("simple", false, "itemID", chi.URLParam(r, "itemID"), &itemID); err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{err, "itemID"})
		return
	}

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := siw.Handler.UpdateItem(w, r, displayID, itemID)
		if resp != nil {
			if resp.body != nil {
				render.Render(w, r, resp)
			} else {
				w.WriteHeader(resp.Code)
			}
		}
	})

	// Operation specific middleware
	handler = siw.Middlewares["session"](handler).ServeHTTP

	handler(w, r.WithContext(ctx))
}

type UnescapedCookieParamError struct {
	err       error
	paramName string
}

// Error implements error.
func (err UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter %s: %v", err.paramName, err.err)
}

func (err UnescapedCookieParamError) Unwrap() error { return err.err }

type UnmarshalingParamError struct {
	err       error
	paramName string
}

// Error implements error.
func (err UnmarshalingParamError) Error() string {
	return fmt.Sprintf("error unmarshaling parameter %s as JSON: %v", err.paramName, err.err)
}

func (err UnmarshalingParamError) Unwrap() error { return err.err }

type RequiredParamError struct {
	err       error
	paramName string
}

// Error implements error.
func (err RequiredParamError) Error() string {
	if err.err == nil {
		return fmt.Sprintf("query parameter %s is required, but not found", err.paramName)
	} else {
		return fmt.Sprintf("query parameter %s is required, but errored: %s", err.paramName, err.err)
	}
}

func (err RequiredParamError) Unwrap() error { return err.err }

type RequiredHeaderError struct {
	paramName string
}

// Error implements error.
func (err RequiredHeaderError) Error() string {
	return fmt.Sprintf("header parameter %s is required, but not found", err.paramName)
}

type InvalidParamFormatError struct {
	err       error
	paramName string
}

// Error implements error.
func (err InvalidParamFormatError) Error() string {
	return fmt.Sprintf("invalid format for parameter %s: %v", err.paramName, err.err)
}

func (err InvalidParamFormatError) Unwrap() error { return err.err }

type TooManyValuesForParamError struct {
	NumValues int
	paramName string
}

// Error implements error.
func (err TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("expected one value for %s, got %d", err.paramName, err.NumValues)
}

// ParameterName is an interface that is implemented by error types that are
// relevant to a specific parameter.
type ParameterError interface {
	error
	// ParamName is the name of the parameter that the error is referring to.
	ParamName() string
}

func (err UnescapedCookieParamError) ParamName() string  { return err.paramName }
func (err UnmarshalingParamError) ParamName() string     { return err.paramName }
func (err RequiredParamError) ParamName() string         { return err.paramName }
func (err RequiredHeaderError) ParamName() string        { return err.paramName }
func (err InvalidParamFormatError) ParamName() string    { return err.paramName }
func (err TooManyValuesForParamError) ParamName() string { return err.paramName }

type ServerOptions struct {
	BaseURL          string
	BaseRouter       chi.Router
	Middlewares      map[string]func(http.Handler) http.Handler
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

type ServerOption func(*ServerOptions)

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface, opts ...ServerOption) http.Handler {
	options := &ServerOptions{
		BaseURL:     "/",
		BaseRouter:  chi.NewRouter(),
		Middlewares: make(map[string]func(http.Handler) http.Handler),
		ErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		},
	}

	for _, f := range opts {
		f(options)
	}

	r := options.BaseRouter
	wrapper := ServerInterfaceWrapper{
		Handler:          si,
		Middlewares:      options.Middlewares,
		ErrorHandlerFunc: options.ErrorHandlerFunc,
	}

	middlewares := []string{"session"}
	for _, m := range middlewares {
		if _, ok := wrapper.Middlewares[m]; !ok {
			panic("goapi-gen: could not find tagged middleware " + m)
		}
	}

	r.Route(options.BaseURL, func(r chi.Router) {
		r.Get("/@me", wrapper.Me)
		r.Post("/assets", wrapper.UploadFile)
		r.Post("/d", wrapper.CreateDisplay)
		r.Delete("/d/{displayID}", wrapper.DeleteDisplay)
		r.Get("/d/{displayID}", wrapper.GetDisplay)
		r.Patch("/d/{displayID}", wrapper.UpdateDisplay)
		r.Post("/d/{displayID}/i", wrapper.CreateItem)
		r.Delete("/d/{displayID}/i/{itemID}", wrapper.DeleteItem)
		r.Get("/d/{displayID}/i/{itemID}", wrapper.GetItem)
		r.Patch("/d/{displayID}/i/{itemID}", wrapper.UpdateItem)

	})
	return r
}

func WithRouter(r chi.Router) ServerOption {
	return func(s *ServerOptions) {
		s.BaseRouter = r
	}
}

func WithServerBaseURL(url string) ServerOption {
	return func(s *ServerOptions) {
		s.BaseURL = url
	}
}

func WithMiddleware(key string, middleware func(http.Handler) http.Handler) ServerOption {
	return func(s *ServerOptions) {
		s.Middlewares[key] = middleware
	}
}

func WithMiddlewares(middlewares map[string]func(http.Handler) http.Handler) ServerOption {
	return func(s *ServerOptions) {
		s.Middlewares = middlewares
	}
}

func WithErrorHandler(handler func(w http.ResponseWriter, r *http.Request, err error)) ServerOption {
	return func(s *ServerOptions) {
		s.ErrorHandlerFunc = handler
	}
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/8xZS2/jNhf9KwK/b6lYnsdKq3YmnUHQtA3SZhUExrV4bXMqkRyS8sQI9N8LkpIlxZQf",
	"8aPdWSJ5H+ccXfLSLyQThRQcudEkfSESFBRoULknyrTMYXVzbR8YJymRYBYkJhwKJGlnPCYKv5dMISWp",
	"USXGRGcLLMAunAlVgCEpKUtGSUzMStrF2ijG56SqYsIMFoNO6sFjPFR2sZaCa3R5fQJ6j99L1MY+ZYIb",
	"5O4nSJmzDAwTPPmmBbfvWjf/VzgjKflf0mKW+FGd/KKUUN4VRZ0pJq0RklpfUePMjtYLrD2/xrqllNnp",
	"kN8pIVEZZuOcQa4xJrLzygZLsZcy4+bD+zZnxg3OUZEqJgVqDXM3exPyFsxHb7Od/7Q2JqbfMDPW1lc0",
	"157sA+PtgbERSEwY3YNArxBnb/1jGxlf0dwYLOzC2hIoBSv7LBfCiIf722AwhpkcgyOlRuUVulvNXWjd",
	"lHpxY7+vkE5ITZoDBLiUDkS/+wHvBBmfDSoO+S3jfx9D1laMtcgY5HdCm0E3R4HdS2LDXQ/tNS8tTCHo",
	"78Bki/Oof0hw1VAYb9DATlaPoysYqtDm4oB1lRD60J4GAj09pK+C6c0ORfEnal0n3PcDmWHLbrpTIXIE",
	"bhdBaRbIjd2tkE7ABOHqTGKCTwo0C0H7dfT19lLIHLdY9CZ2J13Pi/sWQ9m/rtCdmIWagNalAp7hJMcl",
	"5sGY8FkyhXooZEYHXlsvZhUAQeE2UAfsKczEEtVqApQq1PVRYwjnN7koJd22agl5Gd6+lgx2U+bKp7fh",
	"V8TdMHve9+HRb8aTgVTq0VKFKdUGDA6PTLIF8PkwEkYBMwHYsQAWdvgDp5qZPUpLMzGujYWg2MUTKjZj",
	"MM3xjFqxOJX6xDJygSMNV6Q3iKyx1+itDvoY4YX8tVLsCa+RWUBUawkNsBX83A8Lm2ldDpEQSqLeDHrl",
	"LrAN7Cqgg5tCN6JOdQyF/iBzAfTgzqU+pwelNWN5v6+ZMg5qtfPM11qtbWzG6zpMPhMb5wvy14LpiOnI",
	"LDCS5TRnWfTz3U00Eyr6AhlQHK0P7Snxb7wg/HZNxqPx6J2NXkjkIBlJyYfReGTbMdvDurSTnwqX2Bxd",
	"4hYRh/sNJSn5Dcmr1vT9eHyynrQ5VwS60j9+9c1oWRQW5dT2F1E7PyYJaI11CRU6ELrXwBeWY92aozaf",
	"BF29ir4oc8MkKJNYZq8oGNg/gVpnVZ902/1XR8LWR+Ph/taRXuccvEDYxC8mH73XUArr6JLOZUMfcp9d",
	"5FRbxeT5ai6uCkZpjj9A2bQeia4ZebKM0GEyPru60xy5t/HxdjV1D/V7MfLuZK47tw8BLnzy9HhCvKGI",
	"Nq724SR5WTePlReWPexuMnTt3rcMdS/aHsMRt1OStkGtnjZw/rgp6N9F9LkGvp+iD2Od4u4Mg2Wrw8dJ",
	"MxlfSDHB6tehXdpWO1TyKJyOwzN8od2LijMUzeMAP75cUjhAuhtfZ8J21U93E/AfpLW5pLh81W3dnrfk",
	"Nre1h9XbhCUv/t+BPUrvseTGOyfXf1ScpkIzH+5by/O/me34EgoMlnBWy2hr/b4wNGcq8/sXhIvRcaIC",
	"v5fy3UpUy4a+fjQz16pFtiOLibtiIgtjpE6TBCQb+WGQckRxmSzHpHqq/gkAAP//zi526QkdAAA=",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
