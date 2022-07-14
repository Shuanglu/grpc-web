package webclient

import (
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

func Run(httpAddr string, host string, mesh string) error {
	c := http.Client{Timeout: time.Duration(1) * time.Second}
	req, err := http.NewRequest("GET", httpAddr, nil)
	if err != nil {
		log.Printf("Failed to generate request: ", err)
	}
	req.Header.Set("Host", host)
	var wg sync.WaitGroup
	for {
		wg.Add(1)
		go func() {
			resp, err := c.Do(req)
			if err != nil {
				log.Printf("Could not send: %s", err)
			} else {
				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					log.Printf("Could not read the body: ", err)
				}
				defer resp.Body.Close()
				log.Printf("HTTP | Client is running in the mesh: %q | %s ", mesh, body)
			}

			wg.Done()
		}()
		time.Sleep(5 * time.Second)
	}
	wg.Wait()
	return nil
}
