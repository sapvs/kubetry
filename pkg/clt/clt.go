package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type em struct{}

var serverHealthURL string
var httpClient *http.Client

func init() {
	serverHost := os.Getenv("SVR")
	serverHealthURL = fmt.Sprintf("http://%s:8080/health", serverHost)
	httpClient = http.DefaultClient

}

func main() {
	stopchan := make(chan em)
	go checkServerHealth(stopchan)

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, os.Interrupt)
	<-sigchan
	stopchan <- em{}
	log.Println("Shutting down")
	<-stopchan

}

func checkServerHealth(stop chan em) {
	ticker := time.NewTicker(2 * time.Second)

	for {
		select {
		case <-stop:
			log.Println("stopping server health check")
			ticker.Stop()
			stop <- em{}
		case <-ticker.C:
			log.Println("checking server health")
			res, err := httpClient.Get(serverHealthURL) //TODO move this to config map/environment
			if err != nil || res.StatusCode != http.StatusOK {
				log.Printf("health check failed due to %s", err.Error())
			}

		}
	}
}
