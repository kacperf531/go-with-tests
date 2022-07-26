package main

import (
	"os"
	"time"

	clockface "github.com/kacperf531/go-with-tests/clock"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
