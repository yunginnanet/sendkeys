package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"

	"git.tcp.direct/kayos/sendkeys"
)

var (
	sleeptime = 10
	random    = false
	textlines []string
)

func init() {
	parseArgs()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		textlines = append(textlines, scanner.Text())
	}
	fmt.Printf("\ngot %d lines from standard input\n", len(textlines))
}

func parseArgs() {
	for i, arg := range os.Args {
		if arg == "" {
			continue
		}
		if arg == "-n" && len(os.Args) >= i {
			sleep, err := strconv.Atoi(os.Args[i+1])
			if err == nil {
				sleeptime = sleep
				os.Args[i+1] = ""
			}
		}
		if arg == "-r" {
			random = true
		}

		if arg == "-h" || arg == "--help" {
			println("Sendkeys example: pipe text into here, and it will type it out after a delay, will press enter for each line")
			println("Example: uname -a | " + os.Args[0])
			println("Flags: -n <seconds> (wait this many seconds before sending text)")
		}

	}
}

func main() {
	opts := []sendkeys.KBOpt{sendkeys.Noisy}
	if random {
		opts = append(opts, sendkeys.Random)
		println("entropy enabled")
	}
	k, err := sendkeys.NewKBWrapWithOptions(sendkeys.Noisy)
	if err != nil {
		println(err.Error())
		return
	}
	fmt.Printf("\nSleeping for %d seconds", sleeptime)
	for n := 0; n != sleeptime; n++ {
		time.Sleep(1 * time.Second)
		print(".")
	}
	println("Sending keys!")
	for _, line := range textlines {
		k.Type(line)
		time.Sleep(15 * time.Millisecond)
		k.Enter()
		time.Sleep(15 * time.Millisecond)
	}
	println("\ndone!")
}
