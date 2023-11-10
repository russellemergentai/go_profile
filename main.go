package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"runtime"
	"runtime/pprof"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
var memprofile = flag.String("memprofile", "", "write memory profile to `file`")

func calculate(ch chan float64) {
	acc := 0
	n := 10000000

	for i := 0; i < n; i++ {
		var x = rand.Float64()
		var y = rand.Float64()
		xx := math.Pow(x, 2)
		yy := math.Pow(y, 2)
		r := (xx + yy)
		if r < 1.0 {
			acc++
		}
	}

	output := float64(acc) * 4 / float64(n)
	ch <- output
}

func main() {

	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		defer f.Close() // error handling omitted for example
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	ch1 := make(chan float64)
	go calculate(ch1)

	ch2 := make(chan float64)
	go calculate(ch2)

	val1 := <-ch1
	val2 := <-ch2
	fmt.Println(val1)
	fmt.Println(val2)

	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		defer f.Close() // error handling omitted for example
		runtime.GC()    // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
	}
}
