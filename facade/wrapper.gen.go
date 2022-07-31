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

// Display defines model for Display.
type Display struct {
	Description *string  `json:"description,omitempty"`
	ID          *uint64  `json:"id,omitempty"`
	ItemIDS     []uint64 `json:"itemIDs,omitempty"`
	PhotoURL    *string  `json:"photoURL,omitempty"`
	Title       *string  `json:"title,omitempty"`
	Username    *string  `json:"username,omitempty"`
}

// Error defines model for Error.
type Error struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

// Item defines model for Item.
type Item struct {
	DisplayID      *uint64 `json:"displayID,omitempty"`
	ExternalLink   *string `json:"externalLink,omitempty"`
	ID             *uint64 `json:"id,omitempty"`
	PhotoURL       *string `json:"photoURL,omitempty"`
	SocialPostLink *string `json:"socialPostLink,omitempty"`
	Username       *string `json:"username,omitempty"`
}

// Upload defines model for Upload.
type Upload struct {
	Component string `json:"component"`
	File      string `json:"file"`
}

// User defines model for User.
type User struct {
	DisplayIDS  []uint64 `json:"displayIDs,omitempty"`
	Fname       *string  `json:"fname,omitempty"`
	Lname       *string  `json:"lname,omitempty"`
	PhotoURL    *string  `json:"photoURL,omitempty"`
	SocialLinks []string `json:"socialLinks,omitempty"`
	Username    *string  `json:"username,omitempty"`
}

// CreateDisplayParams defines parameters for CreateDisplay.
type CreateDisplayParams struct {
	Display *Display `json:"display,omitempty"`
}

// UpdateDisplayJSONBody defines parameters for UpdateDisplay.
type UpdateDisplayJSONBody Display

// CreateItemParams defines parameters for CreateItem.
type CreateItemParams struct {
	Item *Item `json:"item,omitempty"`
}

// UpdateItemJSONBody defines parameters for UpdateItem.
type UpdateItemJSONBody Item

// UpdateUserJSONBody defines parameters for UpdateUser.
type UpdateUserJSONBody User

// CreateUserJSONBody defines parameters for CreateUser.
type CreateUserJSONBody User

// UpdateDisplayJSONRequestBody defines body for UpdateDisplay for application/json ContentType.
type UpdateDisplayJSONRequestBody UpdateDisplayJSONBody

// Bind implements render.Binder.
func (UpdateDisplayJSONRequestBody) Bind(*http.Request) error {
	return nil
}

// UpdateItemJSONRequestBody defines body for UpdateItem for application/json ContentType.
type UpdateItemJSONRequestBody UpdateItemJSONBody

// Bind implements render.Binder.
func (UpdateItemJSONRequestBody) Bind(*http.Request) error {
	return nil
}

// UpdateUserJSONRequestBody defines body for UpdateUser for application/json ContentType.
type UpdateUserJSONRequestBody UpdateUserJSONBody

// Bind implements render.Binder.
func (UpdateUserJSONRequestBody) Bind(*http.Request) error {
	return nil
}

// CreateUserJSONRequestBody defines body for CreateUser for application/json ContentType.
type CreateUserJSONRequestBody CreateUserJSONBody

// Bind implements render.Binder.
func (CreateUserJSONRequestBody) Bind(*http.Request) error {
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

// UploadFileJSON200Response is a constructor method for a UploadFile response.
// A *Response is returned with the configured status code and content type from the spec.
func UploadFileJSON200Response(body string) *Response {
	return &Response{
		body:        body,
		Code:        200,
		contentType: "application/json",
	}
}

// UploadFileJSONDefaultResponse is a constructor method for a UploadFile response.
// A *Response is returned with the configured status code and content type from the spec.
func UploadFileJSONDefaultResponse(body Error) *Response {
	return &Response{
		body:        body,
		Code:        200,
		contentType: "application/json",
	}
}

// CreateDisplayJSON200Response is a constructor method for a CreateDisplay response.
// A *Response is returned with the configured status code and content type from the spec.
func CreateDisplayJSON200Response(body Display) *Response {
	return &Response{
		body:        body,
		Code:        200,
		contentType: "application/json",
	}
}

// CreateDisplayJSONDefaultResponse is a constructor method for a CreateDisplay response.
// A *Response is returned with the configured status code and content type from the spec.
func CreateDisplayJSONDefaultResponse(body Error) *Response {
	return &Response{
		body:        body,
		Code:        200,
		contentType: "application/json",
	}
}

// DeleteDisplayJSONDefaultResponse is a constructor method for a DeleteDisplay response.
// A *Response is returned with the configured status code and content type from the spec.
func DeleteDisplayJSONDefaultResponse(body Error) *Response {
	return &Response{
		body:        body,
		Code:        200,
		contentType: "application/json",
	}
}

// GetDisplayJSON200Response is a constructor method for a GetDisplay response.
// A *Response is returned with the configured status code and content type from the spec.
func GetDisplayJSON200Response(body Display) *Response {
	return &Response{
		body:        body,
		Code:        200,
		contentType: "application/json",
	}
}

// GetDisplayJSONDefaultResponse is a constructor method for a GetDisplay response.
// A *Response is returned with the configured status code and content type from the spec.
func GetDisplayJSONDefaultResponse(body Error) *Response {
	return &Response{
		body:        body,
		Code:        200,
		contentType: "application/json",
	}
}

// UpdateDisplayJSON200Response is a constructor method for a UpdateDisplay response.
// A *Response is returned with the configured status code and content type from the spec.
func UpdateDisplayJSON200Response(body Display) *Response {
	return &Response{
		body:        body,
		Code:        200,
		contentType: "application/json",
	}
}

// UpdateDisplayJSONDefaultResponse is a constructor method for a UpdateDisplay response.
// A *Response is returned with the configured status code and content type from the spec.
func UpdateDisplayJSONDefaultResponse(body Error) *Response {
	return &Response{
		body:        body,
		Code:        200,
		contentType: "application/json",
	}
}

// CreateItemJSON200Response is a constructor method for a CreateItem response.
// A *Response is returned with the configured status code and content type from the spec.
func CreateItemJSON200Response(body Item) *Response {
	return &Response{
		body:        body,
		Code:        200,
		contentType: "application/json",
	}
}

// CreateItemJSONDefaultResponse is a constructor method for a CreateItem response.
// A *Response is returned with the configured status code and content type from the spec.
func CreateItemJSONDefaultResponse(body Error) *Response {
	return &Response{
		body:        body,
		Code:        200,
		contentType: "application/json",
	}
}

// DeleteItemJSONDefaultResponse is a constructor method for a DeleteItem response.
// A *Response is returned with the configured status code and content type from the spec.
func DeleteItemJSONDefaultResponse(body Error) *Response {
	return &Response{
		body:        body,
		Code:        200,
		contentType: "application/json",
	}
}

// GetItemJSON200Response is a constructor method for a GetItem response.
// A *Response is returned with the configured status code and content type from the spec.
func GetItemJSON200Response(body Item) *Response {
	return &Response{
		body:        body,
		Code:        200,
		contentType: "application/json",
	}
}

// GetItemJSONDefaultResponse is a constructor method for a GetItem response.
// A *Response is returned with the configured status code and content type from the spec.
func GetItemJSONDefaultResponse(body Error) *Response {
	return &Response{
		body:        body,
		Code:        200,
		contentType: "application/json",
	}
}

// UpdateItemJSON200Response is a constructor method for a UpdateItem response.
// A *Response is returned with the configured status code and content type from the spec.
func UpdateItemJSON200Response(body Item) *Response {
	return &Response{
		body:        body,
		Code:        200,
		contentType: "application/json",
	}
}

// UpdateItemJSONDefaultResponse is a constructor method for a UpdateItem response.
// A *Response is returned with the configured status code and content type from the spec.
func UpdateItemJSONDefaultResponse(body Error) *Response {
	return &Response{
		body:        body,
		Code:        200,
		contentType: "application/json",
	}
}

// GetUserJSON200Response is a constructor method for a GetUser response.
// A *Response is returned with the configured status code and content type from the spec.
func GetUserJSON200Response(body User) *Response {
	return &Response{
		body:        body,
		Code:        200,
		contentType: "application/json",
	}
}

// GetUserJSONDefaultResponse is a constructor method for a GetUser response.
// A *Response is returned with the configured status code and content type from the spec.
func GetUserJSONDefaultResponse(body Error) *Response {
	return &Response{
		body:        body,
		Code:        200,
		contentType: "application/json",
	}
}

// UpdateUserJSON200Response is a constructor method for a UpdateUser response.
// A *Response is returned with the configured status code and content type from the spec.
func UpdateUserJSON200Response(body User) *Response {
	return &Response{
		body:        body,
		Code:        200,
		contentType: "application/json",
	}
}

// UpdateUserJSONDefaultResponse is a constructor method for a UpdateUser response.
// A *Response is returned with the configured status code and content type from the spec.
func UpdateUserJSONDefaultResponse(body Error) *Response {
	return &Response{
		body:        body,
		Code:        200,
		contentType: "application/json",
	}
}

// CreateUserJSON200Response is a constructor method for a CreateUser response.
// A *Response is returned with the configured status code and content type from the spec.
func CreateUserJSON200Response(body User) *Response {
	return &Response{
		body:        body,
		Code:        200,
		contentType: "application/json",
	}
}

// CreateUserJSONDefaultResponse is a constructor method for a CreateUser response.
// A *Response is returned with the configured status code and content type from the spec.
func CreateUserJSONDefaultResponse(body Error) *Response {
	return &Response{
		body:        body,
		Code:        200,
		contentType: "application/json",
	}
}

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Upload file
	// (POST /assets)
	UploadFile(w http.ResponseWriter, r *http.Request) *Response
	// Create display
	// (POST /d)
	CreateDisplay(w http.ResponseWriter, r *http.Request, params CreateDisplayParams) *Response
	// Delete display
	// (DELETE /d/{displayID})
	DeleteDisplay(w http.ResponseWriter, r *http.Request, displayID uint64) *Response
	// Get display
	// (GET /d/{displayID})
	GetDisplay(w http.ResponseWriter, r *http.Request, displayID uint64) *Response
	// Update display
	// (PATCH /d/{displayID})
	UpdateDisplay(w http.ResponseWriter, r *http.Request, displayID uint64) *Response
	// Create Item
	// (POST /i)
	CreateItem(w http.ResponseWriter, r *http.Request, params CreateItemParams) *Response
	// Delete item
	// (DELETE /i/{itemID})
	DeleteItem(w http.ResponseWriter, r *http.Request, itemID uint64) *Response
	// Get item
	// (GET /i/{itemID})
	GetItem(w http.ResponseWriter, r *http.Request, itemID uint64) *Response
	// Update item
	// (PATCH /i/{itemID})
	UpdateItem(w http.ResponseWriter, r *http.Request, itemID uint64) *Response
	// Get user
	// (GET /{username})
	GetUser(w http.ResponseWriter, r *http.Request, username string) *Response
	// Update user
	// (PATCH /{username})
	UpdateUser(w http.ResponseWriter, r *http.Request, username string) *Response
	// Create user
	// (POST /{username})
	CreateUser(w http.ResponseWriter, r *http.Request, username string) *Response
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler          ServerInterface
	Middlewares      map[string]func(http.Handler) http.Handler
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
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

	handler(w, r.WithContext(ctx))
}

// CreateDisplay operation middleware
func (siw *ServerInterfaceWrapper) CreateDisplay(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Parameter object where we will unmarshal all parameters from the context
	var params CreateDisplayParams

	// ------------- Optional query parameter "display" -------------

	if err := runtime.BindQueryParameter("form", true, false, "display", r.URL.Query(), &params.Display); err != nil {
		err = fmt.Errorf("invalid format for parameter display: %w", err)
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{err, "display"})
		return
	}

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := siw.Handler.CreateDisplay(w, r, params)
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

// DeleteDisplay operation middleware
func (siw *ServerInterfaceWrapper) DeleteDisplay(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// ------------- Path parameter "displayID" -------------
	var displayID uint64

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

	handler(w, r.WithContext(ctx))
}

// GetDisplay operation middleware
func (siw *ServerInterfaceWrapper) GetDisplay(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// ------------- Path parameter "displayID" -------------
	var displayID uint64

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
	var displayID uint64

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

	handler(w, r.WithContext(ctx))
}

// CreateItem operation middleware
func (siw *ServerInterfaceWrapper) CreateItem(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Parameter object where we will unmarshal all parameters from the context
	var params CreateItemParams

	// ------------- Optional query parameter "item" -------------

	if err := runtime.BindQueryParameter("form", true, false, "item", r.URL.Query(), &params.Item); err != nil {
		err = fmt.Errorf("invalid format for parameter item: %w", err)
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{err, "item"})
		return
	}

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := siw.Handler.CreateItem(w, r, params)
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

// DeleteItem operation middleware
func (siw *ServerInterfaceWrapper) DeleteItem(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// ------------- Path parameter "itemID" -------------
	var itemID uint64

	if err := runtime.BindStyledParameter("simple", false, "itemID", chi.URLParam(r, "itemID"), &itemID); err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{err, "itemID"})
		return
	}

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := siw.Handler.DeleteItem(w, r, itemID)
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

// GetItem operation middleware
func (siw *ServerInterfaceWrapper) GetItem(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// ------------- Path parameter "itemID" -------------
	var itemID uint64

	if err := runtime.BindStyledParameter("simple", false, "itemID", chi.URLParam(r, "itemID"), &itemID); err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{err, "itemID"})
		return
	}

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := siw.Handler.GetItem(w, r, itemID)
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

	// ------------- Path parameter "itemID" -------------
	var itemID uint64

	if err := runtime.BindStyledParameter("simple", false, "itemID", chi.URLParam(r, "itemID"), &itemID); err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{err, "itemID"})
		return
	}

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := siw.Handler.UpdateItem(w, r, itemID)
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

// GetUser operation middleware
func (siw *ServerInterfaceWrapper) GetUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// ------------- Path parameter "username" -------------
	var username string

	if err := runtime.BindStyledParameter("simple", false, "username", chi.URLParam(r, "username"), &username); err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{err, "username"})
		return
	}

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := siw.Handler.GetUser(w, r, username)
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

// UpdateUser operation middleware
func (siw *ServerInterfaceWrapper) UpdateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// ------------- Path parameter "username" -------------
	var username string

	if err := runtime.BindStyledParameter("simple", false, "username", chi.URLParam(r, "username"), &username); err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{err, "username"})
		return
	}

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := siw.Handler.UpdateUser(w, r, username)
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

// CreateUser operation middleware
func (siw *ServerInterfaceWrapper) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// ------------- Path parameter "username" -------------
	var username string

	if err := runtime.BindStyledParameter("simple", false, "username", chi.URLParam(r, "username"), &username); err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{err, "username"})
		return
	}

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := siw.Handler.CreateUser(w, r, username)
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

	r.Route(options.BaseURL, func(r chi.Router) {
		r.Post("/assets", wrapper.UploadFile)
		r.Post("/d", wrapper.CreateDisplay)
		r.Delete("/d/{displayID}", wrapper.DeleteDisplay)
		r.Get("/d/{displayID}", wrapper.GetDisplay)
		r.Patch("/d/{displayID}", wrapper.UpdateDisplay)
		r.Post("/i", wrapper.CreateItem)
		r.Delete("/i/{itemID}", wrapper.DeleteItem)
		r.Get("/i/{itemID}", wrapper.GetItem)
		r.Patch("/i/{itemID}", wrapper.UpdateItem)
		r.Get("/{username}", wrapper.GetUser)
		r.Patch("/{username}", wrapper.UpdateUser)
		r.Post("/{username}", wrapper.CreateUser)

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

	"H4sIAAAAAAAC/+xY32/bNhD+V4TbHjXLS4o96G1r1sJosAXd/FQEAyOeLLYSyZKnIJ7h/30gqR9xLLl2",
	"kybO0jeZInnH+77vPosryFSllURJFtIV2KzAivnHM2F1yZbukXEuSCjJygujNBoSaCHNWWkxBn1raAUc",
	"bWaEdrPdT1pqhBQsGSEXsI5BcDecK1MxghRqIemXVxC3E4UkXKDxMwmr2Znf1D3afdc1I8wYtnS/daFI",
	"zd+fD6ZDgkocfFNbNJJVQy/7IOrqI2bkpv9ujDIH1ipTHDdOJSSdngweqkJr2WIkGYOfa2GQQ/oh7NnP",
	"vxxIdUZYHYpq4MLsbF8Q8IZc+cpzIT/dkwc78bMqE6y8UJZGIx0I5FyXivGDkWxENJhBLspNnK+EZGbZ",
	"n3YUy3bXZo8hNOcWzdeieX9x5SOVjaEcfbMHoA7MzeS2lXsnk4NgdoEwq42g5V+u44W6MC3++YS+4wkJ",
	"KRTIOBqIIezbve+Da/EOl7B2+wmZq60GCH8XwkbCRlRgpOurUmTRrxezKFcmesMyxnECXROCMAIxXKOx",
	"Yf10Mp387I6nNEqmBaRwOplOXJPQjAqfdsKsxdC+tbKegg5t5lKYcUgbSr9xDAoEQ0u/Kb4MzJXU8Laq",
	"SxKaGUocE37ijFhvCe7pR4M5pPBD0ntG0hhG0shmvclhMjX6AauVtKHKJ9PpncBM61JkPt/kow3G0Ufd",
	"LOj8/bmvnqhce9tW0Dq+s+DPd+DHclaXdFDgXccN7X4g3FzijcaMkEftnBhsXVVO8C0UkVeze5Pwcdhe",
	"G2SErQ07wA2rkNBYSD+sAG906Q3EFTkOjP1co+8rDWF5t3a/U7WxXMq09Jx0VID15T0x3CvosWMX8Ih4",
	"l6+DL1l13XQd2Foi4TaYZ358FEyPnlP0FnizM7irqNtwfrlpD4D3altXf6jodVPcIyp5qFpf8hgWOKCU",
	"t0hHUtkXKIu3SLcB0oyyYsiDOHtK/o+Z3gMB9LCW97/gRUB8s12KL7md/zb5GqsTYeF+R/JRHtXkQsRn",
	"4nBNsg6vZBU+w/fwtmHoBoQddvxGvfeYy9u4mWjYN2Zlx1HIF0R752AtKDvt6ymQeXjj6kF5PNd6HkRo",
	"LEt0/W/V3ir4/jcmWH8Fsw8tujuKXcSo2M05ygUVkJ5MY6iEbH+ebn/ufkvx+mM9C/HWtrkq3CXep0bp",
	"4YXcA/R4Qn4epGiE3PFixz/O77R4ObRo/t/WXbIWzXULeW1KSKEg0jZNkjzczmbMaOSI/06ESq6nsL5c",
	"/xcAAP//MUPOfzMbAAA=",
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
