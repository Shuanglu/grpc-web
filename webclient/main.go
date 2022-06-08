package webclient

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func Run(httpAddr string, host string) error {
	c := http.Client{Timeout: time.Duration(1) * time.Second}
	req, err := http.NewRequest("GET", httpAddr, nil)
	req.Header.Set("Host", host)
	resp, err := c.Do(req)
	if err != nil {
		log.Printf("Error %s", err)
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	log.Printf("HTTP Response: %s", body)
	return nil
}
