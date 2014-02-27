package main

import (
	"github.com/mattbaird/elastigo/api"
	"github.com/mattbaird/elastigo/core"
	"log"
	"net/http"
	"time"
)

type RequestResponse struct {
	URL        string `json:"url"`
	Method     string `json:"method"`
	StatusCode int    `json:"status"`
	Duration   int64  `json:"duration"`
}

func TrackRequest(req *http.Request, resp *http.Response, t time.Duration) error {

	api.Domain = "localhost"

	rr := RequestResponse{
		URL:        req.URL.String(),
		Method:     req.Method,
		StatusCode: resp.StatusCode,
		Duration:   t.Nanoseconds(),
	}

	log.Println(req, resp)

	// add single go struct entity
	_, err := core.Index(true, "requests", "pinger", "", rr)

	if err != nil {
		log.Println(err)
	}

	return err
}
