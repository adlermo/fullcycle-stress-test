package cli

import (
	"flag"
	"log"
)

type Config struct {
	URL         string
	Requests    int
	Concurrency int
}

func Parse() Config {
	url := flag.String("url", "", "Target URL")
	requests := flag.Int("requests", 0, "Total requests")
	concurrency := flag.Int("concurrency", 1, "Concurrent requests")

	flag.Parse()

	if *url == "" {
		log.Fatal("--url is required")
	}

	if *requests <= 0 {
		log.Fatal("--requests must be greater than zero")
	}

	if *concurrency <= 0 {
		log.Fatal("--concurrency must be greater than zero")
	}

	return Config{
		URL:         *url,
		Requests:    *requests,
		Concurrency: *concurrency,
	}
}
