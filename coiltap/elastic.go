package main

import (
	"github.com/mattbaird/elastigo/api"
    "fmt"
	"github.com/mattbaird/elastigo/core"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type RequestResponse struct {
	URL        string `json:"url"`
	Method     string `json:"method"`
	StatusCode int    `json:"status"`
	Duration   int64  `json:"duration"`
	Timestamp  string `json:"timestamp"`
}

type Sink struct {
	Index   string
	Type    string
	Indexer *core.BulkIndexer
}

func (s Sink) Run() {
	done := make(chan bool)
	s.Indexer.Run(done)
}

func (s Sink) Put(req *http.Request, resp *http.Response, t time.Duration) error {
	// Timestamp is basic_date_time_no_millis
	rr := RequestResponse{
		URL:        req.URL.String(),
		Method:     req.Method,
		StatusCode: resp.StatusCode,
		Duration:   t.Nanoseconds() / 1000 / 1000,
		Timestamp:  time.Now().UTC().Format("2006-01-02T15:04:05Z"),
	}

	// add single go struct entity
	return s.Indexer.Index(s.Index, s.Type, "", "", nil, rr)
}

func NewSink(raw string) (Sink, error) {
	esurl, err := url.Parse(raw)

	if err != nil {
		return Sink{}, err
	}

	parts := strings.Split(esurl.Host, ":")
	paths := strings.Split(esurl.Path, "/")


	if len(paths) == 0 || len(parts) == 1 {
		return Sink{}, fmt.Errorf("Hostname, port, or index could not be determined")
	}

	if len(parts) == 2 {
		api.Domain = parts[0]
		api.Port = parts[1]
	} else {
		api.Domain = parts[0]
	}

    return Sink{Index: paths[1], Type: "request", Indexer: core.NewBulkIndexerErrors(10, 60)}, nil
}
