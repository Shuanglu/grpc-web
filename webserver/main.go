package webserver

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/shuanglu/grpc-web/webclient"
)

var inputVersion *string
var inputMesh string
var inputIp string
var inputDest string
var inputHost string

func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("HTTP | Received the request from %q", r.Header.Get("X-Forwarded-For"))
	fmt.Fprintf(w, "Server is running in the %q mesh. Version is %q. IP is %q", inputMesh, *inputVersion, inputIp)

}

func ingressHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("HTTP | Received the request from %q", r.Header.Get("X-Forwarded-For"))
	fmt.Fprintf(w, "Server is running in the %q mesh. Version is %q. IP is %q. Request comes from ingress", inputMesh, *inputVersion, inputIp)
	paths := strings.Split(r.URL.Path, "-")
	if len(paths) != 2 {
		webclient.Run(inputDest+"/ingress-downstream", inputHost, inputMesh, r.Header.Get("X-Forwarded-For"))
	}

}

func Run(port *int, version *string, mesh string, ip string, dest, host string) error {
	inputVersion = version
	inputMesh = mesh
	inputIp = ip
	inputDest = dest
	inputHost = host
	http.HandleFunc("/", handler)
	http.HandleFunc("/ingress", ingressHandler)
	log.Printf("web server %q listening at %s", *version, fmt.Sprintf(":%d", *port))
	if err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil); err != nil {
		return err
	}

	return nil
}
