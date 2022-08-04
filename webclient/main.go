package webclient

import (
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

func Run(httpAddr string, host string, mesh string, ip string, client_success_request_total int, client_failure_request_total int) error {
	c := http.Client{
		Timeout: time.Duration(1) * time.Second,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	req, err := http.NewRequest("GET", httpAddr, nil)
	if err != nil {
		log.Printf("Failed to generate request: %s", err)
	}
	req.Host = host
	req.Header.Set("X-Forwarded-For", ip)
	var wg sync.WaitGroup
	successCount := 0
	failureCount := 0
	for {
		wg.Add(1)
		go func(successCount *int, failureCount *int) {
			resp, err := c.Do(req)
			if err != nil {
				log.Printf("Could not send: %s", err)
				*failureCount++
			} else {
				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					log.Printf("Could not read the body: %s", err)
				}
				redirectURL, err := req.Response.Location()
				if err == http.ErrNoLocation || redirectURL == nil {
					log.Printf("HTTP | Client is running in the mesh: %q | %s ", mesh, body)
				} else {
					log.Printf("HTTP | Client is running in the mesh: %q | %s | The request is redirected to %s", mesh, body, redirectURL.Host+redirectURL.Path)
				}
				defer resp.Body.Close()

				*successCount++
			}
			wg.Done()
		}(&successCount, &failureCount)
		if client_failure_request_total == 0 || client_success_request_total == 0 {
			time.Sleep(5 * time.Second)
			continue
		} else if successCount == client_success_request_total {
			log.Printf("Sent %d requests to server. Will sleep forever", client_success_request_total)
			time.Sleep(time.Duration(1<<63 - 1))
		} else if failureCount == client_failure_request_total {
			log.Printf("Failed to send %d requests to server. Will sleep forever", client_failure_request_total)
			time.Sleep(time.Duration(1<<63 - 1))
		} else {
			time.Sleep(5 * time.Second)
			continue
		}

	}
	wg.Wait()
	return nil
}
