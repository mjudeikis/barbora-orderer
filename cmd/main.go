package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/mjudeikis/barbora-orderer/pkg/brb"
)

func usage() {
	fmt.Fprint(flag.CommandLine.Output(), "usage: \n")
	fmt.Fprintf(flag.CommandLine.Output(), "       %s {har_file_name} \n", os.Args[0])
	flag.PrintDefaults()
}

func main() {

	flag.Usage = usage
	flag.Parse()

	if len(flag.Args()) != 1 {
		usage()
		os.Exit(2)
	}

	tickerChannel := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-tickerChannel.C:
			log.Print("checking")
			err := brb.Run(os.Args[1])
			if err != nil {
				// not nice, but we ignore errors and try again
				log.Println(err)
			}
		}
	}

}
