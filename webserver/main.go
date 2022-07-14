package webserver

import (
	"fmt"
	"log"
	"net/http"
)

var inputVersion *string
var inputMesh string
var inputIp string

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Server is running in the %q mesh. Version is %q. IP is %q", inputMesh, *inputVersion, inputIp)
}

func Run(port *int, version *string, mesh string, ip string) error {
	http.HandleFunc("/", handler)
	log.Printf("web server %q listening at %s", *version, fmt.Sprintf(":%d", *port))
	inputVersion = version
	inputMesh = mesh
	inputIp = ip
	if err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil); err != nil {
		return err
	}

	return nil
}
