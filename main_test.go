package main

import (
	"testing"
)

// go test ./...
// go test -cpuprofile cpu.prof
// go test -memprofile mem.prof
// to visualise output: 'go tool pprof cpu.prof' or 'go tool pprof mem.prof'
func TestUnitTestFramework(t *testing.T) {
	main()
}
