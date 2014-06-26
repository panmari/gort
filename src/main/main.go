package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"scenes"
	"time"
	"renderer"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
var memprofile = flag.String("memprofile", "", "write memory profile to this file")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
	start := time.Now()
	
	scene := scenes.MakeSimpleScene()
	renderer.StartRendering(scene)
	
	duration := time.Since(start)
	fmt.Println(duration.String())
	scene.Film.WriteToPng(scene.Filename)
	fmt.Printf("Printed to %s\n", scene.Filename)
	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.WriteHeapProfile(f)
		f.Close()
		return
	}
}
