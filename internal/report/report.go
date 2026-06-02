package report

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

type Report struct {
	Duration      time.Duration
	TotalRequests int
	Success200    int
	NetworkErrors int
	StatusCodes   map[int]int
}

func (r Report) String() string {
	var sb strings.Builder

	sb.WriteString("\n========== STRESS TEST REPORT ==========\n\n")

	sb.WriteString(fmt.Sprintf("Total execution time: %s\n", r.Duration))
	sb.WriteString(fmt.Sprintf("Total requests: %d\n", r.TotalRequests))
	sb.WriteString(fmt.Sprintf("HTTP 200 responses: %d\n", r.Success200))
	sb.WriteString(fmt.Sprintf("Network errors: %d\n\n", r.NetworkErrors))

	sb.WriteString("Status code distribution:\n")

	var codes []int

	for code := range r.StatusCodes {
		codes = append(codes, code)
	}

	sort.Ints(codes)

	for _, code := range codes {
		sb.WriteString(
			fmt.Sprintf(
				"  %d -> %d\n",
				code,
				r.StatusCodes[code],
			),
		)
	}

	return sb.String()
}
