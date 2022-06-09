package webserver

import (
	"fmt"
	"log"
	"net/http"
)

var inputVersion *string
var inputNamespace string

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s/%s", *inputVersion, inputNamespace)
}

func Run(port *int, version *string, namespace string) error {
	http.HandleFunc("/", handler)
	log.Printf("web server %q listening at %s", *version, fmt.Sprintf(":%d", *port))
	inputVersion = version
	inputNamespace = namespace
	if err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil); err != nil {
		return err
	}

	return nil
}
