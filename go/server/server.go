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
	"time"

	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/gorilla/handlers"
	_ "github.com/lib/pq" //database driver
	"github.com/rate-engineering/dai-ly/go/database"
	"github.com/rate-engineering/dai-ly/go/zap"
	"github.com/rs/cors"
	"github.com/unrolled/secure"
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

// NewRouter configure mux router for the server. Must be run after SetAPI
func (s *Server) SetRouter() error {
	// Security headers that we set w/o depending on reverse proxy
	secureMiddleware := secure.New(secure.Options{
		SSLProxyHeaders:      map[string]string{"X-Forwarded-Proto": "https"},
		STSSeconds:           315360000,
		STSIncludeSubdomains: true,
		STSPreload:           true,
		FrameDeny:            true,
		ContentTypeNosniff:   true,
		BrowserXssFilter:     true,
		IsDevelopment:        true,
	})

	// You can create a generic limiter for all your handlers
	// or one for each handler. Your choice.
	// This limiter basically says: allow at most 5 request per 1 second.
	limiter := tollbooth.NewLimiter(3, &limiter.ExpirableOptions{DefaultExpirationTTL: time.Hour})

	// Configure list of places to look for IP address.
	// "X-Forwarded-FOr" should be first but we have gorilla.ProxyHeaders
	limiter.SetIPLookups([]string{"X-Forwarded-For", "RemoteAddr", "X-Real-IP"})

	// This is an example on how to limit only GET and POST requests.
	limiter.SetMethods([]string{"GET", "POST"})

	// if !c.App.InProduction() {
	corHeaders := cors.New(cors.Options{
		AllowedOrigins: []string{
			// localhosts
			"http://192.168.0.199:3000",
			"http://localhost:5100",
			"http://localhost:3000"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"*"},
	})
	// }

	// Routes
	main := chi.NewRouter()
	main.Use(RecoverPanic)
	main.Use(middleware.StripSlashes)

	api := chi.NewRouter()
	api.Use(JSONHeader)
	api.Get("/test", s.API.GetTestHandler)
	api.Get("/transactions", s.API.GetTransactionsHandler)

	// API Router Initiation
	main.Mount("/api", api)

	// Serve client
	fs := http.FileServer(http.Dir("./view"))
	main.Mount("/static", fs)
	main.Mount("/favicon.ico", fs)
	main.Mount("/manifiest.json", fs)
	main.Mount("/asset-manifest.json", fs)
	main.Mount("/service-worker.js", fs)

	s.Handler = (handlers.CombinedLoggingHandler(
		os.Stdout, tollbooth.LimitHandler(limiter, secureMiddleware.Handler(corHeaders.Handler(handlers.ProxyHeaders((main)))))))
	return nil
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
