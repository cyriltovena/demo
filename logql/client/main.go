package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var paths = []string{
	"/api/v1/session",
	"/api/v2/session",
	"/api/v1/issue",
}

func url() string {
	return fmt.Sprintf("%s%s", "http://nginx", paths[rand.Intn(len(paths))])
}

func main() {
	methods := []string{http.MethodGet, http.MethodGet, http.MethodGet, http.MethodPost, http.MethodPost, http.MethodPut}

	for {
		<-time.After(time.Second / time.Duration(rand.Int63n(10)+1))
		go func() {
			req, err := http.NewRequest(methods[rand.Intn(len(methods))], url(), nil)
			if err != nil {
				log.Println(err)
			}
			_, err = http.DefaultClient.Do(req)
			if err != nil {
				log.Println(err)
			}
		}()
	}
}
