package webserver

import (
	"fmt"
	"log"
	"net/http"
)

var inputVersion *string

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, this is version %s!", *inputVersion)
}

func Run(port *int, version *string) error {
	http.HandleFunc("/", handler)
	log.Printf("web server %q listening at %s", *version, fmt.Sprintf(":%d", *port))
	inputVersion = version
	if err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil); err != nil {
		return err
	}

	return nil
}
