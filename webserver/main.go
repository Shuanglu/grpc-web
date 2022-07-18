package webserver

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
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
	fmt.Fprintf(w, "Server is running in the %q mesh. Version is %q. IP is %q. Request comes from ingress\n", inputMesh, *inputVersion, inputIp)
	paths := strings.Split(r.URL.Path, "-")
	if len(paths) != 2 {
		c := http.Client{Timeout: time.Duration(1) * time.Second}
		req, err := http.NewRequest("GET", inputDest+"/ingress-downstream", nil)
		req.Header.Set("Host", inputHost)
		req.Header.Set("X-Forwarded-For", r.Header.Get("X-Forwarded-For"))
		resp, err := c.Do(req)
		if err != nil {
			log.Printf("Could not send: %s", err)
			fmt.Fprintf(w, "Failed to send request to downstream: %s", err)
		} else {
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Printf("Could not read the body: %s", err)
			}
			defer resp.Body.Close()
			log.Printf("HTTP | Client is running in the mesh: %q | %s | Request comes from ingress", inputMesh, body)
			fmt.Fprintf(w, "Response from downstream is %q\n", string(body))
		}
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
