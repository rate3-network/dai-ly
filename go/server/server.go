/*
Package http implements the Server component which uses http as a transport
mechanism to communication over RESTful endpoints for ratex store
*/
package server

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"runtime/debug"

	_ "github.com/lib/pq" //database driver
	"github.com/rate-engineering/dai-ly/go/database"
	"github.com/rate-engineering/dai-ly/go/zap"
)

// Server is a HTTP implementation of the project's Server interface
type Server struct {
	*http.Server
	*API
}

type API struct {
	Store       *database.Store
	Logger      *zap.Logger
	Protocol    string
	Address     string
	HostURL     string
	Environment string
}

// New returns a pointer to a http implentation of store.Server
func New() *Server {
	// Integration with AWS, do not change "PORT"
	port := os.Getenv("PORT")
	if port == "" {
		port = "5100"
	}

	s := Server{}
	s.Server = &http.Server{Addr: ":" + port}

	return &s
}

func (s *Server) Start() error {
	return s.ListenAndServe()
}

// StripWildcard strips anything after "/" to always serve index.html (the app)
func StripWildcard(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		context.WithValue(ctx, "RequestURI", r.RequestURI)
		http.StripPrefix(r.URL.Path, next).ServeHTTP(w, r)
	})
}

func RecoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func(w http.ResponseWriter, r *http.Request) {
			if rVal := recover(); rVal != nil {
				rVal := fmt.Sprint(rVal)
				// We defer again to catch the panic thrown by the logger and write ISE to client
				defer func(w http.ResponseWriter, r *http.Request) {
					if rVal := recover(); rVal != nil {
						w.WriteHeader(http.StatusInternalServerError)
						w.Write([]byte{})
					}
				}(w, r)

				// TODO do logging properly
				// print stack if the app is in development.
				// if c.App.InDevelopment() {
				debug.PrintStack()
				// }

				// Logger has no way to log a Panic level log without triggering a panic
				log.Panicf("recovering panic! panic: %v", rVal)

				return
			}

		}(w, r)

		next.ServeHTTP(w, r)
	})
}

func (a API) GetTestHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	data := map[string]string{}
	// Log the IP in test handler

	data["remote_addr"] = r.RemoteAddr

	j, _ := json.MarshalIndent(data, "", "  ")

	w.Write(j)
	return
}

func JSONHeader(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

type APIResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (res *APIResponse) SetMessage(m string) {
	res.Message = m
}

func (res *APIResponse) SetData(d interface{}) {
	res.Data = d
}

func (res *APIResponse) SetHasMore(b bool) {
	// nop
}

// Write writes and indents the HTTP response. If w.WriteHeader() has not yet
// been called, it will be called with http.StatusOK.
func (res *APIResponse) Write(w http.ResponseWriter) {
	setJSONHeader(w)
	b, _ := encodeJSON(res)
	w.Write(b)
}

// WriteData writes and indents the HTTP response with data in the Data field.
func (res *APIResponse) WriteData(w http.ResponseWriter, d interface{}) {
	res.SetData(d)
	res.Write(w)
}

// Error writes an error HTTP response with the status code attached, and puts
// the error message in the Message field.
func (res *APIResponse) Error(w http.ResponseWriter, err string, code int) {
	w.WriteHeader(code)
	res.SetMessage(err)
	res.Write(w)
}

// setJSONHeader sets the "Content-Type" header to "application/json".
func setJSONHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
}

// encodeJSON returns the JSON encoding of v with indentation.
func encodeJSON(v interface{}) ([]byte, error) {
	return json.MarshalIndent(v, "", "  ") // Two whitespaces!
}

type ListAPIResponse struct {
	APIResponse
	HasMore bool `json:"hasMore"`
}

func (res *ListAPIResponse) SetMessage(m string) {
	res.Message = m
}

func (res *ListAPIResponse) SetData(d interface{}) {
	res.Data = d
}

func (res *ListAPIResponse) SetHasMore(b bool) {
	res.HasMore = b
}

func (res ListAPIResponse) Error(w http.ResponseWriter, error string, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	res.Message = error

	j, _ := json.MarshalIndent(res, "", "  ")
	w.Write(j)
}

func parseJSON(req *http.Request, value interface{}) error {
	body := req.Body
	decoder := json.NewDecoder(body)
	err := decoder.Decode(value)
	defer body.Close()

	return err
}

// decodeJSONRequest reads and parses the JSON body of the HTTP request and
// stores the result in the value pointed to by v. Returns an error if v is nil
// or not a pointer.
//
// Take note that only exported fields of v are visible to the unmarshaller.
func decodeJSONRequest(v interface{}, r *http.Request) error {
	b, _ := ioutil.ReadAll(r.Body)
	return json.Unmarshal(b, &v)
}
