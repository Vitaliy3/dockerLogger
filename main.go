package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
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

	log.SetOutput(os.Stdout)

	if *mode == 0 {
		log.Println(*logFile)
		data, _ := ioutil.ReadFile(*logFile)
		log.Println(string(data))
	}

	if *mode == 1 {
		prevPos := 0
		for {
			time.Sleep(time.Second * 1)
			data, _ := ioutil.ReadFile(*logFile)
			log.Println("prevpos", prevPos, len(data))
			if prevPos != len(data) {
				if len(data) < prevPos {
					prevPos = len(data)
				}
				log.Println("%s", string(data[prevPos:]))
				prevPos = len(data)
			}
		}
	}

}
