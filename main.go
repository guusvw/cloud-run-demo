// Sample run-helloworld is a minimal Cloud Run service.
package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/rs/zerolog/log"
)

func main() {
	log.Info().Int64("start_time", time.Now().Unix()).Msg("starting server...")
	http.HandleFunc("/", handler)

	// Determine port for HTTP service.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Info().Msgf("defaulting to port %s", port)
	}

	// Start HTTP server.
	log.Info().Str("port", port).Msg("listening")
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Error().Err(err).Msg("ListenAndServe failed")
		os.Exit(1)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("NAME")
	if name == "" {
		name = "Gopher"
	}
	fmt.Fprintf(w, "Hello %s!\n", name)
}
