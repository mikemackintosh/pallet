package main

import (
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/logging"
	"github.com/go-chi/chi"
	"github.com/mikemackintosh/pallet"
)

func main() {
	// Set listening port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	// Create the router
	r := chi.NewRouter()

	loggingOptions := pallet.NewOptionSet()
	loggingOptions.SetLabels(map[string]string{
		"test": "test",
	})
	loggingOptions.SetProject(os.Getenv("GOOGLE_CLOUD_PROJECT"))

	// A good base middleware stack.
	//r.Use(MiddlewareSlackWrapper)
	r.Use(pallet.DefaultMiddleware(loggingOptions))
	r.Get("/", handler)

	// Listen and serve
	log.Fatal(http.ListenAndServe(":"+port, r))
}

// handler is a sample route
func handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	logger := pallet.GetLoggerFromRequest(r)
	logger.Logf(logging.Warning, "url: %+v", r.URL.Path)

	_, err := w.Write([]byte("{\"status\": \"ok\"}"))
	if err != nil {
		log.Fatal(err)
	}
}
