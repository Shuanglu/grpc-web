package webclient

import (
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

func Run(httpAddr string, host string, mesh string, ip string, client_request_total int) error {
	c := http.Client{Timeout: time.Duration(1) * time.Second}
	req, err := http.NewRequest("GET", httpAddr, nil)
	if err != nil {
		log.Printf("Failed to generate request: %s", err)
	}
	req.Host = host
	req.Header.Set("X-Forwarded-For", ip)
	var wg sync.WaitGroup
	count := 0
	for {
		wg.Add(1)
		go func(count *int) {
			resp, err := c.Do(req)
			if err != nil {
				log.Printf("Could not send: %s", err)
			} else {
				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					log.Printf("Could not read the body: %s", err)
				}
				defer resp.Body.Close()
				log.Printf("HTTP | Client is running in the mesh: %q | %s ", mesh, body)
				*count++
			}
			wg.Done()
		}(&count)
		if client_request_total == 0 {
			continue
		} else if client_request_total != 0 && count == client_request_total {
			log.Printf("Sent %d requests to server. Will sleep forever", client_request_total)
			time.Sleep(time.Duration(1<<63 - 1))
		}
		time.Sleep(5 * time.Second)
	}
	wg.Wait()
	return nil
}
