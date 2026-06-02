package stress

import (
	"net/http"
	"sync"
	"time"

	"fullcycle-stress-test/internal/report"
)

type Tester struct {
	url         string
	requests    int
	concurrency int
	client      *http.Client
}

type Result struct {
	StatusCode int
	Error      error
}

func NewTester(
	url string,
	requests int,
	concurrency int,
) *Tester {
	return &Tester{
		url:         url,
		requests:    requests,
		concurrency: concurrency,
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (t *Tester) Run() (report.Report, error) {
	start := time.Now()

	jobs := make(chan struct{}, t.requests)
	results := make(chan Result, t.requests)

	var workers sync.WaitGroup

	for i := 0; i < t.concurrency; i++ {
		workers.Add(1)

		go func() {
			defer workers.Done()

			for range jobs {
				results <- t.makeRequest()
			}
		}()
	}

	for i := 0; i < t.requests; i++ {
		jobs <- struct{}{}
	}

	close(jobs)

	go func() {
		workers.Wait()
		close(results)
	}()

	statusCodes := make(map[int]int)

	success200 := 0
	networkErrors := 0

	for result := range results {

		if result.Error != nil {
			networkErrors++
			continue
		}

		statusCodes[result.StatusCode]++

		if result.StatusCode == http.StatusOK {
			success200++
		}
	}

	return report.Report{
		Duration:      time.Since(start),
		TotalRequests: t.requests,
		Success200:    success200,
		NetworkErrors: networkErrors,
		StatusCodes:   statusCodes,
	}, nil
}

func (t *Tester) makeRequest() Result {
	resp, err := t.client.Get(t.url)
	if err != nil {
		return Result{
			Error: err,
		}
	}

	defer resp.Body.Close()

	return Result{
		StatusCode: resp.StatusCode,
	}
}