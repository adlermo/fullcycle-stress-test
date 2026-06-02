package main

import (
	"fmt"
	"log"

	"fullcycle-stress-test/internal/cli"
	"fullcycle-stress-test/internal/stress"
)

func main() {
	cfg := cli.Parse()

	tester := stress.NewTester(
		cfg.URL,
		cfg.Requests,
		cfg.Concurrency,
	)

	report, err := tester.Run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(report.String())
}
