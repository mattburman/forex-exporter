package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	appId := os.Getenv("APP_ID")
	if len(appId) != 32 {
		log.Fatal("APP_ID was not set or not 32 characters")
	}
	port, err := strconv.ParseUint(os.Getenv("PORT"), 10, 32)
	if err != nil {
		log.Printf("PORT defaulting to %d\n", 8080)
		port = 8080
	}

	rootHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "See /metrics for the prometheus metrics.\n")
	}
	http.HandleFunc("/", rootHandler)

	http.Handle("/metrics", promhttp.Handler())

	// TODO: make duration configurable.
	errs := runCollector(appId, NewMetrics(), 1*time.Hour)
	go func() {
		for err := range errs {
			log.Println(errors.Wrap(err, "error collecting"))
		}
	}()

	addr := fmt.Sprintf(":%d", port)
	log.Printf("Listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
