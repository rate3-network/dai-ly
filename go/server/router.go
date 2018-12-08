package server

// NewRouter configure mux router for the server. Must be run after SetAPI
import (
	"net/http"
	"os"
	"time"

	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/gorilla/handlers"
	"github.com/rs/cors"
	"github.com/unrolled/secure"
)

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
	api.Post("/hash", s.API.GetHashHandler)
	api.Post("/transaction", s.API.CreateTransactionHandler)
	api.Get("/transactions/{hash}", s.API.PollTransactionHandler)
	api.Get("/transactions", s.API.GetTransactionsHandler)
	api.Get("/queued", s.API.GetOneQueuedTransactionHandler)

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
