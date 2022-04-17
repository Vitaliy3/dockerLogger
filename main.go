package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"time"
)

func main() {
	logFile := flag.String("file", "", "")
	mode := flag.Int("mode", 0, "")
	flag.Parse()

	if *logFile == "" {
		fmt.Println("specify log_dir path by flag [logdir]")
		return
	}

	if *mode == 0 {
		fmt.Println(*logFile)
		data, _ := ioutil.ReadFile(*logFile)
		fmt.Println(string(data))
	}

	if *mode == 1 {
		prevPos := 0
		for {
			time.Sleep(time.Second * 1)
			data, _ := ioutil.ReadFile(*logFile)
			if prevPos != len(data) {
				if len(data) < prevPos {
					prevPos = len(data)
				}
				fmt.Printf("%s", string(data[prevPos:]))
				prevPos = len(data)
			}
		}
	}

}
